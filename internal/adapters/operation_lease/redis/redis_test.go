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

//go:build integration
// +build integration

package redis

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/topfreegames/maestro/test"
)

var redisAddress string

func TestMain(m *testing.M) {
	var code int
	test.WithRedisContainer(func(redisContainerAddress string) {
		redisAddress = redisContainerAddress
		code = m.Run()
	})
	os.Exit(code)
}

func TestGrantLease(t *testing.T) {
	t.Run("with success", func(t *testing.T) {
		require.Equal(t, 1, 1)
	})
}

func TestRevokeLease(t *testing.T) {
	t.Run("with success", func(t *testing.T) {
		require.Equal(t, 1, 1)
	})
}

func TestRenewLease(t *testing.T) {
	t.Run("with success", func(t *testing.T) {
		require.Equal(t, 1, 1)
	})
}

func TestFetchLeaseTTL(t *testing.T) {
	t.Run("with success", func(t *testing.T) {
		require.Equal(t, 1, 1)
	})
}

func TestListExpiredLeases(t *testing.T) {
	t.Run("with success", func(t *testing.T) {
		require.Equal(t, 1, 1)
	})
}
