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

package remove_rooms

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/topfreegames/maestro/internal/core/entities/operation"
	"go.uber.org/zap"
)

const OperationName = "remove_rooms"

type RemoveRoomsDefinition struct {
	Amount   int      `json:"amount"`
	RoomsIDs []string `json:"rooms_ids"`
}

func (d *RemoveRoomsDefinition) ShouldExecute(_ context.Context, _ []*operation.Operation) bool {
	return true
}

func (d *RemoveRoomsDefinition) Name() string {
	return OperationName
}

func (d *RemoveRoomsDefinition) Marshal() []byte {
	bytes, err := json.Marshal(d)
	if err != nil {
		zap.L().With(zap.Error(err)).Error("error marshalling remove rooms operation definition")
		return nil
	}

	return bytes
}

func (d *RemoveRoomsDefinition) Unmarshal(raw []byte) error {
	err := json.Unmarshal(raw, d)
	if err != nil {
		return fmt.Errorf("error marshalling remove rooms operation definition: %w", err)
	}

	return nil
}
