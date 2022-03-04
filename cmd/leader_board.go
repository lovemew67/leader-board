package cmd

import (
	"log"
	"main/controllerv1"
	"main/repositoryv1/sqlite"
	"main/servicev1"
	"os"
	"os/signal"
	"syscall"

	_ "github.com/jinzhu/gorm/dialects/sqlite"

	"github.com/lovemew67/public-misc/cornerstone"
	"github.com/spf13/cobra"
)

var (
	quit = make(chan os.Signal, 5)
)

func newLeaderBoardCmd() (result *cobra.Command, err error) {
	result = &cobra.Command{
		Use: "leaderboard",
		Run: func(cmd *cobra.Command, args []string) {
			funcName := "leaderBoardCmd.Run"
			log.Printf("[%s] in run function", funcName)

			ctx := cornerstone.NewContext()

			// init repository
			staffV1Repositorier, errRepository := sqlite.NewStaffV1SQLiteRepositorier(ctx)
			if errRepository != nil {
				cornerstone.Panicf(ctx, "[%s] failed to create staff v1 repositiory, err: %+v", funcName, errRepository)
			}

			// init service
			staffV1Service, errService := servicev1.NewStaffV1Servicer(staffV1Repositorier)
			if errService != nil {
				cornerstone.Panicf(ctx, "[%s] failed to create staff v1 service, err: %+v", funcName, errService)
			}

			// init http server
			ginServer := controllerv1.InitGinServer(staffV1Service)
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
