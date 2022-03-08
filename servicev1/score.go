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
	if req.Id == "" {
		err = fmt.Errorf("empty id")
		return
	}
	if req.Score == 0 {
		err = fmt.Errorf("zero score")
		return
	}
	result, err = s.r.InsertScore(cornerstone.NewContext(), req.ScoreV1)
	if err != nil {
		return
	}
	cornerstone.Debugf(ctx, "[%s] inserted score: %f, id: %s", funcName, req.Score, req.Id)
	err = s.cr.CleanTopKScores(cornerstone.NewContext())
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
	results, err = s.cr.GetTopKScores(cornerstone.NewContext())
	if err == nil {
		cornerstone.Debugf(ctx, "[%s] cache hit", funcName)
		if len(results) >= req.Limit {
			results = results[:req.Limit]
		}
	} else {
		if err == lb.ErrRedisKeyNotFound {
			cornerstone.Debugf(ctx, "[%s] cache miss", funcName)
			results, err = s.r.ListTopKHighestScores(cornerstone.NewContext(), lb.DefaultMaxLengthInt)
			if err != nil {
				return
			}
			cornerstone.Debugf(ctx, "[%s] got top k highest scores from databse", funcName)
			err = s.cr.SetTopKScores(cornerstone.NewContext(), results)
			if err != nil {
				return
			}
			cornerstone.Debugf(ctx, "[%s] set top k highest scores to cache", funcName)
			if len(results) >= req.Limit {
				results = results[:req.Limit]
			}
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
