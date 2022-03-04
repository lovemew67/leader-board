package domainv1

import (
	"main/gen/go/proto"
)

// staff v1

type CreateStaffV1ServiceRequest struct {
	*proto.StaffV1
}

type ListStaffV1ServiceRequest struct {
	Offset int `form:"offset" url:"offset"`
	Limit  int `form:"limit" url:"limit"`
}
