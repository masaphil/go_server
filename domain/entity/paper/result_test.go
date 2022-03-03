package paper

import (
	"reflect"
	"testing"
)

func Test_initResult(t *testing.T) {
	type args struct {
		a  float32
		ql *[]*Quiz
		h  float32
	}
	tests := []struct {
		name string
		args args
		want *Result
	}{
		{
			name: "test",
			args: args{
				a: 20,
				ql: &[]*Quiz{
					{},
					{},
				},
				h: 1.1,
			},
			want: &Result{
				Point:     nil,
				Deviation: nil,
				Volume:    2,
				Average:   20,
				Hyojun:    1.1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := initResult(tt.args.a, tt.args.ql, tt.args.h); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("initResult() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestResult_completeResult(t *testing.T) {
	f := float32(1.1)
	p := int(10)

	type fields struct {
		point     *int
		deviation *float32
		volume    int
		average   float32
		hyojun    float32
	}
	type args struct {
		p int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Result
	}{
		{
			name: "test",
			fields: fields{
				point:     nil,
				deviation: nil,
				volume:    2,
				average:   7.52,
				hyojun:    2.1,
			},
			args: args{
				p: 10,
			},
			want: &Result{
				Point:     &p,
				Deviation: &f,
				Volume:    2,
				Average:   7.52,
				Hyojun:    2.1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Result{
				Point:     tt.fields.point,
				Deviation: tt.fields.deviation,
				Volume:    tt.fields.volume,
				Average:   tt.fields.average,
				Hyojun:    tt.fields.hyojun,
			}
			if got := r.completeResult(tt.args.p); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Result.completeResult() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_calculateDeviation(t *testing.T) {
	type args struct {
		p      int
		hyojun float32
		ave    float32
	}
	tests := []struct {
		name string
		args args
		want *string
	}{
		{
			name: "60.5",
			args: args{
				p:      10,
				hyojun: 2.1,
				ave:    7.52,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := calculateDeviation(tt.args.p, tt.args.hyojun, tt.args.ave)
			println(*got)
		})
	}
}
