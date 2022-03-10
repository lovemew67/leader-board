package sqlite

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/lovemew67/leader-board/gen/go/proto"
	"github.com/lovemew67/leader-board/repositoryv1"
	"github.com/lovemew67/public-misc/cornerstone"
	"github.com/spf13/viper"
)

const (
	scoreV1TableName = "score_v1"
)

var (
	_ repositoryv1.ScoreV1Repository = &ScoreV1SQLiteRepositorier{}
)

type ScoreV1SQLiteRepositorier struct{}

func (s *ScoreV1SQLiteRepositorier) InsertScore(ctx cornerstone.Context, score *proto.ScoreV1) (result *proto.ScoreV1, err error) {
	funcName := "ScoreV1SQLiteRepositorier.InsertScore"

	now := time.Now().UnixNano()
	score.Created = now
	score.Updated = now

	// FIXME: sqlite upsert
	db := sqlitedb
	db = db.Where("client_id = ?", score.ClientId)
	db.Delete(&proto.ScoreV1{})
	err = db.Error
	if err != nil {
		cornerstone.Errorf(ctx, "[%s] database opertaion err: %+v", funcName, err)
		return
	}
	db = db.Create(score)

	err = db.Error
	if err != nil {
		cornerstone.Errorf(ctx, "[%s] database opertaion err: %+v", funcName, err)
		return
	}
	cornerstone.Debugf(ctx, "[%s] inserted scores in sqlite", funcName)

	result, err = s.getScore(ctx, score.ClientId)
	return
}

func (s *ScoreV1SQLiteRepositorier) getScore(ctx cornerstone.Context, clientID string) (staff *proto.ScoreV1, err error) {
	funcName := "ScoreV1SQLiteRepositorier.getScore"

	staffList := make([]*proto.ScoreV1, 1)

	db := sqlitedb
	db = db.Where("client_id = ?", clientID)
	db = db.Limit(1)
	db = db.Find(&staffList)

	err = db.Error
	if len(staffList) != 0 {
		staff = staffList[0]
	} else {
		cornerstone.Errorf(ctx, "[%s] database opertaion err: %+v", funcName, err)
		err = fmt.Errorf("not found")
	}
	return
}

func (s *ScoreV1SQLiteRepositorier) ListTopKHighestScores(ctx cornerstone.Context, limit int) (scoreList []*proto.ScoreV1, err error) {
	funcName := "ScoreV1SQLiteRepositorier.ListTopKHighestScores"

	scoreList = make([]*proto.ScoreV1, limit)

	db := sqlitedb
	db = db.Order("score DESC")
	db = db.Limit(limit)
	db = db.Find(&scoreList)

	err = db.Error
	if err != nil {
		cornerstone.Errorf(ctx, "[%s] database opertaion err: %+v", funcName, err)
	} else {
		cornerstone.Debugf(ctx, "[%s] got top k scores: %d, in sqlite", funcName, limit)
	}
	return
}

func (s *ScoreV1SQLiteRepositorier) CleanScores(ctx cornerstone.Context) (err error) {
	funcName := "ScoreV1SQLiteRepositorier.CleanScores"

	db := sqlitedb
	db = db.Where("1 = 1")
	db = db.Delete(&proto.ScoreV1{})

	err = db.Error
	if err != nil {
		cornerstone.Errorf(ctx, "[%s] database opertaion err: %+v", funcName, err)
	} else {
		cornerstone.Debugf(ctx, "[%s] cleaned scores in sqlite", funcName)
	}
	return
}

func NewScoreV1SQLiteRepositorier(ctx cornerstone.Context) (result *ScoreV1SQLiteRepositorier, err error) {
	funcName := "NewScoreV1SQLiteRepositorier"

	viperDataFolder := viper.GetString("database.sqlite.folder")
	if viperDataFolder != "" {
		dataFolder = viperDataFolder
	}

	dbFilePath := fmt.Sprintf(formatSqliteDatabasePath, dataFolder)
	if err = createDirIfNotExist(ctx, dataFolder); err != nil {
		return
	}
	if err = createFileIfNotExist(ctx, dbFilePath); err != nil {
		return
	}

	db, err := gorm.Open(dialect, dbFilePath)
	if err != nil {
		return
	}
	sqlitedb = db

	score := &proto.ScoreV1{}
	hasTable := sqlitedb.HasTable(score)
	if hasTable {
		cornerstone.Infof(ctx, "[%s] continue to reuse the table: %s", funcName, scoreV1TableName)
		db.AutoMigrate(&proto.ScoreV1{})
		return
	}

	err = sqlitedb.CreateTable(score).Error
	if err != nil {
		return
	}

	result = &ScoreV1SQLiteRepositorier{}
	return
}
