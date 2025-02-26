package utils

func MaxPointer[T int | int8 | int64 | float32 | float64](x, y *T) *T {
	if x == nil || y == nil {
		return nil
	}
	t := max(*x, *y)
	return &t
}

func MinPointer[T int | int8 | int64 | float32 | float64](x, y *T) *T {
	if x == nil || y == nil {
		return nil
	}
	t := min(*x, *y)
	return &t
}

func SumPointer[T int | int8 | int64 | float32 | float64](x, y *T) *T {
	if x == nil && y == nil {
		return nil
	}

	if x == nil {
		return y
	}

	if y == nil {
		return x
	}
	t := *x + *y
	return &t
}
