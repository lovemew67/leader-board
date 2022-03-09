package cmd

import (
	"fmt"
	"log"
	"os"
	"testing"
	"time"

	conredis "github.com/lovemew67/public-misc/connection-poller/redis"
	"github.com/spf13/viper"

	"github.com/lovemew67/leader-board/repositoryv1/redis"
	"github.com/lovemew67/public-misc/cornerstone"
	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
	"github.com/stretchr/testify/assert"
)

var (
	dockerPool     *dockertest.Pool
	dockerResource *dockertest.Resource
)

var (
	ctx       = cornerstone.NewContext()
	redisPort = ""
)

func beforeTest() {
	// init error
	var err error

	// init docker
	dockerPool, err = dockertest.NewPool("")
	if err != nil {
		panic("docker test init fail, error:" + err.Error())
	}
	runOpts := &dockertest.RunOptions{
		Name:       "",
		Repository: "redis",
		Tag:        "5.0",
		Env:        []string{},
	}
	hcOpts := func(config *docker.HostConfig) {
		config.AutoRemove = true
		config.RestartPolicy = docker.RestartPolicy{
			Name: "no",
		}
	}
	dockerResource, err = dockerPool.RunWithOptions(runOpts, hcOpts)
	if err != nil {
		panic("redis docker init fail, error:" + err.Error())
	}
	err = dockerResource.Expire(600)
	if err != nil {
		panic("failed to set expire, err: " + err.Error())
	}
	err = dockerPool.Retry(func() error {
		// test connection or setup
		var errNew error
		redisPort = dockerResource.GetPort("6379/tcp")
		_, errNew = redis.NewScoreV1RedisCacheRepositorier(ctx, &conredis.Config{
			Host:    fmt.Sprintf("localhost:%s", redisPort),
			RedisDB: 0,
		})
		return errNew
	})
	if err != nil {
		panic("connect docker fail, error:" + err.Error())
	}
}

func TestMain(m *testing.M) {
	log.SetFlags(log.LstdFlags)
	log.SetOutput(os.Stderr)
	beforeTest()
	retCode := m.Run()
	afterTest()
	os.Exit(retCode)
}

func afterTest() {
	_ = dockerPool.Purge(dockerResource)
}

func Test_newRootCmd(t *testing.T) {
	result, err := newRootCmd()
	assert.NoError(t, err)

	result.SetArgs([]string{"--config", "../test.toml"})
	err = result.Execute()
	assert.NoError(t, err)
}

func Test_newLeaderBoardCmd(t *testing.T) {
	result, err := newLeaderBoardCmd()
	assert.NoError(t, err)

	go func() {
		time.Sleep(5 * time.Second)
		quit <- os.Interrupt
	}()

	viper.Set("database.redis.host", fmt.Sprintf("localhost:%s", redisPort))
	err = result.Execute()
	assert.NoError(t, err)
}
