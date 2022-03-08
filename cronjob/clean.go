package cronjob

import (
	"time"

	"github.com/lovemew67/leader-board/repositoryv1"
	"github.com/lovemew67/public-misc/cornerstone"
	"github.com/robfig/cron/v3"
)

const (
	// At every 10th minute.
	cronSpec = "*/10 * * * *"
)

var (
	cleanUpCron           = cron.New(cron.WithLocation(time.UTC))
	cleanUpCronJobEntryID = cron.EntryID(0)
)

func StartCleanUpCron(ctx cornerstone.Context, r repositoryv1.ScoreV1Repository, cr repositoryv1.ScoreV1CacheRepository) (err error) {
	funcName := "StartCleanUpBackgroundCron"

	if cleanUpCronJobEntryID != 0 {
		cornerstone.Infof(ctx, "[%s] going to remove clean up cron job entry id: %d", funcName, cleanUpCronJobEntryID)
		cleanUpCron.Remove(cleanUpCronJobEntryID)
	}

	routineCtx := ctx.CopyContext()
	cleanUpCronJobEntryID, err = cleanUpCron.AddFunc(cronSpec, func() {
		cleanUpCronJobImpl(routineCtx, r, cr)
	})
	if err != nil {
		return
	}
	cornerstone.Infof(ctx, "[%s] going to start sync cron job entry id: %d", funcName, cleanUpCronJobEntryID)

	cleanUpCron.Start()
	return
}

func recoverCleanUpCronJobImpl(ctx cornerstone.Context) {
	funcName := "recoverCleanUpCronJobImpl"
	cornerstone.Infof(ctx, "[%s] exit proxy sync cronjob impl", funcName)
	if err := recover(); err != nil {
		cornerstone.Errorf(ctx, "[%s] recovered at: %s, panic err: %+v", funcName, time.Now().Format(time.RFC3339Nano), err)
	}
}

func cleanUpCronJobImpl(ctx cornerstone.Context, r repositoryv1.ScoreV1Repository, cr repositoryv1.ScoreV1CacheRepository) {
	funcName := "cleanUpCronJobImpl"
	defer recoverCleanUpCronJobImpl(ctx)
	cornerstone.Infof(ctx, "[%s] triggered", funcName)

	err := r.CleanScores(ctx)
	if err != nil {
		cornerstone.Errorf(ctx, "[%s] failed to clean scores in reposiroty", funcName)
		return
	}

	err = cr.CleanTopKScores(ctx)
	if err != nil {
		cornerstone.Errorf(ctx, "[%s] failed to clean scores in cache repository", funcName)
	}
}
