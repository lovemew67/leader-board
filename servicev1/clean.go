package servicev1

import (
	"github.com/lovemew67/leader-board/cronjob"
	"github.com/lovemew67/leader-board/repositoryv1"
	"github.com/lovemew67/public-misc/cornerstone"
)

var (
	_ CleanUpBackgroundV1Service = &CleanUpBackgroundCronjobdV1Servicer{}
)

type CleanUpBackgroundCronjobdV1Servicer struct {
	r  repositoryv1.ScoreV1Repository
	cr repositoryv1.ScoreV1CacheRepository
}

func (s *CleanUpBackgroundCronjobdV1Servicer) Start(ctx cornerstone.Context) (err error) {
	err = cronjob.StartCleanUpCron(ctx, s.r, s.cr)
	return
}

func NewCleanUpBackgroundV1Servicer(_r repositoryv1.ScoreV1Repository, _cr repositoryv1.ScoreV1CacheRepository) (result *CleanUpBackgroundCronjobdV1Servicer, err error) {
	result = &CleanUpBackgroundCronjobdV1Servicer{
		r:  _r,
		cr: _cr,
	}
	return
}
