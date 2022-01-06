package module2

import (
	"os"
	"testing"
)

func init() {
}

func TestMain(m *testing.M) {
	a = 2
	exitVal := m.Run()

	os.Exit(exitVal)
}

func Test_Sum(t *testing.T) {
	type args struct {
		b int64
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{"Success", args{b: 3}, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Sum(tt.args.b); got != tt.want {
				t.Errorf("Sum() = %v, want %v", got, tt.want)
			}
		})
	}
}
