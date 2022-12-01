package utils

func Reduce[A, T any](s []T, f func(A, T) A, start A) A {
	acc := start
	for _, v := range s {
		acc = f(acc, v)
	}
	return acc
}
