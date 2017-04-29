// maestro
// https://github.com/topfreegames/maestro
//
// Licensed under the MIT license:
// http://www.opensource.org/licenses/mit-license
// Copyright © 2017 Top Free Games <backend@tfgco.com>

package models_test

import (
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"

	pgmocks "github.com/topfreegames/extensions/pg/mocks"
	redismocks "github.com/topfreegames/extensions/redis/mocks"
)

var (
	db          *pgmocks.PGMock
	redisClient *redismocks.MockRedisClient
	mockCtrl    *gomock.Controller
	err         error
)

func TestModels(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Models Suite")
}

var _ = BeforeEach(func() {
	db = pgmocks.NewPGMock(1, 1)
	mockCtrl = gomock.NewController(GinkgoT())
	redisClient = redismocks.NewMockRedisClient(mockCtrl)
})

var _ = AfterEach(func() {
	defer db.Close()
})
