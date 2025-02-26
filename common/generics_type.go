package common

type Frame interface {
	~int | ~int64 | ~string | ~bool | ~float32 | ~float64
}
