package servicev1

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/lovemew67/leader-board/domainv1"
	"github.com/lovemew67/leader-board/gen/go/proto"
	"github.com/lovemew67/leader-board/repositoryv1mock"
	"github.com/stretchr/testify/assert"
)

func Test_InsertScoreV1Service(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	mockScoreV1Repositorier := repositoryv1mock.NewMockScoreV1Repository(controller)
	mockScoreV1Repositorier.EXPECT().InsertScore(gomock.Any()).Return(&proto.ScoreV1{}, nil)

	mockScoreV1CacheRepositorier := repositoryv1mock.NewMockScoreV1CacheRepository(controller)

	scoreV1Servicer, err := NewScoreV1Servicer(mockScoreV1Repositorier, mockScoreV1CacheRepositorier)
	assert.NoError(t, err)

	req := &domainv1.InsertScoreV1ServiceRequest{
		ScoreV1: &proto.ScoreV1{},
	}
	result, err := scoreV1Servicer.InsertScoreV1Service(req)
	assert.Error(t, err)
	assert.Nil(t, result)

	req = &domainv1.InsertScoreV1ServiceRequest{
		ScoreV1: &proto.ScoreV1{
			Id: "id",
		},
	}
	result, err = scoreV1Servicer.InsertScoreV1Service(req)
	assert.Error(t, err)
	assert.Nil(t, result)

	req = &domainv1.InsertScoreV1ServiceRequest{
		ScoreV1: &proto.ScoreV1{
			Id:    "id",
			Score: 1,
		},
	}
	result, err = scoreV1Servicer.InsertScoreV1Service(req)
	assert.NoError(t, err)
	assert.NotNil(t, result)
}

func Test_ListTopKScoresV1Service(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	mockScoreV1Repositorier := repositoryv1mock.NewMockScoreV1Repository(controller)
	mockScoreV1Repositorier.EXPECT().ListTopKHighestScores(gomock.Any()).Return([]*proto.ScoreV1{}, nil)

	mockScoreV1CacheRepositorier := repositoryv1mock.NewMockScoreV1CacheRepository(controller)

	scoreV1Servicer, err := NewScoreV1Servicer(mockScoreV1Repositorier, mockScoreV1CacheRepositorier)
	assert.NoError(t, err)

	req := &domainv1.ListTopKScoresV1ServiceRequest{}
	result, err := scoreV1Servicer.ListTopKScoresV1Service(req)
	assert.NoError(t, err)
	assert.NotNil(t, result)
}
