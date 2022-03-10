package sqlite

import (
	"log"
	"os"
	"testing"
	"time"

	_ "github.com/jinzhu/gorm/dialects/sqlite"

	"github.com/lovemew67/leader-board/gen/go/proto"
	"github.com/lovemew67/leader-board/lb"
	"github.com/lovemew67/public-misc/cornerstone"
	"github.com/stretchr/testify/assert"
)

var (
	ctx = cornerstone.NewContext()
)

func beforeTest() {}

func TestMain(m *testing.M) {
	log.SetFlags(log.LstdFlags)
	log.SetOutput(os.Stderr)
	beforeTest()
	retCode := m.Run()
	afterTest()
	os.Exit(retCode)
}

func afterTest() {
	_ = os.RemoveAll(dataFolder)
}

func Test_All(t *testing.T) {
	// test: init repo
	_, err := NewScoreV1SQLiteRepositorier(ctx)
	assert.NoError(t, err)

	// test: init repo again
	repo, err := NewScoreV1SQLiteRepositorier(ctx)
	assert.NoError(t, err)

	// test: list before insert
	result, err := repo.ListTopKHighestScores(ctx, lb.DefaultMaxLengthInt)
	assert.Equal(t, 0, len(result))
	assert.NoError(t, err)

	// test: delete before insert
	err = repo.CleanScores(ctx)
	assert.NoError(t, err)

	// test: insert more than default scores
	for i := 0; i < lb.DefaultMaxLengthInt/2; i++ {
		_, err = repo.InsertScore(ctx, &proto.ScoreV1{
			ClientId: time.Now().UTC().String(),
			Score:    float32(i),
		})
		assert.NoError(t, err)
	}

	// test: get only top 10
	result, err = repo.ListTopKHighestScores(ctx, lb.DefaultMaxLengthInt)
	assert.Equal(t, lb.DefaultMaxLengthInt/2, len(result))
	assert.NoError(t, err)

	// test: delete after insert
	err = repo.CleanScores(ctx)
	assert.NoError(t, err)

	// test: list after delete
	result, err = repo.ListTopKHighestScores(ctx, lb.DefaultMaxLengthInt)
	assert.Equal(t, 0, len(result))
	assert.NoError(t, err)

	// test: insert more than default scores
	for i := 0; i < 2*lb.DefaultMaxLengthInt; i++ {
		_, err = repo.InsertScore(ctx, &proto.ScoreV1{
			ClientId: time.Now().UTC().String(),
			Score:    float32(i),
		})
		assert.NoError(t, err)
	}

	// test: get only top 10
	result, err = repo.ListTopKHighestScores(ctx, lb.DefaultMaxLengthInt)
	assert.Equal(t, lb.DefaultMaxLengthInt, len(result))
	assert.NoError(t, err)
}
