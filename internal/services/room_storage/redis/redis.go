package redis

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/topfreegames/maestro/internal/core/entities/game_room"

	"github.com/go-redis/redis"
	"github.com/topfreegames/maestro/internal/core/entities"
	"github.com/topfreegames/maestro/internal/services/room_storage"
)

type redisStateStorage struct {
	client *redis.Client
}

var _ room_storage.RoomStorage = (*redisStateStorage)(nil)

func NewRedisStateStorage(client *redis.Client) *redisStateStorage {
	return &redisStateStorage{client: client}
}

func (r redisStateStorage) GetRoom(ctx context.Context, scheduler, roomID string) (*game_room.GameRoom, error) {
	room := &game_room.GameRoom{
		ID:        roomID,
		Scheduler: entities.Scheduler{ID: scheduler},
	}

	p := r.client.WithContext(ctx).Pipeline()
	metadataCmd := p.Get(getRoomRedisKey(room.Scheduler.ID, room.ID))
	statusCmd := p.ZScore(getRoomStatusSetRedisKey(room.Scheduler.ID), room.ID)
	pingCmd := p.ZScore(getRoomPingRedisKey(room.Scheduler.ID), room.ID)
	_, err := p.Exec()
	if err != nil {
		if err == redis.Nil {
			return nil, room_storage.NewRoomNotFoundError(scheduler, roomID)
		}
		return nil, room_storage.WrapError("error storing room on redis", err)
	}

	room.Status = game_room.GameRoomStatus(statusCmd.Val())
	room.LastPingAt = time.Unix(int64(pingCmd.Val()), 0)
	err = json.NewDecoder(strings.NewReader(metadataCmd.Val())).Decode(&room.Metadata)
	if err != nil {
		return nil, room_storage.WrapError("error unmarshalling json", err)
	}

	return room, nil
}

func (r *redisStateStorage) CreateRoom(ctx context.Context, room *game_room.GameRoom) error {
	metadataJson, err := json.Marshal(room.Metadata)
	if err != nil {
		return err
	}

	p := r.client.WithContext(ctx).TxPipeline()
	roomCmd := p.SetNX(getRoomRedisKey(room.Scheduler.ID, room.ID), metadataJson, 0)
	statusCmd := p.ZAddNX(getRoomStatusSetRedisKey(room.Scheduler.ID), redis.Z{
		Member: room.ID,
		Score:  float64(room.Status),
	})
	pingCmd := p.ZAddNX(getRoomPingRedisKey(room.Scheduler.ID), redis.Z{
		Member: room.ID,
		Score:  float64(room.LastPingAt.Unix()),
	})

	_, err = p.Exec()
	if err != nil {
		return room_storage.WrapError("error storing room on redis", err)
	}

	if !roomCmd.Val() || statusCmd.Val() < 1 || pingCmd.Val() < 1 {
		return room_storage.NewRoomAlreadyExistsError(room.Scheduler.ID, room.ID)
	}

	return nil
}

func (r *redisStateStorage) UpdateRoom(ctx context.Context, room *game_room.GameRoom) error {
	metadataJson, err := json.Marshal(room.Metadata)
	if err != nil {
		return err
	}

	p := r.client.WithContext(ctx).TxPipeline()
	roomCmd := p.SetXX(getRoomRedisKey(room.Scheduler.ID, room.ID), metadataJson, 0)
	p.ZAddXXCh(getRoomStatusSetRedisKey(room.Scheduler.ID), redis.Z{
		Member: room.ID,
		Score:  float64(room.Status),
	})
	p.ZAddXXCh(getRoomPingRedisKey(room.Scheduler.ID), redis.Z{
		Member: room.ID,
		Score:  float64(room.LastPingAt.Unix()),
	})

	_, err = p.Exec()
	if err != nil {
		return room_storage.WrapError("error updating room on redis", err)
	}

	if !roomCmd.Val() {
		return room_storage.NewRoomNotFoundError(room.Scheduler.ID, room.ID)
	}

	return nil
}

func (r *redisStateStorage) RemoveRoom(ctx context.Context, scheduler, roomID string) error {
	p := r.client.WithContext(ctx).TxPipeline()
	p.Del(getRoomRedisKey(scheduler, roomID))
	p.ZRem(getRoomStatusSetRedisKey(scheduler), roomID)
	p.ZRem(getRoomPingRedisKey(scheduler), roomID)
	cmders, err := p.Exec()
	if err != nil {
		return room_storage.WrapError("error removing room from redis", err)
	}
	for _, cmder := range cmders {
		cmd := cmder.(*redis.IntCmd)
		if cmd.Val() == 0 {
			return room_storage.NewRoomNotFoundError(scheduler, roomID)
		}
	}
	return nil
}

func (r *redisStateStorage) SetRoomStatus(ctx context.Context, scheduler, roomID string, status game_room.GameRoomStatus) error {
	err := r.client.WithContext(ctx).ZAddXXCh(getRoomStatusSetRedisKey(scheduler), redis.Z{
		Member: roomID,
		Score:  float64(status),
	}).Err()
	if err != nil {
		return room_storage.WrapError("error updating room on redis", err)
	}
	return nil
}

func (r *redisStateStorage) GetAllRoomIDs(ctx context.Context, scheduler string) ([]string, error) {
	rooms, err := r.client.WithContext(ctx).ZRange(getRoomStatusSetRedisKey(scheduler), 0, -1).Result()
	if err != nil {
		return nil, room_storage.WrapError("error listing rooms on redis", err)
	}
	return rooms, nil
}

func (r *redisStateStorage) GetRoomIDsByLastPing(ctx context.Context, scheduler string, threshold time.Time) ([]string, error) {
	rooms, err := r.client.WithContext(ctx).ZRangeByScore(getRoomPingRedisKey(scheduler), redis.ZRangeBy{
		Min: "-inf",
		Max: strconv.FormatInt(threshold.Unix(), 10),
	}).Result()
	if err != nil {
		return nil, room_storage.WrapError("error listing rooms on redis", err)
	}
	return rooms, nil
}

func (r *redisStateStorage) GetRoomCount(ctx context.Context, scheduler string) (int, error) {
	client := r.client.WithContext(ctx)
	count, err := client.ZCard(getRoomStatusSetRedisKey(scheduler)).Result()
	if err != nil {
		return 0, room_storage.WrapError("error counting rooms on redis", err)
	}
	return int(count), nil
}

func (r *redisStateStorage) GetRoomCountByStatus(ctx context.Context, scheduler string, status game_room.GameRoomStatus) (int, error) {
	client := r.client.WithContext(ctx)
	statusIntStr := fmt.Sprint(int(status))
	count, err := client.ZCount(getRoomStatusSetRedisKey(scheduler), statusIntStr, statusIntStr).Result()
	if err != nil {
		return 0, room_storage.WrapError("error counting rooms on redis", err)
	}
	return int(count), nil
}

func getRoomRedisKey(scheduler, roomID string) string {
	return fmt.Sprintf("scheduler:%s:rooms:%s", scheduler, roomID)
}

func getRoomStatusSetRedisKey(scheduler string) string {
	return fmt.Sprintf("scheduler:%s:status", scheduler)
}

func getRoomPingRedisKey(scheduler string) string {
	return fmt.Sprintf("scheduler:%s:ping", scheduler)
}
