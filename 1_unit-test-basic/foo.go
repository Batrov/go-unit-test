package module1

func sum(a, b int64) int64 {
	return a + b
}

func sumPointer(a, b *int) int {
	return *a + *b
}
