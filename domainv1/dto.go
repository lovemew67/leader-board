package domainv1

import (
	"fmt"

	"github.com/lovemew67/leader-board/gen/go/proto"
)

type InsertScoreV1ServiceRequest struct {
	*proto.ScoreV1
}

func (r *InsertScoreV1ServiceRequest) String() string {
	if r != nil {
		return fmt.Sprintf("%+v", *r)
	}
	return ""
}

type ListTopKScoresV1ServiceRequest struct {
	Limit int `form:"limit" url:"limit"`
}

func (r *ListTopKScoresV1ServiceRequest) String() string {
	if r != nil {
		return fmt.Sprintf("%+v", *r)
	}
	return ""
}
