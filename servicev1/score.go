package servicev1

import (
	"fmt"

	"github.com/lovemew67/leader-board/domainv1"
	"github.com/lovemew67/leader-board/gen/go/proto"
	"github.com/lovemew67/leader-board/repositoryv1"
)

var (
	_ ScoreV1Service = &ScoreV1Servicer{}
)

type ScoreV1Servicer struct {
	r  repositoryv1.ScoreV1Repository
	cr repositoryv1.ScoreV1CacheRepository
}

func (s *ScoreV1Servicer) InsertScoreV1Service(req *domainv1.InsertScoreV1ServiceRequest) (result *proto.ScoreV1, err error) {
	if req.Id == "" {
		err = fmt.Errorf("empty id")
		return
	}
	if req.Score == 0 {
		err = fmt.Errorf("zero score")
		return
	}
	result, err = s.r.InsertScore(req.ScoreV1)
	if err != nil {
		return
	}
	return
}

func (s *ScoreV1Servicer) ListTopKScoresV1Service(req *domainv1.ListTopKScoresV1ServiceRequest) (results []*proto.ScoreV1, err error) {
	results, err = s.r.ListTopKHighestScores(req.Limit)
	return
}

func NewScoreV1Servicer(_r repositoryv1.ScoreV1Repository, _cr repositoryv1.ScoreV1CacheRepository) (result *ScoreV1Servicer, err error) {
	result = &ScoreV1Servicer{
		r:  _r,
		cr: _cr,
	}
	return
}
