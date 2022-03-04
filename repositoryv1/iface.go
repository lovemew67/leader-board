package repositoryv1

import (
	"main/gen/go/proto"
)

type StaffV1Repository interface {
	CreateStaff(*proto.StaffV1) (*proto.StaffV1, error)
	CountTotalStaff() (int, error)
	QueryAllStaffWithOffsetAndLimit(offset, limit int) ([]*proto.StaffV1, error)
}
