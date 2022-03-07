package redis

import (
	"github.com/lovemew67/leader-board/gen/go/proto"
	"github.com/lovemew67/leader-board/repositoryv1"
	"github.com/lovemew67/public-misc/connection-poller/redis"
	"github.com/lovemew67/public-misc/cornerstone"
)

var (
	_ repositoryv1.ScoreV1CacheRepository = &ScoreV1RedisCacheRepositorier{}
)

type ScoreV1RedisCacheRepositorier struct {
	pool *redis.Pool
}

func (s *ScoreV1RedisCacheRepositorier) SetTopKScores(staff []*proto.ScoreV1) (err error) {
	return
}

func (s *ScoreV1RedisCacheRepositorier) GetTopKScores() (staff []*proto.ScoreV1, err error) {
	return
}

func (s *ScoreV1RedisCacheRepositorier) CleanTopKScores() (err error) {
	return
}

func NewScoreV1RedisCacheRepositorier(ctx cornerstone.Context, cfg *redis.Config) (result *ScoreV1RedisCacheRepositorier, err error) {
	funcName := "NewScoreV1RedisCacheRepositorier"

	_pool, err := redis.NewPool(cfg)
	if err != nil {
		cornerstone.Errorf(ctx, "[%s] failed to init redis pool, err: %+v", funcName, err)
		return
	}

	result = &ScoreV1RedisCacheRepositorier{
		pool: _pool,
	}
	return
}
