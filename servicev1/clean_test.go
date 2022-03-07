package servicev1

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/lovemew67/leader-board/repositoryv1mock"
	"github.com/lovemew67/public-misc/cornerstone"
	"github.com/stretchr/testify/assert"
)

func Test_Start(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	mockScoreV1Repositorier := repositoryv1mock.NewMockScoreV1Repository(controller)

	mockScoreV1CacheRepositorier := repositoryv1mock.NewMockScoreV1CacheRepository(controller)

	cleanUpBackgroundV1Servicer, err := NewCleanUpBackgroundV1Servicer(mockScoreV1Repositorier, mockScoreV1CacheRepositorier)
	assert.NoError(t, err)

	ctx := cornerstone.NewContext()
	err = cleanUpBackgroundV1Servicer.Start(ctx)
	assert.NoError(t, err)
}
