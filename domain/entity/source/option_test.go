package source

import (
	"reflect"
	"testing"
)

func Test_checkDuplicatedOptions(t *testing.T) {
	type args struct {
		ol []Option
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "重複してないとき",
			args: args{
				ol: []Option{"aaa", "bbb", "ccc"},
			},
			wantErr: false,
		},
		{
			name: "重複してるとき",
			args: args{
				ol: []Option{"aaa", "aaa", "ccc"},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := checkDuplicatedOptions(tt.args.ol); (err != nil) != tt.wantErr {
				t.Errorf("checkDuplicatedOptions() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
func Test_checkContains(t *testing.T) {
	type args struct {
		ol []Option
		a  Option
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "答えを含んでいるとき",
			args: args{
				ol: []Option{"aaa", "bbb", "ccc"},
				a:  Option("aaa"),
			},
			wantErr: false,
		},
		{
			name: "答えを含んでいないとき",
			args: args{
				ol: []Option{"aaa", "bbb", "ccc"},
				a:  Option("ddd"),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := checkContains(tt.args.ol, tt.args.a); (err != nil) != tt.wantErr {
				t.Errorf("checkContains() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestShuffleSourceOptions(t *testing.T) {
	type args struct {
		data []Option
	}
	tests := []struct {
		name string
		args args
		want []Option
	}{
		{
			name: "test",
			args: args{
				data: []Option{"bbb", "aaa", "ccc"},
			},
			want: []Option{"aaa", "bbb", "ccc"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ShuffleSourceOptions(tt.args.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ShuffleSourceOptions() = %v, want %v", got, tt.want)
				println(got)
			}
		})
	}
}

func TestToOptionList(t *testing.T) {
	type args struct {
		data []string
	}
	tests := []struct {
		name string
		args args
		want []Option
	}{
		{
			name: "test",
			args: args{
				data: []string{"aaa", "bbb", "ccc"},
			},
			want: []Option{"aaa", "bbb", "ccc"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToOptionList(tt.args.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToOptionList() = %v, want %v", got, tt.want)
			}
		})
	}
}
