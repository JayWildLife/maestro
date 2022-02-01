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

package management

import (
	"context"
	"fmt"
	"testing"
	"time"

	_struct "github.com/golang/protobuf/ptypes/struct"
	"google.golang.org/protobuf/types/known/structpb"

	"github.com/go-redis/redis/v8"
	"github.com/stretchr/testify/require"

	"github.com/topfreegames/maestro/e2e/framework/maestro"

	"github.com/topfreegames/maestro/e2e/framework"
	maestroApiV1 "github.com/topfreegames/maestro/pkg/api/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func TestEventsForwarding(t *testing.T) {
	framework.WithClients(t, func(roomsApiClient *framework.APIClient, managementApiClient *framework.APIClient, kubeClient kubernetes.Interface, redisClient *redis.Client, maestro *maestro.MaestroInstance) {

		t.Run("[Player event success] Forward player event return success true when no error occurs while forwarding events call", func(t *testing.T) {
			t.Parallel()

			schedulerName, roomName := createSchedulerWithForwarderAndRooms(t, maestro, kubeClient, managementApiClient, maestro.ServerMocks.GrpcForwarderAddress)

			// This configuration make the grpc service return with success
			err := addStubRequestToMockedGrpcServer("events-forwarder-grpc-send-player-event-success")
			require.NoError(t, err)

			playerEventRequest := &maestroApiV1.ForwardPlayerEventRequest{
				RoomName: roomName,

				Event:     "playerLeft",
				Timestamp: time.Now().Unix(),
				Metadata: &_struct.Struct{
					Fields: map[string]*structpb.Value{
						"playerId": {
							Kind: &structpb.Value_StringValue{
								StringValue: "c50acc91-4d88-46fa-aa56-48d63c5b5311",
							},
						},
						"eventMetadata1": {
							Kind: &structpb.Value_StringValue{
								StringValue: "value1",
							},
						},
						"eventMetadata2": {
							Kind: &structpb.Value_BoolValue{
								BoolValue: true,
							},
						},
					},
				},
			}
			playerEventResponse := &maestroApiV1.ForwardPlayerEventResponse{}
			err = roomsApiClient.Do("POST", fmt.Sprintf("/scheduler/%s/rooms/%s/playerevent", schedulerName, roomName), playerEventRequest, playerEventResponse)
			require.NoError(t, err)
			require.Equal(t, true, playerEventResponse.Success)
		})

		t.Run("[Player event success] Forward player event return success true when no forwarder is configured for the scheduler", func(t *testing.T) {
			t.Parallel()

			schedulerName, roomName := createSchedulerWithRooms(t, maestro, kubeClient, managementApiClient)

			playerEventRequest := &maestroApiV1.ForwardPlayerEventRequest{
				RoomName:  roomName,
				Event:     "playerLeft",
				Timestamp: time.Now().Unix(),
				Metadata: &_struct.Struct{
					Fields: map[string]*structpb.Value{
						"playerId": {
							Kind: &structpb.Value_StringValue{
								StringValue: "5280087d-6dff-4bbf-abc8-45cb8786ad00",
							},
						},
					},
				},
			}
			playerEventResponse := &maestroApiV1.ForwardPlayerEventResponse{}
			err := roomsApiClient.Do("POST", fmt.Sprintf("/scheduler/%s/rooms/%s/playerevent", schedulerName, roomName), playerEventRequest, playerEventResponse)
			require.NoError(t, err)
			require.Equal(t, true, playerEventResponse.Success)
		})

		t.Run("[Player event failure] Forward player event return success false when some error occurs in GRPC call", func(t *testing.T) {
			t.Parallel()

			schedulerName, roomName := createSchedulerWithForwarderAndRooms(t, maestro, kubeClient, managementApiClient, maestro.ServerMocks.GrpcForwarderAddress)

			// This configuration make the grpc service return with error
			err := addStubRequestToMockedGrpcServer("events-forwarder-grpc-send-player-event-failure")

			playerEventRequest := &maestroApiV1.ForwardPlayerEventRequest{
				RoomName: roomName,

				Event:     "playerLeft",
				Timestamp: time.Now().Unix(),
				Metadata: &_struct.Struct{

					Fields: map[string]*structpb.Value{
						"playerId": {
							Kind: &structpb.Value_StringValue{
								StringValue: "446bb3d0-0334-4468-a4e7-8068a97caa53",
							},
						},
						"eventMetadata1": {
							Kind: &structpb.Value_StringValue{
								StringValue: "value1",
							},
						},
						"eventMetadata2": {
							Kind: &structpb.Value_BoolValue{
								BoolValue: true,
							},
						},
					},
				},
			}
			playerEventResponse := &maestroApiV1.ForwardPlayerEventResponse{}
			err = roomsApiClient.Do("POST", fmt.Sprintf("/scheduler/%s/rooms/%s/playerevent", schedulerName, roomName), playerEventRequest, playerEventResponse)
			require.NoError(t, err)
			require.Equal(t, false, playerEventResponse.Success)
			require.Equal(t, "failed to forward event room at \"matchmaker-grpc\"", playerEventResponse.Message)
		})

		t.Run("[Player event failure] Forward player event return success false when forwarding event for an inexistent room", func(t *testing.T) {
			t.Parallel()

			schedulerName, _ := createSchedulerWithForwarderAndRooms(t, maestro, kubeClient, managementApiClient, maestro.ServerMocks.GrpcForwarderAddress)
			roomName := "inexistent-room"

			err := addStubRequestToMockedGrpcServer("events-forwarder-grpc-send-player-event-failure")

			playerEventRequest := &maestroApiV1.ForwardPlayerEventRequest{
				RoomName:  roomName,
				Event:     "playerLeft",
				Timestamp: time.Now().Unix(),
				Metadata: &_struct.Struct{
					Fields: map[string]*structpb.Value{
						"playerId": {
							Kind: &structpb.Value_StringValue{
								StringValue: "c50c9d8a-5a40-40ee-97ea-d477d7b0abd9",
							},
						},
					},
				},
			}
			playerEventResponse := &maestroApiV1.ForwardPlayerEventResponse{}
			err = roomsApiClient.Do("POST", fmt.Sprintf("/scheduler/%s/rooms/%s/playerevent", schedulerName, roomName), playerEventRequest, playerEventResponse)
			require.NoError(t, err)
			require.Equal(t, false, playerEventResponse.Success)
		})

		t.Run("[Player event failure] Forward player event return success false when the forwarder connection can't be established", func(t *testing.T) {
			t.Parallel()

			schedulerName, roomName := createSchedulerWithForwarderAndRooms(t, maestro, kubeClient, managementApiClient, "invalid-grpc-address:9982")

			playerEventRequest := &maestroApiV1.ForwardPlayerEventRequest{
				RoomName:  roomName,
				Event:     "playerLeft",
				Timestamp: time.Now().Unix(),
				Metadata: &_struct.Struct{
					Fields: map[string]*structpb.Value{
						"playerId": {
							Kind: &structpb.Value_StringValue{
								StringValue: "c50c9d8a-5a40-40ee-97ea-d477d7b0abd9",
							},
						},
					},
				},
			}
			playerEventResponse := &maestroApiV1.ForwardPlayerEventResponse{}
			err := roomsApiClient.Do("POST", fmt.Sprintf("/scheduler/%s/rooms/%s/playerevent", schedulerName, roomName), playerEventRequest, playerEventResponse)
			require.NoError(t, err)
			require.Equal(t, false, playerEventResponse.Success)
			require.Contains(t, playerEventResponse.Message, "transport: Error while dialing dial tcp: lookup invalid-grpc-address")
		})
	})

}

