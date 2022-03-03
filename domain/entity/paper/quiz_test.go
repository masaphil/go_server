package paper

import (
	"reflect"
	"testing"

	"github.com/quantum-box/skillforest_platform/go/services/test/domain/entity/source"
)

//answerOptionのテスト
func TestQuiz_answerOption(t *testing.T) {
	type fields struct {
		statement     string
		Options       []Option
		answer        Option
		writtenAnswer *Option
		isReadAnswer  bool
	}
	type args struct {
		op Option
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "最初の回答",
			fields: fields{
				isReadAnswer: false,
			},
			args: args{Option("回答")},
		},
		{
			name: "正答見たあとの回答修正はエラー",
			fields: fields{
				isReadAnswer:  true,
				writtenAnswer: nil,
			},
			args:    args{Option("修正回答")},
			wantErr: true,
		},
		{
			name: "正答を見る前の回答修正はok",
			fields: fields{
				isReadAnswer:  false,
				writtenAnswer: nil,
			},
			args: args{Option("修正回答")},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Quiz{
				Statement:    tt.fields.statement,
				Options:      tt.fields.Options,
				Answer:       tt.fields.answer,
				Answered:     tt.fields.writtenAnswer,
				IsReadAnswer: tt.fields.isReadAnswer,
			}
			if err := p.answerQuiz(tt.args.op); (err != nil) != tt.wantErr {
				t.Errorf("PaperQuiz.answerOption() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_newQuizList(t *testing.T) {
	type args struct {
		s []*source.SourceQuiz
	}
	tests := []struct {
		name string
		args args
		want []*Quiz
	}{
		{
			name: "test",
			args: args{
				s: []*source.SourceQuiz{
					{
						Statement: "aaaaaaaa問題文です。",
						Options:   []source.Option{"aaa", "bbb", "ccc"},
						Answer:    source.Option("aaa"),
					},
					{
						Statement: "bbbbbb問題文です。",
						Options:   []source.Option{"aaa", "bbb", "ccc"},
						Answer:    source.Option("aaa"),
					},
					{
						Statement: "ccc問題文です。",
						Options:   []source.Option{"aaa", "bbb", "ccc"},
						Answer:    source.Option("aaa"),
					},
				},
			},
			want: []*Quiz{
				{
					Statement:    "aaaaaaaa問題文です。",
					Options:      []Option{"aaa", "bbb", "ccc"},
					Answer:       Option("aaa"),
					Answered:     nil,
					IsReadAnswer: false,
				},
				/*
					{
						statement:     "bbbbbb問題文です。",
						Options:       []Option{"aaa", "bbb", "ccc"},
						answer:        Option("aaa"),
						writtenAnswer: nil,
						isReadAnswer:  false,
					},
				*/
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newQuizList(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newPaperQuizList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQuiz_checkAnswer(t *testing.T) {
	ans := Option("aaa")
	type fields struct {
		statement     string
		Options       []Option
		answer        Option
		writtenAnswer *Option
		isReadAnswer  bool
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "test",
			fields: fields{
				statement:     "aaaaaaaa問題文です。",
				Options:       []Option{"aaa", "bbb", "ccc"},
				answer:        Option("aaa"),
				writtenAnswer: &ans,
				isReadAnswer:  false,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &Quiz{
				Statement:    tt.fields.statement,
				Options:      tt.fields.Options,
				Answer:       tt.fields.answer,
				Answered:     tt.fields.writtenAnswer,
				IsReadAnswer: tt.fields.isReadAnswer,
			}
			if got := q.checkAnswer(); got != tt.want {
				t.Errorf("Quiz.checkAnswer() = %v, want %v", got, tt.want)
			}
		})
	}
}
