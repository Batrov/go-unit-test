package module2

var a int64

/*
	Execution sequence:
	foo.go.init() -> foo_test.go.init() -> foo_test.go.TestMain()
*/

func init() {
	a = 10
}

func Sum(b int64) int64 {
	return a + b
}
