package redis

import (
	"github.com/lovemew67/leader-board/gen/go/proto"
	"github.com/lovemew67/leader-board/repositoryv1"

	"github.com/lovemew67/public-misc/connection-poller/redis"
	"github.com/lovemew67/public-misc/cornerstone"
)

var (
	_ repositoryv1.StaffV1CacheRepository = &StaffV1RedisCacheRepositorier{}
)

type StaffV1RedisCacheRepositorier struct {
	pool *redis.Pool
}

func (s *StaffV1RedisCacheRepositorier) SetTopNScores(staff []*proto.StaffV1) (err error) {
	return
}

func (s *StaffV1RedisCacheRepositorier) GetTopNScores() (staff []*proto.StaffV1, err error) {
	return
}

func (s *StaffV1RedisCacheRepositorier) CleanTopNScores() (err error) {
	return
}

func NewStaffV1RedisCacheRepositorier(ctx cornerstone.Context, cfg *redis.Config) (result *StaffV1RedisCacheRepositorier, err error) {
	funcName := "NewStaffV1RedisCacheRepositorier"

	_pool, err := redis.NewPool(cfg)
	if err != nil {
		cornerstone.Errorf(ctx, "[%s] failed to init redis pool, err: %+v", funcName, err)
		return
	}

	result = &StaffV1RedisCacheRepositorier{
		pool: _pool,
	}
	return
}
