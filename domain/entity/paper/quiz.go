package paper

import (
	"errors"

	"github.com/quantum-box/skillforest_platform/go/services/test/domain/entity/source"
)

type QuizID string

//テスト答案問題
type Quiz struct {

	//問題文
	Statement string

	//選択肢
	Options []Option

	//正答
	Answer Option

	//回答されたOption
	Answered *Option

	//正解を見たかどうか
	IsReadAnswer bool
}

func NewQuiz(st string, op []Option, ans Option, writea Option, isRead bool) *Quiz {
	return &Quiz{
		Statement:    st,
		Options:      op,
		Answer:       ans,
		Answered:     &writea,
		IsReadAnswer: isRead,
	}
}

//from []source_quiz to []paper_quiz
func newQuizList(s []*source.SourceQuiz) []*Quiz {
	//問題の順番はランダム
	shuffled := source.ShuffleSourceQuizzes(s)

	slice := make([]*Quiz, 0)
	for _, sss := range shuffled {
		slice = append(slice, &Quiz{
			Statement: sss.Statement,
			//optionの順番シャッフル
			Options:      FromSourceOption(sss.Options),
			Answer:       Option(sss.Answer),
			Answered:     nil,
			IsReadAnswer: false,
		})
	}

	return slice
}

//回答を受け取ると、正答が開示される
//正答が開示されたら、回答は変更できない
func (q *Quiz) answerQuiz(op Option) error {
	if q.IsReadAnswer {
		return errors.New("正答確認後は選択を修正できません")
	}
	q.Answered = &op
	return nil
}

//isReadAnswerをtrueにするメソッド
func (q *Quiz) readAnswer() {
	q.IsReadAnswer = true
}

//答えと合っているか採点する
func (q *Quiz) checkAnswer() bool {
	for _, v := range q.Options {
		if *q.Answered == v {
			return true
		}
	}
	return false
}
