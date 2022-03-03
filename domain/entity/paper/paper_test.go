package paper

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"

	"github.com/quantum-box/skillforest_platform/go/services/test/domain/entity/source"
	"github.com/quantum-box/skillforest_platform/go/services/test/domain/entity/tester"
)

func TestPaper_Answer(t *testing.T) {
	sample := tester.TesterID("123")
	type fields struct {
		ID       PaperID
		TesterId *tester.TesterID
		SourceId source.SourceID
		QuizList []*Quiz
		Result   *Result
	}
	type args struct {
		index int
		o     Option
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "test",
			fields: fields{
				ID:       "123",
				TesterId: &sample,
				SourceId: "123",
				QuizList: []*Quiz{
					{
						Statement:    "11問題文です。",
						Options:      []Option{"aaa", "bbb", "ccc"},
						Answer:       Option("aaa"),
						Answered:     nil,
						IsReadAnswer: false,
					},
					{
						Statement:    "22問題文です。",
						Options:      []Option{"aaa", "bbb", "ccc"},
						Answer:       Option("aaa"),
						Answered:     nil,
						IsReadAnswer: false,
					},
				},
				Result: nil,
			},
			args: args{
				index: 0,
				o:     Option("aaa"),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Paper{
				ID:       tt.fields.ID,
				TesterId: tt.fields.TesterId,
				SourceId: tt.fields.SourceId,
				QuizList: tt.fields.QuizList,
				Result:   tt.fields.Result,
			}
			if err := p.Answer(tt.args.index, tt.args.o); (err != nil) != tt.wantErr {
				t.Errorf("Paper.Answer() error = %v, wantErr %v", err, tt.wantErr)
			}
			//indexで指定したquizに回答が入っているか確認できる
			//s, _ := json.Marshal(p.QuizList[tt.args.index].writtenAnswer)
			//fmt.Println(string(s))
			ta, _ := json.Marshal(p.QuizList[0].Answered)
			fmt.Println(string(ta))
		})
	}
}

func TestPaper_MakeResult(t *testing.T) {
	sample := tester.TesterID("123")
	//&Option("aaa")がダメだった
	ans := Option("aaa")
	type fields struct {
		ID       PaperID
		TesterId *tester.TesterID
		SourceId source.SourceID
		QuizList []*Quiz
		Result   *Result
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{
			name: "test",
			fields: fields{
				ID:       "123",
				TesterId: &sample,
				SourceId: "123",
				QuizList: []*Quiz{
					{
						Statement:    "11問題文です。",
						Options:      []Option{"aaa", "bbb", "ccc"},
						Answer:       Option("aaa"),
						Answered:     &ans,
						IsReadAnswer: false,
					},
					{
						Statement:    "22問題文です。",
						Options:      []Option{"aaa", "bbb", "ccc"},
						Answer:       Option("aaa"),
						Answered:     &ans,
						IsReadAnswer: false,
					},
				},
				Result: nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Paper{
				ID:       tt.fields.ID,
				TesterId: tt.fields.TesterId,
				SourceId: tt.fields.SourceId,
				QuizList: tt.fields.QuizList,
				Result:   tt.fields.Result,
			}
			got := p.MakeResult()
			s, _ := json.Marshal(got.Result.Point)
			println(string(s))
		})
	}
}

func TestNewPaperFromSource(t *testing.T) {
	sample := tester.TesterID("123")
	type args struct {
		ti     *tester.TesterID
		source *source.Source
	}
	tests := []struct {
		name string
		args args
		want *Paper
	}{
		{
			name: "test",
			args: args{
				ti: &sample,
				source: &source.Source{
					ID: "12123345",
					QuizList: []*source.SourceQuiz{
						{
							Statement: "aaaaaaaa問題文です。",
							Options:   []source.Option{"111", "bbb", "ccc"},
							Answer:    source.Option("aaa"),
						},
						{
							Statement: "bbbbbb問題文です。",
							Options:   []source.Option{"aaa", "bbb", "ccc"},
							Answer:    source.Option("aaa"),
						},
					},
					SkillIdList: []source.SkillID{"aaa", "www"},
					Average:     20,
					Hyojun:      1.1,
				},
			},
			want: &Paper{
				ID:       "",
				TesterId: &sample,
				SourceId: "12123345",
				QuizList: []*Quiz{
					{
						Statement:    "aaaaaaaa問題文です。",
						Options:      []Option{"aaa", "bbb", "ccc"},
						Answer:       Option("aaa"),
						Answered:     nil,
						IsReadAnswer: false,
					},
					{
						Statement:    "bbbbbb問題文です。",
						Options:      []Option{"aaa", "bbb", "ccc"},
						Answer:       Option("aaa"),
						Answered:     nil,
						IsReadAnswer: false,
					},
				},
				Result: nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewPaperFromSource(tt.args.ti, tt.args.source); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewPaperFromSource() = %v, want %v", *got.QuizList[0], tt.want)
			}
		})
	}
}
