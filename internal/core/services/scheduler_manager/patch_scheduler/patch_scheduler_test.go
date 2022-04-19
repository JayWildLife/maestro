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

package patch_scheduler_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/topfreegames/maestro/internal/core/entities"
	"github.com/topfreegames/maestro/internal/core/entities/forwarder"
	"github.com/topfreegames/maestro/internal/core/entities/game_room"
	"github.com/topfreegames/maestro/internal/core/services/scheduler_manager/patch_scheduler"
)

func TestPatchScheduler(t *testing.T) {
	type Input struct {
		Scheduler *entities.Scheduler
		PatchMap  map[string]interface{}
	}

	type Output struct {
		ChangeSchedulerFunc func() *entities.Scheduler
		Error               error
	}

	testCases := []struct {
		Title string
		Input
		Output
	}{
		// Scheduler success cases

		{
			Title: "Have max surge return scheduler with changed MaxSurge",
			Input: Input{
				Scheduler: basicSchedulerToPatchSchedulerTests(),
				PatchMap: map[string]interface{}{
					patch_scheduler.LabelSchedulerMaxSurge: "30%",
				},
			},
			Output: Output{
				ChangeSchedulerFunc: func() *entities.Scheduler {
					scheduler := basicSchedulerToPatchSchedulerTests()
					scheduler.MaxSurge = "30%"

					return scheduler
				},
				Error: nil,
			},
		},
		{
			Title: "Have port range return scheduler with changed PortRange",
			Input: Input{
				Scheduler: basicSchedulerToPatchSchedulerTests(),
				PatchMap: map[string]interface{}{
					patch_scheduler.LabelSchedulerPortRange: &entities.PortRange{
						Start: 10000,
						End:   60000,
					},
				},
			},
			Output: Output{
				ChangeSchedulerFunc: func() *entities.Scheduler {
					scheduler := basicSchedulerToPatchSchedulerTests()
					scheduler.PortRange = &entities.PortRange{
						Start: 10000,
						End:   60000,
					}

					return scheduler
				},
				Error: nil,
			},
		},
		{
			Title: " Have forwarders return scheduler with changed Forwarders",
			Input: Input{
				Scheduler: basicSchedulerToPatchSchedulerTests(),
				PatchMap: map[string]interface{}{
					patch_scheduler.LabelSchedulerForwarders: []*forwarder.Forwarder{
						{
							Name:        "first-forwarder",
							Enabled:     false,
							ForwardType: forwarder.TypeGrpc,
							Address:     "localhost:2000",
						},
					},
				},
			},
			Output: Output{
				ChangeSchedulerFunc: func() *entities.Scheduler {
					scheduler := basicSchedulerToPatchSchedulerTests()
					scheduler.Forwarders = []*forwarder.Forwarder{
						{
							Name:        "first-forwarder",
							Enabled:     false,
							ForwardType: forwarder.TypeGrpc,
							Address:     "localhost:2000",
						},
					}
					return scheduler
				},
				Error: nil,
			},
		},

		// Scheduler Errors

		{
			Title: "Have port range wrong return error",
			Input: Input{
				Scheduler: basicSchedulerToPatchSchedulerTests(),
				PatchMap: map[string]interface{}{
					patch_scheduler.LabelSchedulerPortRange: "wrong port range",
				},
			},
			Output: Output{
				ChangeSchedulerFunc: func() *entities.Scheduler {
					return basicSchedulerToPatchSchedulerTests()
				},
				Error: fmt.Errorf("error parsing scheduler: port range malformed"),
			},
		},
		{
			Title: " Have wrong forwarders return error",
			Input: Input{
				Scheduler: basicSchedulerToPatchSchedulerTests(),
				PatchMap: map[string]interface{}{
					patch_scheduler.LabelSchedulerForwarders: "wrong forwarders",
				},
			},
			Output: Output{
				ChangeSchedulerFunc: func() *entities.Scheduler {
					return basicSchedulerToPatchSchedulerTests()
				},
				Error: fmt.Errorf("error parsing scheduler: forwarders malformed"),
			},
		},

		// Spec Success cases

		{
			Title: "Have termination grace period return scheduler with changed TerminationGracePeriod",
			Input: Input{
				Scheduler: basicSchedulerToPatchSchedulerTests(),
				PatchMap: map[string]interface{}{
					patch_scheduler.LabelSchedulerSpec: map[string]interface{}{
						patch_scheduler.LabelSpecTerminationGracePeriod: time.Duration(12),
					},
				},
			},
			Output: Output{
				ChangeSchedulerFunc: func() *entities.Scheduler {
					scheduler := basicSchedulerToPatchSchedulerTests()
					scheduler.Spec.TerminationGracePeriod = time.Duration(12)

					return scheduler
				},
				Error: nil,
			},
		},
		{
			Title: "Have Toleration return scheduler with changed Toleration",
			Input: Input{
				Scheduler: basicSchedulerToPatchSchedulerTests(),
				PatchMap: map[string]interface{}{
					patch_scheduler.LabelSchedulerSpec: map[string]interface{}{
						patch_scheduler.LabelSpecToleration: "toleration",
					},
				},
			},
			Output: Output{
				ChangeSchedulerFunc: func() *entities.Scheduler {
					scheduler := basicSchedulerToPatchSchedulerTests()
					scheduler.Spec.Toleration = "toleration"

					return scheduler
				},
				Error: nil,
			},
		},
		{
			Title: "Have affinity return scheduler with changed Affinity",
			Input: Input{
				Scheduler: basicSchedulerToPatchSchedulerTests(),
				PatchMap: map[string]interface{}{
					patch_scheduler.LabelSchedulerSpec: map[string]interface{}{
						patch_scheduler.LabelSpecAffinity: "affinity",
					},
				},
			},
			Output: Output{
				ChangeSchedulerFunc: func() *entities.Scheduler {
					scheduler := basicSchedulerToPatchSchedulerTests()
					scheduler.Spec.Affinity = "affinity"

					return scheduler
				},
				Error: nil,
			},
		},

		// Spec errors

		{
			Title: "Have termination grace period return scheduler with changed TerminationGracePeriod",
			Input: Input{
				Scheduler: basicSchedulerToPatchSchedulerTests(),
				PatchMap: map[string]interface{}{
					patch_scheduler.LabelSchedulerSpec: map[string]interface{}{
						patch_scheduler.LabelSpecTerminationGracePeriod: "wrong termination grace period",
					},
				},
			},
			Output: Output{
				ChangeSchedulerFunc: func() *entities.Scheduler {
					return basicSchedulerToPatchSchedulerTests()
				},
				Error: fmt.Errorf("error parsing scheduler: error parsing spec: termination grace period malformed"),
			},
		},

		// Containers success

		{
			Title: "Have name return scheduler with changed Name",
			Input: Input{
				Scheduler: basicSchedulerToPatchSchedulerTests(),
				PatchMap: map[string]interface{}{
					patch_scheduler.LabelSchedulerSpec: map[string]interface{}{
						patch_scheduler.LabelSpecContainers: []map[string]interface{}{
							map[string]interface{}{
								patch_scheduler.LabelContainerName: "some-name",
							},
						},
					},
				},
			},
			Output: Output{
				ChangeSchedulerFunc: func() *entities.Scheduler {
					scheduler := basicSchedulerToPatchSchedulerTests()
					scheduler.Spec.Containers[0].Name = "some-name"

					return scheduler
				},
				Error: nil,
			},
		},
		{
			Title: "Have image return scheduler with changed Image",
			Input: Input{
				Scheduler: basicSchedulerToPatchSchedulerTests(),
				PatchMap: map[string]interface{}{
					patch_scheduler.LabelSchedulerSpec: map[string]interface{}{
						patch_scheduler.LabelSpecContainers: []map[string]interface{}{
							map[string]interface{}{
								patch_scheduler.LabelContainerImage: "some-image",
							},
						},
					},
				},
			},
			Output: Output{
				ChangeSchedulerFunc: func() *entities.Scheduler {
					scheduler := basicSchedulerToPatchSchedulerTests()
					scheduler.Spec.Containers[0].Image = "some-image"

					return scheduler
				},
				Error: nil,
			},
		},
		{
			Title: "Have ImagePullPolicy return scheduler with changed ImagePullPolicy",
			Input: Input{
				Scheduler: basicSchedulerToPatchSchedulerTests(),
				PatchMap: map[string]interface{}{
					patch_scheduler.LabelSchedulerSpec: map[string]interface{}{
						patch_scheduler.LabelSpecContainers: []map[string]interface{}{
							map[string]interface{}{
								patch_scheduler.LabelContainerImagePullPolicy: "LocalFirst",
							},
						},
					},
				},
			},
			Output: Output{
				ChangeSchedulerFunc: func() *entities.Scheduler {
					scheduler := basicSchedulerToPatchSchedulerTests()
					scheduler.Spec.Containers[0].ImagePullPolicy = "LocalFirst"

					return scheduler
				},
				Error: nil,
			},
		},
		{
			Title: "Have Command return scheduler with changed Command",
			Input: Input{
				Scheduler: basicSchedulerToPatchSchedulerTests(),
				PatchMap: map[string]interface{}{
					patch_scheduler.LabelSchedulerSpec: map[string]interface{}{
						patch_scheduler.LabelSpecContainers: []map[string]interface{}{
							map[string]interface{}{
								patch_scheduler.LabelContainerCommand: []string{"/bin/sh", "command"},
							},
						},
					},
				},
			},
			Output: Output{
				ChangeSchedulerFunc: func() *entities.Scheduler {
					scheduler := basicSchedulerToPatchSchedulerTests()
					scheduler.Spec.Containers[0].Command = []string{"/bin/sh", "command"}

					return scheduler
				},
				Error: nil,
			},
		},
		{
			Title: "Have Environment return scheduler with changed Environment",
			Input: Input{
				Scheduler: basicSchedulerToPatchSchedulerTests(),
				PatchMap: map[string]interface{}{
					patch_scheduler.LabelSchedulerSpec: map[string]interface{}{
						patch_scheduler.LabelSpecContainers: []map[string]interface{}{
							map[string]interface{}{
								patch_scheduler.LabelContainerEnvironment: []game_room.ContainerEnvironment{
									{
										Name:  "some-name",
										Value: "some-value",
									},
								},
							},
						},
					},
				},
			},
			Output: Output{
				ChangeSchedulerFunc: func() *entities.Scheduler {
					scheduler := basicSchedulerToPatchSchedulerTests()
					scheduler.Spec.Containers[0].Environment = []game_room.ContainerEnvironment{
						{
							Name:  "some-name",
							Value: "some-value",
						},
					}

					return scheduler
				},
				Error: nil,
			},
		},
		{
			Title: "Have Requests return scheduler with changed Requests",
			Input: Input{
				Scheduler: basicSchedulerToPatchSchedulerTests(),
				PatchMap: map[string]interface{}{
					patch_scheduler.LabelSchedulerSpec: map[string]interface{}{
						patch_scheduler.LabelSpecContainers: []map[string]interface{}{
							map[string]interface{}{
								patch_scheduler.LabelContainerRequests: game_room.ContainerResources{
									Memory: "1000",
									CPU:    "500m",
								},
							},
						},
					},
				},
			},
			Output: Output{
				ChangeSchedulerFunc: func() *entities.Scheduler {
					scheduler := basicSchedulerToPatchSchedulerTests()
					scheduler.Spec.Containers[0].Requests = game_room.ContainerResources{
						Memory: "1000",
						CPU:    "500m",
					}

					return scheduler
				},
				Error: nil,
			},
		},
		{
			Title: "Have Limits return scheduler with changed Limits",
			Input: Input{
				Scheduler: basicSchedulerToPatchSchedulerTests(),
				PatchMap: map[string]interface{}{
					patch_scheduler.LabelSchedulerSpec: map[string]interface{}{
						patch_scheduler.LabelSpecContainers: []map[string]interface{}{
							map[string]interface{}{
								patch_scheduler.LabelContainerLimits: game_room.ContainerResources{
									Memory: "1000Mi",
									CPU:    "500m",
								},
							},
						},
					},
				},
			},
			Output: Output{
				ChangeSchedulerFunc: func() *entities.Scheduler {
					scheduler := basicSchedulerToPatchSchedulerTests()
					scheduler.Spec.Containers[0].Limits = game_room.ContainerResources{
						Memory: "1000Mi",
						CPU:    "500m",
					}

					return scheduler
				},
				Error: nil,
			},
		},
		{
			Title: "Have Ports return scheduler with changed Ports",
			Input: Input{
				Scheduler: basicSchedulerToPatchSchedulerTests(),
				PatchMap: map[string]interface{}{
					patch_scheduler.LabelSchedulerSpec: map[string]interface{}{
						patch_scheduler.LabelSpecContainers: []map[string]interface{}{
							map[string]interface{}{
								patch_scheduler.LabelContainerPorts: []game_room.ContainerPort{
									{
										Name:     "some-port",
										Protocol: "TCP",
										Port:     2048,
									},
								},
							},
						},
					},
				},
			},
			Output: Output{
				ChangeSchedulerFunc: func() *entities.Scheduler {
					scheduler := basicSchedulerToPatchSchedulerTests()
					scheduler.Spec.Containers[0].Ports = []game_room.ContainerPort{
						{
							Name:     "some-port",
							Protocol: "TCP",
							Port:     2048,
						},
					}

					return scheduler
				},
				Error: nil,
			},
		},

		{
			Title: "Creating a new container return scheduler adding that container",
			Input: Input{
				Scheduler: basicSchedulerToPatchSchedulerTests(),
				PatchMap: map[string]interface{}{
					patch_scheduler.LabelSchedulerSpec: map[string]interface{}{
						patch_scheduler.LabelSpecContainers: []map[string]interface{}{
							map[string]interface{}{},
							map[string]interface{}{
								patch_scheduler.LabelContainerPorts: []game_room.ContainerPort{
									{
										Name:     "some-port",
										Protocol: "TCP",
										Port:     2048,
									},
								},
							},
						},
					},
				},
			},
			Output: Output{
				ChangeSchedulerFunc: func() *entities.Scheduler {
					scheduler := basicSchedulerToPatchSchedulerTests()
					scheduler.Spec.Containers = append(scheduler.Spec.Containers, game_room.Container{
						Ports: []game_room.ContainerPort{
							{
								Name:     "some-port",
								Protocol: "TCP",
								Port:     2048,
							},
						},
					})

					return scheduler
				},
				Error: nil,
			},
		},

		// Containers errors

		{
			Title: "Have wrong Command return error",
			Input: Input{
				Scheduler: basicSchedulerToPatchSchedulerTests(),
				PatchMap: map[string]interface{}{
					patch_scheduler.LabelSchedulerSpec: map[string]interface{}{
						patch_scheduler.LabelSpecContainers: []map[string]interface{}{
							map[string]interface{}{
								patch_scheduler.LabelContainerCommand: "wrong-command",
							},
						},
					},
				},
			},
			Output: Output{
				ChangeSchedulerFunc: func() *entities.Scheduler {
					return basicSchedulerToPatchSchedulerTests()
				},
				Error: fmt.Errorf("error parsing scheduler: error parsing spec: error parsing containers: command malformed"),
			},
		},
		{
			Title: "Have wrong Environment return error",
			Input: Input{
				Scheduler: basicSchedulerToPatchSchedulerTests(),
				PatchMap: map[string]interface{}{
					patch_scheduler.LabelSchedulerSpec: map[string]interface{}{
						patch_scheduler.LabelSpecContainers: []map[string]interface{}{
							map[string]interface{}{
								patch_scheduler.LabelContainerEnvironment: "wrong-environment",
							},
						},
					},
				},
			},
			Output: Output{
				ChangeSchedulerFunc: func() *entities.Scheduler {
					return basicSchedulerToPatchSchedulerTests()
				},
				Error: fmt.Errorf("error parsing scheduler: error parsing spec: error parsing containers: environment malformed"),
			},
		},
		{
			Title: "Have wrong Requests return error",
			Input: Input{
				Scheduler: basicSchedulerToPatchSchedulerTests(),
				PatchMap: map[string]interface{}{
					patch_scheduler.LabelSchedulerSpec: map[string]interface{}{
						patch_scheduler.LabelSpecContainers: []map[string]interface{}{
							map[string]interface{}{
								patch_scheduler.LabelContainerRequests: "wrong-requests",
							},
						},
					},
				},
			},
			Output: Output{
				ChangeSchedulerFunc: func() *entities.Scheduler {
					return basicSchedulerToPatchSchedulerTests()
				},
				Error: fmt.Errorf("error parsing scheduler: error parsing spec: error parsing containers: requests malformed"),
			},
		},
		{
			Title: "Have wrong Limits return error",
			Input: Input{
				Scheduler: basicSchedulerToPatchSchedulerTests(),
				PatchMap: map[string]interface{}{
					patch_scheduler.LabelSchedulerSpec: map[string]interface{}{
						patch_scheduler.LabelSpecContainers: []map[string]interface{}{
							map[string]interface{}{
								patch_scheduler.LabelContainerLimits: "wrong-limits",
							},
						},
					},
				},
			},
			Output: Output{
				ChangeSchedulerFunc: func() *entities.Scheduler {
					return basicSchedulerToPatchSchedulerTests()
				},
				Error: fmt.Errorf("error parsing scheduler: error parsing spec: error parsing containers: limits malformed"),
			},
		},
		{
			Title: "Have wrong Ports return error",
			Input: Input{
				Scheduler: basicSchedulerToPatchSchedulerTests(),
				PatchMap: map[string]interface{}{
					patch_scheduler.LabelSchedulerSpec: map[string]interface{}{
						patch_scheduler.LabelSpecContainers: []map[string]interface{}{
							map[string]interface{}{
								patch_scheduler.LabelContainerPorts: "wrong-ports",
							},
						},
					},
				},
			},
			Output: Output{
				ChangeSchedulerFunc: func() *entities.Scheduler {
					return basicSchedulerToPatchSchedulerTests()
				},
				Error: fmt.Errorf("error parsing scheduler: error parsing spec: error parsing containers: ports malformed"),
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.Title, func(t *testing.T) {
			expectedScheduler := testCase.Output.ChangeSchedulerFunc()
			scheduler, err := patch_scheduler.PatchScheduler(*testCase.Input.Scheduler, testCase.Input.PatchMap)

			if testCase.Output.Error != nil {
				assert.EqualError(t, err, testCase.Output.Error.Error())
				return
			}

			assert.Equal(t, expectedScheduler.MaxSurge, scheduler.MaxSurge)
			assert.EqualValues(t, expectedScheduler.PortRange, scheduler.PortRange)
			forwarders := expectedScheduler.Forwarders
			for i, expectedForwarder := range expectedScheduler.Forwarders {
				assert.EqualValues(t, expectedForwarder, forwarders[i])
			}

			expectedSpec := expectedScheduler.Spec
			spec := scheduler.Spec

			assert.Equal(t, expectedSpec.TerminationGracePeriod, spec.TerminationGracePeriod)
			assert.Equal(t, expectedSpec.Toleration, spec.Toleration)
			assert.Equal(t, expectedSpec.Affinity, spec.Affinity)

			containers := spec.Containers
			for i, expectedContainer := range expectedSpec.Containers {
				assert.Equal(t, expectedContainer.Image, containers[i].Image)
				assert.Equal(t, expectedContainer.ImagePullPolicy, containers[i].ImagePullPolicy)
				assert.EqualValues(t, expectedContainer.Command, containers[i].Command)
				assert.Equal(t, expectedContainer.ImagePullPolicy, containers[i].ImagePullPolicy)

				environment := containers[i].Environment
				for j, expectedEnvironment := range expectedContainer.Environment {
					assert.EqualValues(t, expectedEnvironment, environment[j])
				}
				assert.EqualValues(t, expectedContainer.Requests, containers[i].Requests)
				assert.EqualValues(t, expectedContainer.Limits, containers[i].Limits)

				containerPorts := containers[i].Ports
				for j, expectedPort := range expectedContainer.Ports {
					assert.EqualValues(t, expectedPort, containerPorts[j])
				}
			}
		})
	}
}

// basicSchedulerToPatchSchedulerTests generates a valid scheduler with the required fields.
func basicSchedulerToPatchSchedulerTests() *entities.Scheduler {
	return &entities.Scheduler{
		Name:            "scheduler",
		Game:            "game",
		State:           entities.StateCreating,
		MaxSurge:        "10%",
		RollbackVersion: "v1.0.0",
		Spec: game_room.Spec{
			Version:                "v1.1.0",
			TerminationGracePeriod: 60,
			Toleration:             "toleration",
			Affinity:               "affinity",
			Containers: []game_room.Container{
				{
					Name:            "default",
					Image:           "some-image",
					ImagePullPolicy: "Always",
					Command:         []string{"hello"},
					Ports: []game_room.ContainerPort{
						{Name: "tcp", Protocol: "tcp", Port: 80},
					},
					Requests: game_room.ContainerResources{
						CPU:    "10m",
						Memory: "100Mi",
					},
					Limits: game_room.ContainerResources{
						CPU:    "10m",
						Memory: "100Mi",
					},
				},
			},
		},
		PortRange: &entities.PortRange{
			Start: 40000,
			End:   60000,
		},
	}
}