package source

import (
	"errors"
	"math/rand"
)

type SourceQuiz struct {

	//タイトル
	Title string

	//問題文
	Statement string

	//選択肢
	Options []Option

	//正答
	Answer Option
}

func NewSourceQuiz(t string, s string, o []Option, a Option) (*SourceQuiz, error) {

	err := checkDuplicatedOptions(o)
	if err != nil {
		return nil, err
	}

	err = checkContains(o, a)
	if err != nil {
		return nil, err
	}

	//optionに正答が含まれているか確認
	if len(o) < 2 {
		return nil, errors.New("選択肢が少ないです")
	}

	return &SourceQuiz{
		Title:     t,
		Statement: s,
		Options:   o,
		Answer:    a,
	}, nil
}

//SourceQuizリストの要素をシャッフル
func ShuffleSourceQuizzes(data []*SourceQuiz) []*SourceQuiz {
	rand.Shuffle(len(data), func(i, j int) {
		data[i], data[j] = data[j], data[i]
	})
	return data
}
