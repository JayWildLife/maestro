// MIT License
//
// Copyright (c) 2021 TFG Co
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package events_forwarder

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/topfreegames/maestro/internal/core/logs"

	"github.com/topfreegames/maestro/internal/core/entities"
	"github.com/topfreegames/maestro/internal/core/entities/game_room"

	"github.com/topfreegames/maestro/internal/core/entities/forwarder"

	"github.com/topfreegames/maestro/internal/core/entities/events"

	"github.com/topfreegames/maestro/internal/core/ports"
	"go.uber.org/zap"
)

var (
	_ ports.EventsService = (*EventsForwarderService)(nil)
)

type EventsForwarderService struct {
	eventsForwarder  ports.EventsForwarder
	logger           *zap.Logger
	schedulerStorage ports.SchedulerStorage
	instanceStorage  ports.GameRoomInstanceStorage
	roomStorage      ports.RoomStorage
	schedulerCache   ports.SchedulerCache
	config           EventsForwarderConfig
}

func NewEventsForwarderService(
	eventsForwarder ports.EventsForwarder,
	schedulerStorage ports.SchedulerStorage,
	instanceStorage ports.GameRoomInstanceStorage,
	roomStorage ports.RoomStorage,
	schedulerCache ports.SchedulerCache,
	config EventsForwarderConfig,
) ports.EventsService {
	return &EventsForwarderService{
		eventsForwarder,
		zap.L().With(zap.String(logs.LogFieldComponent, "service"), zap.String(logs.LogFieldServiceName, "events_forwarder")),
		schedulerStorage,
		instanceStorage,
		roomStorage,
		schedulerCache,
		config,
	}
}

func (es *EventsForwarderService) ProduceEvent(ctx context.Context, event *events.Event) error {
	if _, ok := event.Attributes["eventType"].(string); !ok {
		return errors.New("eventAttributes must contain key \"eventType\"")
	}
	eventType := event.Attributes["eventType"].(string)

	scheduler, err := es.getScheduler(ctx, event.SchedulerID)
	if err != nil {
		return err
	}

	if forwarderList := scheduler.Forwarders; len(forwarderList) > 0 {
		for _, _forwarder := range forwarderList {
			if _forwarder.Enabled {
				switch event.Name {
				case events.RoomEvent:
					err = es.forwardRoomEvent(ctx, event, eventType, scheduler, _forwarder)
					if err != nil {
						return err
					}
				case events.PlayerEvent:
					err = es.forwardPlayerEvent(ctx, event, eventType, scheduler, _forwarder)
					if err != nil {
						return err
					}
				}
			}
		}
	} else {
		es.logger.Debug(fmt.Sprintf("scheduler \"%v\" do not have forwarders configured", event.SchedulerID))
	}

	return nil
}

func (es *EventsForwarderService) forwardRoomEvent(
	ctx context.Context,
	event *events.Event,
	eventType string,
	scheduler *entities.Scheduler,
	_forwarder *forwarder.Forwarder,
) error {
	var instance *game_room.Instance

	if es.isRoomInUnreliableState(event) {
		isValidationRoom, err := es.isValidationRoom(ctx, event)
		if err == nil && isValidationRoom {
			return nil
		}

		instance, err = es.instanceStorage.GetInstance(ctx, event.SchedulerID, event.RoomID)
		if err != nil {
			instance = &game_room.Instance{Address: &game_room.Address{Host: "", Ports: []game_room.Port{{Port: 0}}}}
		}

	} else {
		isValidationRoom, err := es.isValidationRoom(ctx, event)
		if err != nil {
			return err
		}

		if isValidationRoom {
			return nil
		}

		instance, err = es.instanceStorage.GetInstance(ctx, event.SchedulerID, event.RoomID)
		if err != nil {
			es.logger.Error(fmt.Sprintf("Failed to get instance for room \"%v\" from scheduler \"%v\" info", event.RoomID, event.SchedulerID), zap.Error(err))
			return err
		}
	}

	selectedPort, err := es.selectPort(instance.Address)
	if err != nil {
		return fmt.Errorf("no room port found to forward roomEvent. Forwarder name: \"%v\", Scheduler: \"%v\"", _forwarder.Name, event.SchedulerID)
	}
	err = es.incrementEventAttributesWithPortsInfo(event, instance.Address.Ports)
	if err != nil {
		return err
	}

	roomEvent, err := events.ConvertToRoomEventType(eventType)
	if err != nil {
		return err
	}

	var pingType events.RoomStatusType
	if roomEvent == events.Ping {
		if _, ok := event.Attributes["pingType"]; !ok {
			return errors.New("roomEvent of type ping must contain key \"pingType\" in eventAttributes")
		}
		pingType, err = events.ConvertToRoomPingEventType(event.Attributes["pingType"].(string))
		if err != nil {
			return err
		}
	}

	roomAttributes := events.RoomEventAttributes{
		Game:           scheduler.Game,
		RoomId:         event.RoomID,
		Host:           instance.Address.Host,
		Port:           selectedPort,
		EventType:      roomEvent,
		RoomStatusType: &pingType,
		Other:          event.Attributes,
	}
	err = es.eventsForwarder.ForwardRoomEvent(ctx, roomAttributes, *_forwarder)
	if err != nil {
		reportRoomEventForwardingFailed(scheduler.Game, event.SchedulerID)
		es.logger.Error(fmt.Sprintf("Failed to forward room events for room %s and scheduler %s", event.RoomID, event.SchedulerID), zap.Error(err))
		return err
	}

	reportRoomEventForwardingSuccess(scheduler.Game, event.SchedulerID)
	return nil
}

