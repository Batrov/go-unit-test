package module4

import "testing"

func Test_SumNew(t *testing.T) {

	const (
		Success       = "Test_SumNew_Success"
		SuccessPatch1 = "Test_SumNew_SuccessPatch1"
		SuccessPatch2 = "Test_SumNew_SuccessPatch2"
	)

	type args struct {
		a int64
		b int64
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{Success, args{a: 5, b: 7}, ""},
		{SuccessPatch1, args{a: 5, b: 7}, "Patch1"},
		{SuccessPatch2, args{a: 5, b: 7}, "Patch2"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			switch tt.name {
			case SuccessPatch1:
				fooSumOrig := fooSum
				fooGetMessageOrig := fooGetMessage
				fooSum = func(a, b int64) int64 {
					return 12
				}

				fooGetMessage = func(id int64) string {
					return "Patch1"
				}

				defer func() {
					fooSum = fooSumOrig
					fooGetMessage = fooGetMessageOrig
				}()

			case SuccessPatch2:
				fooSumOrig := fooSum
				fooGetMessageOrig := fooGetMessage
				fooSum = func(a, b int64) int64 {
					return -1
				}

				fooGetMessage = func(id int64) string {
					return "Patch2"
				}

				defer func() {
					fooSum = fooSumOrig
					fooGetMessage = fooGetMessageOrig
				}()

			}

			if got := SumNew(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("SumNew() = %v, want %v", got, tt.want)
			}
		})
	}
}
