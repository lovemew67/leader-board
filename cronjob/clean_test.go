package cronjob

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/lovemew67/leader-board/repositoryv1mock"
	"github.com/lovemew67/public-misc/cornerstone"
	"github.com/stretchr/testify/assert"
)

var (
	ctx = cornerstone.NewContext()
)

func Test_StartCleanUpCron(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	mockScoreV1Repositorier := repositoryv1mock.NewMockScoreV1Repository(controller)
	mockScoreV1Repositorier.EXPECT().CleanScores(ctx).AnyTimes().Return(nil)

	mockScoreV1CacheRepositorier := repositoryv1mock.NewMockScoreV1CacheRepository(controller)
	mockScoreV1CacheRepositorier.EXPECT().CleanTopKScores(ctx).AnyTimes().Return(nil)

	err := StartCleanUpCron(ctx, mockScoreV1Repositorier, mockScoreV1CacheRepositorier)
	assert.NoError(t, err)
}

func Test_cleanUpCronJobImpl(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	mockScoreV1Repositorier := repositoryv1mock.NewMockScoreV1Repository(controller)
	mockScoreV1Repositorier.EXPECT().CleanScores(ctx).Return(nil)

	mockScoreV1CacheRepositorier := repositoryv1mock.NewMockScoreV1CacheRepository(controller)
	mockScoreV1CacheRepositorier.EXPECT().CleanTopKScores(ctx).Return(nil)

	cleanUpCronJobImpl(ctx, mockScoreV1Repositorier, mockScoreV1CacheRepositorier)
}
