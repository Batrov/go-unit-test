package module3

/*
	To create mock using mockgen:
	mockgen -source=./3_mocking/foo.go -destination=./3_mocking/foo_mock.go -package=module3
*/

type FooInterface interface {
	Sum(a, b int64) int64
	GetMessage(id int64) string
}

type FooImpl struct {
	message1 string
	message2 string
}

type Bar struct {
	foo FooInterface
}

func (x *FooImpl) Sum(a, b int64) int64 {
	return a + b
}

func (x *FooImpl) GetMessage(id int64) string {
	if id > 0 {
		return x.message1
	}
	return x.message2
}

func (bar *Bar) SumNew(a, b int64) string {
	if bar.foo.Sum(a, b) > 0 {
		return bar.foo.GetMessage(1)
	}

	return bar.foo.GetMessage(-1)
}
