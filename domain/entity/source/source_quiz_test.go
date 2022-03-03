package source

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestNewSourceQuiz(t *testing.T) {
	type args struct {
		t string
		s string
		o []Option
		a Option
	}
	tests := []struct {
		name    string
		args    args
		want    *SourceQuiz
		wantErr bool
	}{
		{
			name: "選択肢2つ答え1つok",
			args: args{
				o: []Option{"aaa", "bbbb"},
				a: Option("aaa"),
			},
			want: &SourceQuiz{
				Options: []Option{"aaa", "bbbb"},
				Answer:  Option("aaa"),
			},
		},
		{
			name:    "選択肢1つerr",
			args:    args{o: []Option{"aaa"}},
			want:    nil,
			wantErr: true,
		},
		{
			name: "[]optionに重複した要素が存在する",
			args: args{
				o: []Option{"aaa", "aaa", "ccc"},
				a: Option("aaa"),
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "[]optionに正答が含まれていない",
			args: args{
				o: []Option{"aaa", "bbb", "ccc"},
				a: Option("zzz"),
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewSourceQuiz(tt.args.t, tt.args.s, tt.args.o, tt.args.a)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewSourceQuiz() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewSourceQuiz() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestShuffleSourceQuizzes(t *testing.T) {
	type args struct {
		data []*SourceQuiz
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "shufflesourcequiz",
			args: args{
				data: []*SourceQuiz{
					{
						Statement: "11111",
						Options:   []Option{"aaa", "bbb", "ccc"},
						Answer:    Option("aaa"),
					},
					{
						Statement: "2222",
						Options:   []Option{"aaa", "bbb", "ccc"},
						Answer:    Option("aaa"),
					},
					{
						Statement: "33333333",
						Options:   []Option{"aaa", "bbb", "ccc"},
						Answer:    Option("aaa"),
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ShuffleSourceQuizzes(tt.args.data)
			s, _ := json.Marshal(got)
			println(string(s))
		})
	}
}