func createSchedulerWithRooms(t *testing.T, maestro *maestro.MaestroInstance, kubeClient kubernetes.Interface, managementApiClient *framework.APIClient) (string, string) {
	schedulerName, err := createSchedulerAndWaitForIt(t,
		maestro,
		managementApiClient,
		kubeClient,
		[]string{"/bin/sh", "-c", "apk add curl && curl --request POST " +
			"$ROOMS_API_ADDRESS:9097/scheduler/$MAESTRO_SCHEDULER_NAME/rooms/$MAESTRO_ROOM_ID/ping " +
			"--data-raw '{\"status\": \"ready\",\"timestamp\": \"12312312313\"}'"},
	)

	addRoomsRequest := &maestroApiV1.AddRoomsRequest{SchedulerName: schedulerName, Amount: 1}
	addRoomsResponse := &maestroApiV1.AddRoomsResponse{}
	err = managementApiClient.Do("POST", fmt.Sprintf("/schedulers/%s/add-rooms", schedulerName), addRoomsRequest, addRoomsResponse)

	require.Eventually(t, func() bool {
		listOperationsRequest := &maestroApiV1.ListOperationsRequest{}
		listOperationsResponse := &maestroApiV1.ListOperationsResponse{}
		err = managementApiClient.Do("GET", fmt.Sprintf("/schedulers/%s/operations", schedulerName), listOperationsRequest, listOperationsResponse)
		require.NoError(t, err)

		if len(listOperationsResponse.FinishedOperations) < 2 {
			return false
		}

		require.Equal(t, "add_rooms", listOperationsResponse.FinishedOperations[0].DefinitionName)
		return true
	}, 240*time.Second, time.Second)

	pods, err := kubeClient.CoreV1().Pods(schedulerName).List(context.Background(), metav1.ListOptions{})
	require.NoError(t, err)
	require.NotEmpty(t, pods.Items)

	return schedulerName, pods.Items[0].Name
}

