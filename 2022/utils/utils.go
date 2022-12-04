package utils

import "strconv"

func Reduce[A, T any](s []T, f func(A, T) A, start A) A {
	acc := start
	for _, v := range s {
		acc = f(acc, v)
	}
	return acc
}

func MapToInt(s []string) ([]int, error) {
	var res []int
	for _, v := range s {
		num, err := strconv.Atoi(v)
		if err != nil {
			return nil, err
		}
		res = append(res, num)
	}

	return res, nil
}
