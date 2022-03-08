package servicev1

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/lovemew67/leader-board/domainv1"
	"github.com/lovemew67/leader-board/gen/go/proto"
	"github.com/lovemew67/leader-board/lb"
	"github.com/lovemew67/leader-board/repositoryv1mock"
	"github.com/stretchr/testify/assert"
)

func Test_InsertScoreV1Service(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	mockScoreV1Repositorier := repositoryv1mock.NewMockScoreV1Repository(controller)
	mockScoreV1Repositorier.EXPECT().InsertScore(ctx, gomock.Any()).Return(&proto.ScoreV1{}, nil)

	mockScoreV1CacheRepositorier := repositoryv1mock.NewMockScoreV1CacheRepository(controller)
	mockScoreV1CacheRepositorier.EXPECT().CleanTopKScores(ctx).Return(nil)

	scoreV1Servicer, err := NewScoreV1Servicer(mockScoreV1Repositorier, mockScoreV1CacheRepositorier)
	assert.NoError(t, err)

	req := &domainv1.InsertScoreV1ServiceRequest{
		ScoreV1: &proto.ScoreV1{},
	}
	result, err := scoreV1Servicer.InsertScoreV1Service(ctx, req)
	assert.Error(t, err)
	assert.Nil(t, result)

	req = &domainv1.InsertScoreV1ServiceRequest{
		ScoreV1: &proto.ScoreV1{
			ClientId: "client_id",
		},
	}
	result, err = scoreV1Servicer.InsertScoreV1Service(ctx, req)
	assert.Error(t, err)
	assert.Nil(t, result)

	req = &domainv1.InsertScoreV1ServiceRequest{
		ScoreV1: &proto.ScoreV1{
			ClientId: "client_id",
			Score:    1,
		},
	}
	result, err = scoreV1Servicer.InsertScoreV1Service(ctx, req)
	assert.NoError(t, err)
	assert.NotNil(t, result)
}

func Test_ListTopKScoresV1Service_CacheHitWithoutLimit(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	mockScoreList := make([]*proto.ScoreV1, lb.DefaultMaxLengthInt)
	for i := 0; i < lb.DefaultMaxLengthInt; i++ {
		mockScoreList[i] = &proto.ScoreV1{}
	}

	mockScoreV1Repositorier := repositoryv1mock.NewMockScoreV1Repository(controller)

	mockScoreV1CacheRepositorier := repositoryv1mock.NewMockScoreV1CacheRepository(controller)
	mockScoreV1CacheRepositorier.EXPECT().GetTopKScores(ctx).Return(mockScoreList, nil)

	scoreV1Servicer, err := NewScoreV1Servicer(mockScoreV1Repositorier, mockScoreV1CacheRepositorier)
	assert.NoError(t, err)

	req := &domainv1.ListTopKScoresV1ServiceRequest{}
	result, err := scoreV1Servicer.ListTopKScoresV1Service(ctx, req)
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, lb.DefaultMaxLengthInt, len(result))
}

func Test_ListTopKScoresV1Service_CacheHitWithLimit(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	mockScoreList := make([]*proto.ScoreV1, lb.DefaultMaxLengthInt)
	for i := 0; i < lb.DefaultMaxLengthInt; i++ {
		mockScoreList[i] = &proto.ScoreV1{}
	}

	mockScoreV1Repositorier := repositoryv1mock.NewMockScoreV1Repository(controller)

	mockScoreV1CacheRepositorier := repositoryv1mock.NewMockScoreV1CacheRepository(controller)
	mockScoreV1CacheRepositorier.EXPECT().GetTopKScores(ctx).Times(2).Return(mockScoreList, nil)

	scoreV1Servicer, err := NewScoreV1Servicer(mockScoreV1Repositorier, mockScoreV1CacheRepositorier)
	assert.NoError(t, err)

	req := &domainv1.ListTopKScoresV1ServiceRequest{
		Limit: 5,
	}
	result, err := scoreV1Servicer.ListTopKScoresV1Service(ctx, req)
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, 5, len(result))

	req = &domainv1.ListTopKScoresV1ServiceRequest{
		Limit: 2 * lb.DefaultMaxLengthInt,
	}
	result, err = scoreV1Servicer.ListTopKScoresV1Service(ctx, req)
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, lb.DefaultMaxLengthInt, len(result))
}

func Test_ListTopKScoresV1Service_CacheMissWithLimitLess(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	mockScoreList := make([]*proto.ScoreV1, lb.DefaultMaxLengthInt)
	for i := 0; i < lb.DefaultMaxLengthInt; i++ {
		mockScoreList[i] = &proto.ScoreV1{}
	}

	mockScoreV1Repositorier := repositoryv1mock.NewMockScoreV1Repository(controller)
	mockScoreV1Repositorier.EXPECT().ListTopKHighestScores(ctx, lb.DefaultMaxLengthInt).Return(mockScoreList, nil)

	mockScoreV1CacheRepositorier := repositoryv1mock.NewMockScoreV1CacheRepository(controller)
	mockScoreV1CacheRepositorier.EXPECT().GetTopKScores(ctx).Return(nil, lb.ErrRedisKeyNotFound)
	mockScoreV1CacheRepositorier.EXPECT().SetTopKScores(ctx, mockScoreList).Return(nil)

	scoreV1Servicer, err := NewScoreV1Servicer(mockScoreV1Repositorier, mockScoreV1CacheRepositorier)
	assert.NoError(t, err)

	req := &domainv1.ListTopKScoresV1ServiceRequest{
		Limit: 5,
	}
	result, err := scoreV1Servicer.ListTopKScoresV1Service(ctx, req)
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, 5, len(result))
}

func Test_ListTopKScoresV1Service_CacheMissWithLimitMore(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	mockScoreList := make([]*proto.ScoreV1, lb.DefaultMaxLengthInt)
	for i := 0; i < lb.DefaultMaxLengthInt; i++ {
		mockScoreList[i] = &proto.ScoreV1{}
	}

	mockScoreV1Repositorier := repositoryv1mock.NewMockScoreV1Repository(controller)
	mockScoreV1Repositorier.EXPECT().ListTopKHighestScores(ctx, lb.DefaultMaxLengthInt).Return(mockScoreList, nil)

	mockScoreV1CacheRepositorier := repositoryv1mock.NewMockScoreV1CacheRepository(controller)
	mockScoreV1CacheRepositorier.EXPECT().GetTopKScores(ctx).Return(nil, lb.ErrRedisKeyNotFound)
	mockScoreV1CacheRepositorier.EXPECT().SetTopKScores(ctx, mockScoreList).Return(nil)

	scoreV1Servicer, err := NewScoreV1Servicer(mockScoreV1Repositorier, mockScoreV1CacheRepositorier)
	assert.NoError(t, err)

	req := &domainv1.ListTopKScoresV1ServiceRequest{
		Limit: 2 * lb.DefaultMaxLengthInt,
	}
	result, err := scoreV1Servicer.ListTopKScoresV1Service(ctx, req)
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, lb.DefaultMaxLengthInt, len(result))
}
