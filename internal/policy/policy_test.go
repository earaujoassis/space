package policy

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/earaujoassis/space/internal/config"
	"github.com/earaujoassis/space/internal/gateways/redis"
	"github.com/earaujoassis/space/test/utils"
)

type PolicyTestSuite struct {
	suite.Suite
	memory *redis.MemoryService
	rls    *RateLimitService
}

func (s *PolicyTestSuite) SetupSuite() {
	utils.SetupConfigEnv()
	config, _ := config.Load()
	memory, err := redis.NewMemoryService(config)
	s.Require().NoError(err)
	if err != nil {
		s.T().Fatalf("Could not create new memory service: %v", err)
	}
	s.rls = NewRateLimitService(memory)
	s.memory = memory
}

func (s *PolicyTestSuite) SetupTest() {
	s.cleanupRedis()
}

func (s *PolicyTestSuite) cleanupRedis() {
	s.memory.Do("FLUSHDB")
}

func (s *PolicyTestSuite) TearDownSuite() {
	s.memory.Close()
}

func TestPolicySuite(t *testing.T) {
	suite.Run(t, new(PolicyTestSuite))
}
