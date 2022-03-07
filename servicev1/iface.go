package servicev1

import (
	"github.com/lovemew67/leader-board/domainv1"
	"github.com/lovemew67/leader-board/gen/go/proto"
	"github.com/lovemew67/public-misc/cornerstone"
)

type ScoreV1Service interface {
	InsertScoreV1Service(*domainv1.InsertScoreV1ServiceRequest) (*proto.ScoreV1, error)
	ListTopKScoresV1Service(*domainv1.ListTopKScoresV1ServiceRequest) ([]*proto.ScoreV1, error)
}

type CleanUpBackgroundV1Service interface {
	Start(cornerstone.Context) error
}
