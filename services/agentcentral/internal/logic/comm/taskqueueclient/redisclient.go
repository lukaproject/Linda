package taskqueueclient

import (
	"Linda/baselibs/codes/errno"
	"Linda/services/agentcentral/internal/config"
	"context"

	"github.com/redis/go-redis/v9"
)

var _ Client = &redisTaskQueueClient{}

type redisTaskQueueClient struct {
	rc redis.UniversalClient
}

func NewRedisTaskQueueClient(config *config.RedisConfig) Client {
	return NewRedisTaskQueueClientWithUniversalRedisClient(
		redis.NewUniversalClient(&redis.UniversalOptions{
			Addrs:    config.Addrs,
			Password: config.Password,
			DB:       config.Db,
		}),
	)
}

func NewRedisTaskQueueClientWithUniversalRedisClient(rc redis.UniversalClient) Client {
	return &redisTaskQueueClient{
		rc: rc,
	}
}

func (rdsClient *redisTaskQueueClient) Enque(taskName string, bagName string, priority uint16, orderId uint32) (err error) {
	priorityValue := float64(priority)*(1e10) + float64(orderId)
	rdsClient.rc.ZAdd(context.Background(), bagName, redis.Z{
		Score:  priorityValue,
		Member: taskName,
	})
	return
}

func (rdsClient *redisTaskQueueClient) Deque(bagName string) (taskName string, err error) {
	cmd := rdsClient.rc.ZPopMin(context.Background(), bagName, 1)
	if cmd.Err() != nil {
		err = cmd.Err()
		return
	}
	res, err := cmd.Result()
	if err != nil {
		return
	}
	if len(res) == 0 {
		err = errno.ErrEmptyBag
		return
	}
	taskName = res[0].Member.(string)
	return
}
