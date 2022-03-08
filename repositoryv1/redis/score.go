package redis

import (
	goproto "google.golang.org/protobuf/proto"

	"github.com/lovemew67/leader-board/gen/go/proto"
	"github.com/lovemew67/leader-board/lb"
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

func (s *ScoreV1RedisCacheRepositorier) SetTopKScores(ctx cornerstone.Context, scores []*proto.ScoreV1) (err error) {
	funcName := "ScoreV1RedisCacheRepositorier.SetTopKScores"

	scoreList := &proto.ScoreListV1{
		Scores: scores,
	}
	b, err := goproto.Marshal(scoreList)
	if err != nil {
		cornerstone.Errorf(ctx, "[%s] failed to marshal json, err: %+v", funcName, err)
		return
	}

	err = s.pool.Set(lb.DefaultMaxLengthStr, b)
	if err != nil {
		cornerstone.Errorf(ctx, "[%s] s.pool.Set failed, err: %+v", funcName, err)
	} else {
		cornerstone.Debugf(ctx, "[%s] set top k scores in redis", funcName)
	}
	return
}

func (s *ScoreV1RedisCacheRepositorier) GetTopKScores(ctx cornerstone.Context) (scores []*proto.ScoreV1, err error) {
	funcName := "ScoreV1RedisCacheRepositorier.GetTopKScores"

	b, err := s.pool.GetBytes(lb.DefaultMaxLengthStr)
	if err != nil {
		if err.Error() == lb.ErrRedisKeyNotFoundStr {
			cornerstone.Debugf(ctx, "[%s] value not found with key: %s", funcName, lb.DefaultMaxLengthStr)
			err = lb.ErrRedisKeyNotFound
		} else {
			cornerstone.Errorf(ctx, "[%s] s.pool.GetBytes failed, err: %+v", funcName, err)
		}
		return
	}

	scoreList := &proto.ScoreListV1{}
	err = goproto.Unmarshal(b, scoreList)
	if err != nil {
		cornerstone.Errorf(ctx, "[%s] failed to unmarshal protobuf, err: %+v", funcName, err)
		return
	} else {
		cornerstone.Debugf(ctx, "[%s] got top k scores in redis", funcName)
	}

	scores = scoreList.Scores
	return
}

func (s *ScoreV1RedisCacheRepositorier) CleanTopKScores(ctx cornerstone.Context) (err error) {
	funcName := "ScoreV1RedisCacheRepositorier.CleanTopKScores"

	err = s.pool.Delete(lb.DefaultMaxLengthStr)
	if err != nil {
		cornerstone.Errorf(ctx, "[%s] s.pool.Delete failed, err: %+v", funcName, err)
	} else {
		cornerstone.Debugf(ctx, "[%s] cleaned top k scores in redis", funcName)
	}
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
