package module1

func Sum(a, b int64) int64 {
	return a + b
}

func SumPointer(a, b *int) int {
	return *a + *b
}
