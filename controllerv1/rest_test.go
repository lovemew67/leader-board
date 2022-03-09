package controllerv1

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/lovemew67/leader-board/domainv1"
	"github.com/lovemew67/leader-board/gen/go/proto"
	"github.com/lovemew67/leader-board/servicev1mock"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func Test_GinServer(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	mockScoreV1Servicer := servicev1mock.NewMockScoreV1Service(controller)

	gs := InitGinServer(mockScoreV1Servicer)
	assert.NotNil(t, gs)

	viper.Set("rest.port", "7999")
	canceller := HTTPListenAndServe(ctx, gs)
	assert.NotNil(t, canceller)

	canceller()
}

func Test_noRouteMiddleware(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	mockScoreV1Servicer := servicev1mock.NewMockScoreV1Service(controller)

	gs := InitGinServer(mockScoreV1Servicer)
	assert.NotNil(t, gs)

	req, _ := http.NewRequest(http.MethodGet, "/v1/no-route", nil)
	resp := httptest.NewRecorder()

	gs.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusNotFound, resp.Code)
}

func Test_panicMiddleware(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	mockScoreV1Servicer := servicev1mock.NewMockScoreV1Service(controller)

	gs := InitGinServer(mockScoreV1Servicer)
	assert.NotNil(t, gs)

	req, _ := http.NewRequest(http.MethodPatch, "/v1/scores", nil)
	resp := httptest.NewRecorder()

	gs.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusInternalServerError, resp.Code)

	result := &middlewareResponse{}
	err := json.Unmarshal(resp.Body.Bytes(), result)
	assert.NoError(t, err)
	assert.Equal(t, "panic occured, err: panic on purpose", result.ErrorMessage)
}

func Test_version(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	mockScoreV1Servicer := servicev1mock.NewMockScoreV1Service(controller)

	gs := InitGinServer(mockScoreV1Servicer)
	assert.NotNil(t, gs)

	req, _ := http.NewRequest(http.MethodGet, "/version", nil)
	resp := httptest.NewRecorder()

	gs.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusOK, resp.Code)
}

func Test_config(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	mockScoreV1Servicer := servicev1mock.NewMockScoreV1Service(controller)

	gs := InitGinServer(mockScoreV1Servicer)
	assert.NotNil(t, gs)

	req, _ := http.NewRequest(http.MethodGet, "/config", nil)
	resp := httptest.NewRecorder()

	gs.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusOK, resp.Code)
}

func Test_insertScoreV1Handler_emptyBody(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	mockScoreV1Servicer := servicev1mock.NewMockScoreV1Service(controller)

	gs := InitGinServer(mockScoreV1Servicer)
	assert.NotNil(t, gs)

	req, _ := http.NewRequest(http.MethodPost, "/v1/scores", nil)
	resp := httptest.NewRecorder()

	gs.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusInternalServerError, resp.Code)
}

func Test_insertScoreV1Handler_serviceError(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	mockScoreV1Servicer := servicev1mock.NewMockScoreV1Service(controller)
	mockScoreV1Servicer.EXPECT().InsertScoreV1Service(gomock.Any(), gomock.Any()).Return(nil, fmt.Errorf("mock error"))

	gs := InitGinServer(mockScoreV1Servicer)
	assert.NotNil(t, gs)

	p := &domainv1.InsertScoreV1ServiceRequest{}
	d, err := json.Marshal(p)
	assert.NoError(t, err)
	req, _ := http.NewRequest(http.MethodPost, "/v1/scores", bytes.NewBuffer(d))
	resp := httptest.NewRecorder()

	gs.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusInternalServerError, resp.Code)
}

func Test_insertScoreV1Handler_correct(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	mockScoreV1Servicer := servicev1mock.NewMockScoreV1Service(controller)
	mockScoreV1Servicer.EXPECT().InsertScoreV1Service(gomock.Any(), gomock.Any()).Return(&proto.ScoreV1{}, nil)

	gs := InitGinServer(mockScoreV1Servicer)
	assert.NotNil(t, gs)

	p := &domainv1.InsertScoreV1ServiceRequest{}
	d, err := json.Marshal(p)
	assert.NoError(t, err)
	req, _ := http.NewRequest(http.MethodPost, "/v1/scores", bytes.NewBuffer(d))
	resp := httptest.NewRecorder()

	gs.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusOK, resp.Code)
}

func Test_listTopKScoresV1Handler_serviceError(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	mockScoreV1Servicer := servicev1mock.NewMockScoreV1Service(controller)
	mockScoreV1Servicer.EXPECT().ListTopKScoresV1Service(gomock.Any(), gomock.Any()).Return(nil, fmt.Errorf("mock error"))

	gs := InitGinServer(mockScoreV1Servicer)
	assert.NotNil(t, gs)

	req, _ := http.NewRequest(http.MethodGet, "/v1/scores", nil)
	resp := httptest.NewRecorder()

	gs.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusInternalServerError, resp.Code)
}

func Test_listTopKScoresV1Handler_correct(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	mockScoreV1Servicer := servicev1mock.NewMockScoreV1Service(controller)
	mockScoreV1Servicer.EXPECT().ListTopKScoresV1Service(gomock.Any(), gomock.Any()).Return([]*proto.ScoreV1{}, nil)

	gs := InitGinServer(mockScoreV1Servicer)
	assert.NotNil(t, gs)

	req, _ := http.NewRequest(http.MethodGet, "/v1/scores", nil)
	resp := httptest.NewRecorder()

	gs.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusOK, resp.Code)
}
