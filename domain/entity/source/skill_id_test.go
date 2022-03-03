package source

import (
	"reflect"
	"testing"
)

func Test_toSkillIDList(t *testing.T) {
	type args struct {
		data []string
	}
	tests := []struct {
		name string
		args args
		want []SkillID
	}{
		{
			name: "test",
			args: args{
				data: []string{"123", "345", "456"},
			},
			want: []SkillID{"123", "345", "456"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToSkillIDList(tt.args.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("toSkillIDList() = %v, want %v", got, tt.want)
			}
		})
	}
}
