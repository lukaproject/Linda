package taskqueueclient_test

import (
	"Linda/services/agentcentral/internal/config"
	"Linda/services/agentcentral/internal/logic/comm/taskqueueclient"
	"testing"

	"github.com/alicebob/miniredis/v2"
	"github.com/stretchr/testify/suite"
)

type redisClientTestSuite struct {
	suite.Suite
	rds           *miniredis.Miniredis
	rdsTestConfig *config.RedisConfig
}

func (s *redisClientTestSuite) SetupSuite() {
	s.rds = miniredis.RunT(s.T())
	s.rdsTestConfig = &config.RedisConfig{
		Addrs: []string{s.rds.Addr()},
	}
}

func (s *redisClientTestSuite) TestGeneralWorkFlow() {
	cli := taskqueueclient.NewRedisTaskQueueClient(s.rdsTestConfig)
	bagName := s.T().Name()
	enques := map[string]struct {
		taskName string
		orderId  uint32
		priority uint16
	}{
		"t1":     {"t1", 1, 10},
		"t2":     {"t2", 2, 10},
		"t3":     {"t3", 3, 10},
		"t4":     {"t4", 4, 10},
		"t5":     {"t5", 5, 10},
		"t6":     {"t6", 6, 10},
		"t7-min": {"t7-min", 7, 9},
		"t8":     {"t8", 8, 10},
		"t9":     {"t9", 9, 10},
		"t10":    {"t10", 10, 10},
	}

	for taskName, enqueItem := range enques {
		cli.Enque(taskName, bagName, enqueItem.priority, enqueItem.orderId)
	}

	last := ""
	for id := 0; id < len(enques); id++ {
		taskName, err := cli.Deque(bagName)
		s.Nil(err)
		if last != "" {
			if enques[taskName].priority == enques[last].priority {
				s.Greater(enques[taskName].orderId, enques[last].orderId)
			} else {
				s.Greater(enques[taskName].priority, enques[last].priority)
			}
			last = taskName
		}
	}

	_, err := cli.Deque(bagName)
	s.NotNil(err)
}

func TestRedisClientMain(t *testing.T) {
	suite.Run(t, new(redisClientTestSuite))
}