func (es *EventsForwarderService) forwardPlayerEvent(
	ctx context.Context,
	event *events.Event,
	eventType string,
	scheduler *entities.Scheduler,
	_forwarder *forwarder.Forwarder,
) error {
	playerId, err := es.getPlayerInfo(event)
	if err != nil {
		return fmt.Errorf("eventAttributes must contain key \"playerId\" in playerEvent events. Forwarder name: \"%v\", Scheduler: \"%v\"", _forwarder.Name, event.SchedulerID)
	}

	playerEvent, err := events.ConvertToPlayerEventType(eventType)
	if err != nil {
		return err
	}

	playerAttributes := events.PlayerEventAttributes{
		RoomId:    event.RoomID,
		PlayerId:  playerId,
		EventType: playerEvent,
		Other:     event.Attributes,
	}

	err = es.eventsForwarder.ForwardPlayerEvent(ctx, playerAttributes, *_forwarder)
	if err != nil {
		reportPlayerEventForwardingFailed(scheduler.Game, event.SchedulerID)
		es.logger.Error(fmt.Sprintf("Failed to forward player events for room %s and scheduler %s", event.RoomID, event.SchedulerID), zap.Error(err))
		return err
	}
	reportPlayerEventForwardingSuccess(scheduler.Game, event.SchedulerID)
	return nil
}

func (es *EventsForwarderService) getScheduler(ctx context.Context, schedulerName string) (*entities.Scheduler, error) {
	scheduler, err := es.schedulerCache.GetScheduler(ctx, schedulerName)
	if err != nil {
		es.logger.Error(fmt.Sprintf("Failed to get scheduler \"%v\" from cache", schedulerName), zap.Error(err))
	}
	if scheduler == nil {
		scheduler, err = es.schedulerStorage.GetScheduler(ctx, schedulerName)
		if err != nil {
			es.logger.Error(fmt.Sprintf("Failed to get scheduler \"%v\" info", schedulerName), zap.Error(err))
			return nil, err
		}
		if err = es.schedulerCache.SetScheduler(ctx, scheduler, es.config.SchedulerCacheTtl); err != nil {
			es.logger.Error(fmt.Sprintf("Failed to set scheduler \"%v\" in cache", schedulerName), zap.Error(err))
		}
	}
	return scheduler, nil
}

func (es *EventsForwarderService) selectPort(address *game_room.Address) (int32, error) {
	if address == nil {
		return 0, errors.New("port not found, address is nil")
	}
	ports := address.Ports
	if len(ports) == 0 {
		return 0, errors.New("port not found")
	}
	selectedPort := ports[0].Port

	for _, port := range ports {
		if port.Name == "clientPort" {
			selectedPort = port.Port
			break
		}
	}

	return selectedPort, nil
}

func (es *EventsForwarderService) incrementEventAttributesWithPortsInfo(event *events.Event, ports []game_room.Port) error {
	if len(ports) == 0 {
		return nil
	}
	portsMap := make([]map[string]interface{}, 0)
	for _, port := range ports {
		portsMap = append(portsMap, map[string]interface{}{
			"port":     port.Port,
			"name":     port.Name,
			"protocol": port.Protocol,
		})
	}

	portsJson, err := json.Marshal(portsMap)
	if err != nil {
		return err
	}

	event.Attributes["ports"] = string(portsJson)
	return nil
}

func (es *EventsForwarderService) getPlayerInfo(event *events.Event) (string, error) {
	if _, ok := event.Attributes["playerId"]; !ok {
		return "", fmt.Errorf("playerId not found on eventAttributes")
	}

	playerId, ok := event.Attributes["playerId"].(string)
	if !ok {
		return "", fmt.Errorf("playerId must be a string")
	}
	return playerId, nil
}

func (es *EventsForwarderService) isValidationRoom(ctx context.Context, event *events.Event) (bool, error) {
	gameRoom, err := es.roomStorage.GetRoom(ctx, event.SchedulerID, event.RoomID)
	if err != nil {
		es.logger.Error(fmt.Sprintf("cannot produce event since room \"%s\" is not registered on storage for scheduler \"%s\"", event.RoomID, event.SchedulerID))
		return false, err
	}

	if gameRoom.IsValidationRoom {
		es.logger.Info(fmt.Sprintf("not producing events for room \"%s\", scheduler \"%s\" since it's a validation room", gameRoom.ID, gameRoom.SchedulerID))
	}

	return gameRoom.IsValidationRoom, nil
}

func (es *EventsForwarderService) isRoomInUnreliableState(event *events.Event) bool {
	if roomEvent, ok := event.Attributes["roomEvent"].(string); ok {
		if roomEvent == game_room.GameRoomPingStatusTerminating.String() ||
			roomEvent == game_room.GameRoomPingStatusTerminated.String() {
			return true
		}
	}
	if roomEvent, ok := event.Attributes["pingType"].(string); ok {
		if roomEvent == game_room.GameRoomPingStatusTerminating.String() ||
			roomEvent == game_room.GameRoomPingStatusTerminated.String() {
			return true
		}
	}

	return false
}
