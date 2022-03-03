package source

import (
	valueobject "github.com/quantum-box/skillforest_platform/go/services/test/domain/value_object"
)

type SourceID string

//テストソース<<root>>
//テスト答案を生成する
type Source struct {
	ID          SourceID
	Title       string
	Statement   string
	QuizList    []*SourceQuiz
	SkillIdList []SkillID

	//平均正答数
	Average float32

	//TODO:標準偏差 valueobjectにする
	Hyojun float32
}

func New(title string, state string, quizzes []*SourceQuiz, skillIds []string, ave float32, hyo float32) *Source {

	//idをジェネレイト
	sid := SourceID(valueobject.GenerateNewID().String())

	return &Source{
		ID:          sid,
		Title:       title,
		Statement:   state,
		QuizList:    quizzes,
		SkillIdList: ToSkillIDList(skillIds),
		Average:     ave,
		Hyojun:      hyo,
	}

}

func NewSource(sid string, title string, state string, quizzes []*SourceQuiz, skillIds []string, ave float32, hyo float32) *Source {

	return &Source{
		ID:          SourceID(sid),
		Title:       title,
		Statement:   state,
		QuizList:    quizzes,
		SkillIdList: ToSkillIDList(skillIds),
		Average:     ave,
		Hyojun:      hyo,
	}

}
