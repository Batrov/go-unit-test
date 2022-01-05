package module2

import "testing"

func TestMain(m *testing.M) {
	a = 2
}

func Test_sum(t *testing.T) {
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
			if got := sum(tt.args.b); got != tt.want {
				t.Errorf("sum() = %v, want %v", got, tt.want)
			}
		})
	}
}
