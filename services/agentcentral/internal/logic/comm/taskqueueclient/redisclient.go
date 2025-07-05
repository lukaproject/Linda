package taskqueueclient

import (
	"Linda/baselibs/codes/errno"
	"Linda/services/agentcentral/internal/config"
	"context"

	"github.com/redis/go-redis/v9"
)

var _ QuesManageClient = &redisQuesManageClient{}

func NewRedisQuesManageClient(config *config.RedisConfig) QuesManageClient {
	opts := &redis.UniversalOptions{
		Addrs:    config.Addrs,
		Password: config.Password,
		DB:       config.Db,
	}
	rc := redis.NewUniversalClient(opts)
	return &redisQuesManageClient{
		rc:      rc,
		options: opts,
	}
}

type redisQuesManageClient struct {
	rc      redis.UniversalClient
	options *redis.UniversalOptions
}

func (mcli *redisQuesManageClient) Create(bagName string) (err error) {
	return nil
}

func (mcli *redisQuesManageClient) Delete(bagName string) (err error) {
	return mcli.rc.Del(context.Background(), bagName).Err()
}

func (mcli *redisQuesManageClient) Get(bagName string) (queClient QueClient, err error) {
	return &redisQueClient{
		bagName: bagName,
		rc:      redis.NewUniversalClient(mcli.options),
	}, nil
}

type redisQueClient struct {
	bagName string
	rc      redis.UniversalClient
}

func (rqcli *redisQueClient) Enque(taskName string, priority uint16, orderId uint32) (err error) {
	priorityValue := float64(priority)*(1e10) + float64(orderId)
	rqcli.rc.ZAdd(context.Background(), rqcli.bagName, redis.Z{
		Score:  priorityValue,
		Member: taskName,
	})
	return
}

func (rqcli *redisQueClient) Deque() (taskName string, err error) {
	cmd := rqcli.rc.ZPopMin(context.Background(), rqcli.bagName, 1)
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

func (rqcli *redisQueClient) Deques(count int64) (taskNames []string, err error) {
	cmd := rqcli.rc.ZPopMin(context.Background(), rqcli.bagName, count)
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
	taskNames = make([]string, 0)
	for i := range res {
		taskNames = append(taskNames, res[i].Member.(string))
	}
	return
}
