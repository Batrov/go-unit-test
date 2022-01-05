package module4

var foo FooImpl
var fooSum = foo.sum
var fooGetMessage = foo.getMessage

type FooImpl struct {
	message1 string
	message2 string
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

func sumNew(a, b int64) string {
	if fooSum(a, b) > 0 {
		return fooGetMessage(1)
	}

	return fooGetMessage(-1)
}
