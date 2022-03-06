package servicev1

import (
	"github.com/lovemew67/leader-board/domainv1"
	"github.com/lovemew67/leader-board/gen/go/proto"
	"github.com/lovemew67/leader-board/repositoryv1"
)

var (
	_ StaffV1Service = &StaffV1Servicer{}
)

type StaffV1Servicer struct {
	r repositoryv1.StaffV1Repository
}

func (s *StaffV1Servicer) CreateStaffV1Service(req *domainv1.CreateStaffV1ServiceRequest) (result *proto.StaffV1, err error) {
	result, err = s.r.CreateStaff(req.StaffV1)
	return
}

func (s *StaffV1Servicer) ListStaffV1Service(req *domainv1.ListStaffV1ServiceRequest) (results []*proto.StaffV1, total int, err error) {
	total, err = s.r.CountTotalStaff()
	if err != nil {
		return
	}
	results, err = s.r.QueryAllStaffWithOffsetAndLimit(req.Offset, req.Limit)
	return
}

func NewStaffV1Servicer(_r repositoryv1.StaffV1Repository) (result *StaffV1Servicer, err error) {
	result = &StaffV1Servicer{
		r: _r,
	}
	return
}
