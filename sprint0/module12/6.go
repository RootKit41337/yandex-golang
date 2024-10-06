package main

import (
	"errors"
	"strconv"
)

func SumTwoIntegers(a, b string) (int, error) {
    intA, errA := strconv.Atoi(a)
    intB, errB := strconv.Atoi(b)

    if errA != nil || errB != nil {
        return 0, errors.New("invalid input, please provide two integers")
    }
    return intA + intB, nil
}