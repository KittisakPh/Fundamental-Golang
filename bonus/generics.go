package bonus

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func minFloat64(a, b float64) float64 {
	if a < b {
		return a
	}
	return b
}

type Number interface {
	int | float64
}

func min[T Number](a, b T) T {
	if a < b {
		return a
	}
	return b
}

func Generics() {

}