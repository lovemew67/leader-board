package controllerv1

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/lovemew67/leader-board/domainv1"
	"github.com/lovemew67/leader-board/servicev1"
	"github.com/lovemew67/public-misc/cornerstone"
	"github.com/spf13/viper"
)

type middlewareResponse struct {
	ErrorMessage string `json:"error_message,omitempty"`
}

func InitGinServer(_ss servicev1.ScoreV1Service) (gs *GinServer) {
	// create gin server.
	gin.SetMode(viper.GetString("rest.mode"))
	gs = &GinServer{
		ss: _ss,

		Engine: gin.New(),
	}
	gs.initRoutings()
	gs.NoRoute(noRouteMiddleware())
	return
}

func HTTPListenAndServe(ctx cornerstone.Context, gs *GinServer) (canceller func()) {
	funcName := "HTTPListenAndServe"
	restPort := viper.GetString("rest.port")
	httpServer := &http.Server{
		Addr:         ":" + restPort,
		Handler:      gs,
		ReadTimeout:  viper.GetDuration("rest.read_timeout"),
		WriteTimeout: viper.GetDuration("rest.write_timeout"),
	}
	go func() {
		cornerstone.Infof(ctx, "[%s] http server is running and listening port: %s", funcName, restPort)
		if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			cornerstone.Panicf(ctx, "[%s] http server failed to listen on port: %s, err: %+v", funcName, restPort, err)
		}
	}()

	routineCtx := ctx.CopyContext()
	canceller = func() {
		cornerstone.Infof(routineCtx, "[%s] shuting down http server", cornerstone.GetAppName())
		nativeCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if errShutdown := httpServer.Shutdown(nativeCtx); errShutdown != nil {
			cornerstone.Panicf(routineCtx, "[%s] failed to shut down http server, err: %+v", cornerstone.GetAppName(), errShutdown)
		}
		cornerstone.Infof(routineCtx, "[%s] http server exiting", cornerstone.GetAppName())
	}
	return
}

type GinServer struct {
	ss servicev1.ScoreV1Service

	*gin.Engine
}

func panicMiddleware() (result gin.HandlerFunc) {
	result = func(c *gin.Context) {
		defer func() {
			funcName := "panicMiddleware"
			if err := recover(); err != nil {
				log.Printf("[%s] recovered from panic, err: %+v", funcName, err)
				c.AbortWithStatusJSON(http.StatusInternalServerError, middlewareResponse{
					ErrorMessage: fmt.Sprintf("panic occured, err: %+v", err),
				})
			}
		}()
		c.Next()
	}
	return
}

func noRouteMiddleware() (result gin.HandlerFunc) {
	result = func(c *gin.Context) {
		c.AbortWithStatusJSON(http.StatusInternalServerError, middlewareResponse{
			ErrorMessage: fmt.Sprintf("page not found, method: %s uri: %s", c.Request.Method, c.Request.Host+c.Request.URL.Path),
		})
	}
	return
}

// for swagger
func getCORSConfig() (config cors.Config) {
	config = cors.Config{
		MaxAge:           1728000 * time.Second,
		AllowAllOrigins:  true,
		AllowCredentials: true,
		AllowMethods: []string{
			http.MethodGet,
			http.MethodPost,
			http.MethodOptions,
			http.MethodDelete,
			http.MethodPatch,
			http.MethodPut,
		},
		AllowHeaders: []string{
			"Authorization",
			"Cache-Control",
			"Content-Type",
			"DNT",
			"If-Modified-Since",
			"Keep-Alive",
			"User-Agent",
			"X-CustomHeader",
			"X-Requested-With",
		},
		ExposeHeaders: []string{
			"content-disposition",
		},
	}
	return
}

func (gs *GinServer) initRoutings() {
	// add data retention group
	rootGroup := gs.Group("")
	rootGroup.Use(panicMiddleware())
	rootGroup.Use(cors.New(getCORSConfig()))

	// general service for debugging
	{
		rootGroup.GET("/config", gs.config)
		rootGroup.GET("/version", gs.version)
	}

	// add staff v1 handlers
	staffGroup := rootGroup.Group("/v1/scores")
	{
		staffGroup.GET("", gs.listTopKScoresV1Handler)
		staffGroup.POST("", gs.insertScoreV1Handler)
		staffGroup.PATCH("", gs.panicTester)

		staffGroup.OPTIONS("")
	}
}

func (gs *GinServer) version(c *gin.Context) {
	c.JSON(http.StatusOK, cornerstone.GetVersion())
}

func (gs *GinServer) config(c *gin.Context) {
	c.JSON(http.StatusOK, viper.AllSettings())
}

// score v1 handlers

func (gs *GinServer) insertScoreV1Handler(c *gin.Context) {
	input := &domainv1.InsertScoreV1ServiceRequest{}
	if errBind := c.ShouldBindJSON(input); errBind != nil {
		cornerstone.FromCodeErrorWithStatus(c, cornerstone.FromNativeError(errBind))
		return
	}
	result, err := gs.ss.InsertScoreV1Service(input)
	if err != nil {
		cornerstone.FromCodeErrorWithStatus(c, cornerstone.FromNativeError(err))
		return
	}
	cornerstone.DoneWithStatus(c, result)
}

func (gs *GinServer) listTopKScoresV1Handler(c *gin.Context) {
	input := &domainv1.ListTopKScoresV1ServiceRequest{}
	if errBind := c.BindQuery(&input); errBind != nil {
		cornerstone.FromCodeErrorWithStatus(c, cornerstone.FromNativeError(errBind))
		return
	}
	results, err := gs.ss.ListTopKScoresV1Service(input)
	if err != nil {
		cornerstone.FromCodeErrorWithStatus(c, cornerstone.FromNativeError(err))
		return
	}
	cornerstone.DoneWithStatus(c, gin.H{
		"scores": results,
	})
}

func (gs *GinServer) panicTester(c *gin.Context) {
	panic("panic on purpose")
}
