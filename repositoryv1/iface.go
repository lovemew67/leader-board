package repositoryv1

import (
	"github.com/lovemew67/leader-board/gen/go/proto"
	"github.com/lovemew67/public-misc/cornerstone"
)

type ScoreV1Repository interface {
	InsertScore(cornerstone.Context, *proto.ScoreV1) (*proto.ScoreV1, error)
	ListTopKHighestScores(cornerstone.Context, int) ([]*proto.ScoreV1, error)
	CleanScores(cornerstone.Context) error
}

type ScoreV1CacheRepository interface {
	SetTopKScores(cornerstone.Context, []*proto.ScoreV1) error
	GetTopKScores(cornerstone.Context) ([]*proto.ScoreV1, error)
	CleanTopKScores(cornerstone.Context) error
}
