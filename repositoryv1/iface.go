package repositoryv1

import (
	"github.com/lovemew67/leader-board/gen/go/proto"
)

type ScoreV1Repository interface {
	InsertScore(*proto.ScoreV1) (*proto.ScoreV1, error)
	ListTopKHighestScores(limit int) ([]*proto.ScoreV1, error)
	CleanScores() error
}

type ScoreV1CacheRepository interface {
	SetTopKScores([]*proto.ScoreV1) error
	GetTopKScores() ([]*proto.ScoreV1, error)
	CleanTopKScores() error
}
