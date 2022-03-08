package sqlite

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/lovemew67/leader-board/gen/go/proto"
	"github.com/lovemew67/leader-board/repositoryv1"
	"github.com/lovemew67/public-misc/cornerstone"
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

	db := sqlitedb
	db = db.Create(score)
	err = db.Error
	if err != nil {
		cornerstone.Errorf(ctx, "[%s] database opertaion err: %+v", funcName, err)
		return
	}

	result, err = s.getScore(ctx, score.Id)
	return
}

func (s *ScoreV1SQLiteRepositorier) getScore(ctx cornerstone.Context, id string) (staff *proto.ScoreV1, err error) {
	funcName := "ScoreV1SQLiteRepositorier.getScore"

	staffList := make([]*proto.ScoreV1, 1)

	db := sqlitedb
	db = db.Where("id = ?", id)
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
	}
	return
}

func NewScoreV1SQLiteRepositorier(ctx cornerstone.Context) (result *ScoreV1SQLiteRepositorier, err error) {
	funcName := "NewScoreV1SQLiteRepositorier"

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
	if hasTable := sqlitedb.HasTable(score); hasTable {
		cornerstone.Infof(ctx, "[%s] continue to reuse the table: %s", funcName, scoreV1TableName)
		db.AutoMigrate(&proto.ScoreV1{})
		return
	}

	if err = sqlitedb.CreateTable(score).Error; err != nil {
		return
	}

	result = &ScoreV1SQLiteRepositorier{}
	return
}