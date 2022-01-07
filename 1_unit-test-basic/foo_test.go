package module1

import (
	"testing"
)

func Test_Sum(t *testing.T) {
	type args struct {
		a int64
		b int64
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{"Success", args{a: 2, b: 3}, 5},
		{"Success1", args{a: 3, b: 5}, 8},
		{"Success2", args{a: 6, b: 3}, 9},
		{"Success3", args{a: 2, b: 8}, 10},
		{"Success4", args{a: 2, b: 9}, 11},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Sum(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("sum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_SumPointer(t *testing.T) {
	a1 := 1
	b1 := 2
	type args struct {
		a *int
		b *int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Success", args{a: &a1, b: &b1}, 3},
		// {"Panic", args{a: &a1}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SumPointer(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("sumPointer() = %v, want %v", got, tt.want)
			}
		})
	}
}
