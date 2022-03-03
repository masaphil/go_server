package paper

import (
	"errors"

	"github.com/quantum-box/skillforest_platform/go/services/test/domain/entity/source"
	"github.com/quantum-box/skillforest_platform/go/services/test/domain/entity/tester"
	valueobject "github.com/quantum-box/skillforest_platform/go/services/test/domain/value_object"
)

type PaperID string

/*テスト答案<<root>>
テストの精度を上げるため同じテストは一回しか受けられない
試験時間は一問につき１分とする
中断はなし。最初から*/
//PaperQuizを答えるごとに上書きする
type Paper struct {
	ID       PaperID
	TesterId *tester.TesterID
	SourceId source.SourceID
	QuizList []*Quiz
	Result   *Result
}

func New(pi PaperID, ti *tester.TesterID, si source.SourceID, ql []*Quiz, re *Result) *Paper {
	return &Paper{
		ID:       pi,
		TesterId: ti,
		SourceId: si,
		QuizList: ql,
		Result:   re,
	}
}

//sourceとtesteridを受け取ってpaperを返す関数
func NewPaperFromSource(ti *tester.TesterID, source *source.Source) *Paper {

	//idをジェネレイト
	pid := PaperID(valueobject.GenerateNewID().String())

	//SourceQuiz から PaperQuizに変換するメソッド
	ql := newQuizList(source.QuizList)

	//resultの初期化
	re := initResult(source.Average, &ql, source.Hyojun)

	return New(pid, ti, source.ID, ql, re)
}

//問題に回答するメソッド
func (p *Paper) Answer(index int, o Option) error {
	if index < 0 || index > len(p.QuizList)-1 {
		return errors.New("indexが存在しません")
	}

	err := p.QuizList[index].answerQuiz(o)
	if err != nil {
		return err
	}

	return nil
}

//テストを終了するメソッド
func (p *Paper) FinishTest() *Paper {
	p.MakeResult()
	return p
}

//採点して結果をpaperに格納
func (p *Paper) MakeResult() *Paper {
	cnt := 0
	for _, q := range p.QuizList {
		if q.checkAnswer() {
			cnt += 1
		}
	}

	//resultの結果を作るメソッドを呼び出す
	p.Result = p.Result.completeResult(cnt)

	return p
}
