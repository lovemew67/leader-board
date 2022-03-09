package cmd

import (
	"os"
	"os/signal"
	"syscall"

	_ "github.com/jinzhu/gorm/dialects/sqlite"
	conredis "github.com/lovemew67/public-misc/connection-poller/redis"

	"github.com/lovemew67/leader-board/controllerv1"
	"github.com/lovemew67/leader-board/repositoryv1/redis"
	"github.com/lovemew67/leader-board/repositoryv1/sqlite"
	"github.com/lovemew67/leader-board/servicev1"
	"github.com/lovemew67/public-misc/cornerstone"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	quit = make(chan os.Signal, 5)
)

func newLeaderBoardCmd() (result *cobra.Command, err error) {
	result = &cobra.Command{
		Use: "leaderboard",
		Run: func(cmd *cobra.Command, args []string) {
			funcName := "leaderBoardCmd.Run"

			// init context
			ctx := cornerstone.NewContext()

			// init repository
			ScoreV1Repositorier, errRepository := sqlite.NewScoreV1SQLiteRepositorier(ctx)
			if errRepository != nil {
				cornerstone.Panicf(ctx, "[%s] failed to create staff v1 repositiory, err: %+v", funcName, errRepository)
			}
			ScoreV1CacheRepositorier, errRepository := redis.NewScoreV1RedisCacheRepositorier(ctx, &conredis.Config{
				Host: viper.GetString("database.redis.host"),
			})
			if errRepository != nil {
				cornerstone.Panicf(ctx, "[%s] failed to create staff v1 cache repositiory, err: %+v", funcName, errRepository)
			}

			// init service
			ScoreV1Service, errService := servicev1.NewScoreV1Servicer(ScoreV1Repositorier, ScoreV1CacheRepositorier)
			if errService != nil {
				cornerstone.Panicf(ctx, "[%s] failed to create staff v1 service, err: %+v", funcName, errService)
			}
			cleanUpBackgroundV1Service, errService := servicev1.NewCleanUpBackgroundV1Servicer(ScoreV1Repositorier, ScoreV1CacheRepositorier)
			if errService != nil {
				cornerstone.Panicf(ctx, "[%s] failed to create clean up background v1 service, err: %+v", funcName, errService)
			}
			errService = cleanUpBackgroundV1Service.Start(ctx)
			if errService != nil {
				cornerstone.Panicf(ctx, "[%s] failed to start clean up background v1 service, err: %+v", funcName, errService)
			}

			// init http server
			ginServer := controllerv1.InitGinServer(ScoreV1Service)
			ginCanceller := controllerv1.HTTPListenAndServe(ctx, ginServer)
			defer ginCanceller()

			// add graceful shutdown
			signal.Notify(quit, os.Interrupt, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

			// blocking
			sig := <-quit
			cornerstone.Infof(ctx, "[%s] receive exit signal: %+v", funcName, sig)
		},
	}
	return
}
