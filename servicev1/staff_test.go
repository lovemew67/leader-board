package servicev1

import (
	"testing"

	"github.com/lovemew67/leader-board/domainv1"
	"github.com/lovemew67/leader-board/gen/go/proto"
	"github.com/lovemew67/leader-board/repositoryv1mock"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func Test_CreateStaffV1Service(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	mockStaffV1Repositorier := repositoryv1mock.NewMockStaffV1Repository(controller)
	mockStaffV1Repositorier.EXPECT().CreateStaff(gomock.Any()).Return(&proto.StaffV1{}, nil)

	mockStaffV1CacheRepositorier := repositoryv1mock.NewMockStaffV1CacheRepository(controller)

	mockStaffV1Servicer, err := NewStaffV1Servicer(mockStaffV1Repositorier, mockStaffV1CacheRepositorier)
	assert.NoError(t, err)

	req := &domainv1.CreateStaffV1ServiceRequest{}
	result, err := mockStaffV1Servicer.CreateStaffV1Service(req)
	assert.NoError(t, err)
	assert.NotNil(t, result)
}