func createSchedulerWithForwarderAndRooms(t *testing.T, maestro *maestro.MaestroInstance, kubeClient kubernetes.Interface, managementApiClient *framework.APIClient, forwarderAddress string) (string, string) {
	forwarders := []*maestroApiV1.Forwarder{
		{
			Name:    "matchmaker-grpc",
			Enable:  true,
			Type:    "grpc",
			Address: forwarderAddress,
			Options: &maestroApiV1.ForwarderOptions{
				Timeout: 5000,
				Metadata: &_struct.Struct{
					Fields: map[string]*structpb.Value{
						"roomType": {
							Kind: &structpb.Value_StringValue{
								StringValue: "green",
							},
						},
						"forwarderMetadata1": {
							Kind: &structpb.Value_StringValue{
								StringValue: "value1",
							},
						},
						"forwarderMetadata2": {
							Kind: &structpb.Value_NumberValue{
								NumberValue: 245,
							},
						},
					},
				},
			},
		},
	}

	schedulerName, err := createSchedulerWithForwardersAndWaitForIt(
		t,
		maestro,
		managementApiClient,
		kubeClient,
		[]string{"/bin/sh", "-c", "apk add curl && curl --request POST " +
			"$ROOMS_API_ADDRESS:9097/scheduler/$MAESTRO_SCHEDULER_NAME/rooms/$MAESTRO_ROOM_ID/ping " +
			"--data-raw '{\"status\": \"ready\",\"timestamp\": \"12312312313\"}'"},
		forwarders,
	)

	addRoomsRequest := &maestroApiV1.AddRoomsRequest{SchedulerName: schedulerName, Amount: 1}
	addRoomsResponse := &maestroApiV1.AddRoomsResponse{}
	err = managementApiClient.Do("POST", fmt.Sprintf("/schedulers/%s/add-rooms", schedulerName), addRoomsRequest, addRoomsResponse)

	require.Eventually(t, func() bool {
		listOperationsRequest := &maestroApiV1.ListOperationsRequest{}
		listOperationsResponse := &maestroApiV1.ListOperationsResponse{}
		err = managementApiClient.Do("GET", fmt.Sprintf("/schedulers/%s/operations", schedulerName), listOperationsRequest, listOperationsResponse)
		require.NoError(t, err)

		if len(listOperationsResponse.FinishedOperations) < 2 {
			return false
		}

		require.Equal(t, "add_rooms", listOperationsResponse.FinishedOperations[0].DefinitionName)
		return true
	}, 240*time.Second, time.Second)

	pods, err := kubeClient.CoreV1().Pods(schedulerName).List(context.Background(), metav1.ListOptions{})
	require.NoError(t, err)
	require.NotEmpty(t, pods.Items)

	return schedulerName, pods.Items[0].Name
}
