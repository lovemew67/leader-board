package servicev1

import (
	"fmt"

	"github.com/lovemew67/leader-board/domainv1"
	"github.com/lovemew67/leader-board/gen/go/proto"
	"github.com/lovemew67/leader-board/lb"
	"github.com/lovemew67/leader-board/repositoryv1"
	"github.com/lovemew67/public-misc/cornerstone"
)

var (
	_ ScoreV1Service = &ScoreV1Servicer{}
)

type ScoreV1Servicer struct {
	r  repositoryv1.ScoreV1Repository
	cr repositoryv1.ScoreV1CacheRepository
}

func (s *ScoreV1Servicer) InsertScoreV1Service(ctx cornerstone.Context, req *domainv1.InsertScoreV1ServiceRequest) (result *proto.ScoreV1, err error) {
	funcName := "ScoreV1Servicer.InsertScoreV1Service"
	if req.ClientId == "" {
		err = fmt.Errorf("empty client id")
		return
	}
	if req.Score == 0 {
		err = fmt.Errorf("zero score")
		return
	}
	result, err = s.r.InsertScore(ctx, req.ScoreV1)
	if err != nil {
		return
	}
	cornerstone.Debugf(ctx, "[%s] inserted score: %f, client id: %s", funcName, req.Score, req.ClientId)
	err = s.cr.CleanTopKScores(ctx)
	return
}

func (s *ScoreV1Servicer) ListTopKScoresV1Service(ctx cornerstone.Context, req *domainv1.ListTopKScoresV1ServiceRequest) (results []*proto.ScoreV1, err error) {
	funcName := "ScoreV1Servicer.ListTopKScoresV1Service"
	if req.Limit <= 0 {
		req.Limit = lb.DefaultMaxLengthInt
	}
	if req.Limit > lb.DefaultMaxLengthInt {
		req.Limit = lb.DefaultMaxLengthInt
	}
	results, err = s.cr.GetTopKScores(ctx)
	if err == nil {
		cornerstone.Debugf(ctx, "[%s] cache hit", funcName)
		if len(results) >= req.Limit {
			results = results[:req.Limit]
		}
	} else if err == lb.ErrRedisKeyNotFound {
		cornerstone.Debugf(ctx, "[%s] cache miss", funcName)
		results, err = s.r.ListTopKHighestScores(ctx, lb.DefaultMaxLengthInt)
		if err != nil {
			return
		}
		cornerstone.Debugf(ctx, "[%s] got top k highest scores from databse", funcName)
		// FIXME: json null after first get empty list
		if len(results) != 0 {
			err = s.cr.SetTopKScores(ctx, results)
			if err != nil {
				return
			}
			cornerstone.Debugf(ctx, "[%s] set top k highest scores to cache", funcName)
		}
		if len(results) >= req.Limit {
			results = results[:req.Limit]
		}
	}
	return
}

func NewScoreV1Servicer(_r repositoryv1.ScoreV1Repository, _cr repositoryv1.ScoreV1CacheRepository) (result *ScoreV1Servicer, err error) {
	result = &ScoreV1Servicer{
		r:  _r,
		cr: _cr,
	}
	return
}
