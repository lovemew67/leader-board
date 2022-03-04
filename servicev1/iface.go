package servicev1

import (
	"main/domainv1"
	"main/gen/go/proto"
)

type StaffV1Service interface {
	CreateStaffV1Service(*domainv1.CreateStaffV1ServiceRequest) (*proto.StaffV1, error)
	ListStaffV1Service(*domainv1.ListStaffV1ServiceRequest) ([]*proto.StaffV1, int, error)
}
