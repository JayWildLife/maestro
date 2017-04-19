// maestro
// https://github.com/topfreegames/maestro
//
// Licensed under the MIT license:
// http://www.opensource.org/licenses/mit-license
// Copyright © 2017 Top Free Games <backend@tfgco.com>

package models

// RoomPingPayload is the struct that defines the payload for the ping route
type RoomPingPayload struct {
	Timestamp int64 `json:"timestamp" valid:"int64,required"`
}

// RoomStatusPayload is the struct that defines the payload for the status route
type RoomStatusPayload struct {
	Status    string `json:"status" valid:"in(room-ready|match-started|match-ended),required"`
	Timestamp int64  `json:"timestamp" valid:"int64,required"`
}

// SchedulerPayload is the struct that defines the payload for the scheduler routes
type SchedulerPayload struct {
	Yaml string `json:"yaml" valid:"required"`
}
