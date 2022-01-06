package module4

var foo FooImpl
var fooSum = foo.Sum
var fooGetMessage = foo.GetMessage

type FooImpl struct {
	message1 string
	message2 string
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

func SumNew(a, b int64) string {
	if fooSum(a, b) > 0 {
		return fooGetMessage(1)
	}

	return fooGetMessage(-1)
}
