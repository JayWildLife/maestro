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

package handlers

import (
	"context"
	"fmt"

	"github.com/topfreegames/maestro/internal/core/entities/operation"
	"github.com/topfreegames/maestro/internal/core/services/operation_manager"
	api "github.com/topfreegames/maestro/pkg/api/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type OperationsHandler struct {
	operationManager *operation_manager.OperationManager
	api.UnimplementedOperationsServiceServer
}

func ProvideOperationsHandler(operationManager *operation_manager.OperationManager) *OperationsHandler {
	return &OperationsHandler{
		operationManager: operationManager,
	}
}

func (h *OperationsHandler) ListOperations(ctx context.Context, request *api.ListOperationsRequest) (*api.ListOperationsResponse, error) {
	pendingOperationEntities, err := h.operationManager.ListSchedulerPendingOperations(ctx, request.GetSchedulerName())
	if err != nil {
		return nil, status.Error(codes.Unknown, err.Error())
	}

	pendingOperationResponse, err := h.fromOperationsToResponses(pendingOperationEntities)
	if err != nil {
		return nil, status.Error(codes.Unknown, err.Error())
	}

	activeOperationEntities, err := h.operationManager.ListSchedulerActiveOperations(ctx, request.GetSchedulerName())
	if err != nil {
		return nil, status.Error(codes.Unknown, err.Error())
	}

	activeOperationResponses, err := h.fromOperationsToResponses(activeOperationEntities)
	if err != nil {
		return nil, status.Error(codes.Unknown, err.Error())
	}

	finishedOperationEntities, err := h.operationManager.ListSchedulerFinishedOperations(ctx, request.GetSchedulerName())
	if err != nil {
		return nil, status.Error(codes.Unknown, err.Error())
	}

	finishedOperationResponse, err := h.fromOperationsToResponses(finishedOperationEntities)
	if err != nil {
		return nil, status.Error(codes.Unknown, err.Error())
	}

	return &api.ListOperationsResponse{
		PendingOperations:  pendingOperationResponse,
		ActiveOperations:   activeOperationResponses,
		FinishedOperations: finishedOperationResponse,
	}, nil
}

func (h *OperationsHandler) fromOperationsToResponses(entities []*operation.Operation) ([]*api.Operation, error) {
	responses := make([]*api.Operation, len(entities))
	for i, entity := range entities {
		response, err := h.fromOperationToResponse(entity)
		if err != nil {
			return nil, err
		}
		responses[i] = response
	}

	return responses, nil
}

func (h *OperationsHandler) fromOperationToResponse(entity *operation.Operation) (*api.Operation, error) {
	status, err := entity.Status.String()
	if err != nil {
		return nil, fmt.Errorf("failed to convert operation entity to response: %w", err)
	}

	return &api.Operation{
		Id:             entity.ID,
		Status:         status,
		DefinitionName: entity.DefinitionName,
		SchedulerName:  entity.SchedulerName,
		CreatedAt:      timestamppb.New(entity.CreatedAt),
	}, nil
}
