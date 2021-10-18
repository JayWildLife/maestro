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

package ports

import (
	"context"
)

type OperationCancellationRequest struct {
	SchedulerName string `json:"schedulerName"`
	OperationID   string `json:"operationID"`
}

type OperationFlow interface {
	InsertOperationID(ctx context.Context, schedulerName, operationID string) error
	// NextOperationID fetches the next scheduler operation to be
	// processed and return its ID.
	NextOperationID(ctx context.Context, schedulerName string) (string, error)
	// ListSchedulerPendingOperationIDs list scheduler pending operation IDs.
	ListSchedulerPendingOperationIDs(ctx context.Context, schedulerName string) ([]string, error)
	// EnqueueOperationCancellationRequest enqueue a operation cancellation request
	EnqueueOperationCancellationRequest(ctx context.Context, request OperationCancellationRequest) error
	// WatchOperationCancellationRequests watches for operation cancellation requests
	WatchOperationCancellationRequests(ctx context.Context) chan OperationCancellationRequest
}
