package module1

import (
	"testing"
)

func Test_sum(t *testing.T) {
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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sum(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("sum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sumPointer(t *testing.T) {
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
			if got := sumPointer(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("sumPointer() = %v, want %v", got, tt.want)
			}
		})
	}
}
