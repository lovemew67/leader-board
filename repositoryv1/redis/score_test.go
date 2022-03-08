package redis

import (
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/lovemew67/leader-board/gen/go/proto"
	"github.com/lovemew67/leader-board/lb"
	"github.com/lovemew67/public-misc/connection-poller/redis"
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
	ctx = cornerstone.NewContext()

	repo *ScoreV1RedisCacheRepositorier
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
		repo, errNew = NewScoreV1RedisCacheRepositorier(ctx, &redis.Config{
			Host:    fmt.Sprintf("localhost:%s", dockerResource.GetPort("6379/tcp")),
			RedisDB: 1,
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

func Test_All(t *testing.T) {
	// test: get before set
	scores, err := repo.GetTopKScores(ctx)
	assert.Error(t, err)
	assert.Equal(t, lb.ErrRedisKeyNotFound, err)
	assert.Equal(t, 0, len(scores))

	// test: set empty
	testScores := []*proto.ScoreV1{}
	err = repo.SetTopKScores(ctx, testScores)
	assert.NoError(t, err)

	// test: get after set
	scores, err = repo.GetTopKScores(ctx)
	assert.NoError(t, err)
	assert.Equal(t, 0, len(scores))

	// test: delete
	err = repo.CleanTopKScores(ctx)
	assert.NoError(t, err)

	// test: get after delete
	scores, err = repo.GetTopKScores(ctx)
	assert.Error(t, err)
	assert.Equal(t, lb.ErrRedisKeyNotFound, err)
	assert.Equal(t, 0, len(scores))

	// test: set non-empty
	testScores = []*proto.ScoreV1{
		{
			Id:    "id1",
			Score: 1.0,
		},
		{
			Id:    "id2",
			Score: 2.0,
		},
	}
	err = repo.SetTopKScores(ctx, testScores)
	assert.NoError(t, err)

	// test: get after set
	scores, err = repo.GetTopKScores(ctx)
	assert.NoError(t, err)
	assert.Equal(t, 2, len(scores))

	// test: delete
	err = repo.CleanTopKScores(ctx)
	assert.NoError(t, err)

	// test: get after delete
	scores, err = repo.GetTopKScores(ctx)
	assert.Error(t, err)
	assert.Equal(t, lb.ErrRedisKeyNotFound, err)
	assert.Equal(t, 0, len(scores))
}
