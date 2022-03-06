package repositoryv1

import (
	"github.com/lovemew67/leader-board/gen/go/proto"
)

type StaffV1Repository interface {
	CreateStaff(*proto.StaffV1) (*proto.StaffV1, error)
	CountTotalStaff() (int, error)
	QueryAllStaffWithOffsetAndLimit(offset, limit int) ([]*proto.StaffV1, error)
}

type StaffV1CacheRepository interface {
	SetTopNScores([]*proto.StaffV1) error
	GetTopNScores() ([]*proto.StaffV1, error)
	CleanTopNScores() error
}
