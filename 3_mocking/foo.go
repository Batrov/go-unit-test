package module3

/*
	To create mock using mockgen:
	mockgen -source=./3_mocking/foo.go -destination=./3_mocking/foo_mock.go -package=module3
*/

type FooInterface interface {
	sum(a, b int64) int64
	getMessage(id int64) string
}

type FooImpl struct {
	message1 string
	message2 string
}

type Bar struct {
	foo FooInterface
}

func (x *FooImpl) sum(a, b int64) int64 {
	return a + b
}

func (x *FooImpl) getMessage(id int64) string {
	if id > 0 {
		return x.message1
	}
	return x.message2
}

func (bar *Bar) sumNew(a, b int64) string {
	if bar.foo.sum(a, b) > 0 {
		return bar.foo.getMessage(1)
	}

	return bar.foo.getMessage(-1)
}
