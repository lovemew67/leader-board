package servicev1

import (
	"github.com/lovemew67/leader-board/domainv1"
	"github.com/lovemew67/leader-board/gen/go/proto"
	"github.com/lovemew67/public-misc/cornerstone"
)

type StaffV1Service interface {
	CreateStaffV1Service(*domainv1.CreateStaffV1ServiceRequest) (*proto.StaffV1, error)
	ListStaffV1Service(*domainv1.ListStaffV1ServiceRequest) ([]*proto.StaffV1, int, error)
}

type CleanUpBackgroundV1Service interface {
	Start(cornerstone.Context) error
}
