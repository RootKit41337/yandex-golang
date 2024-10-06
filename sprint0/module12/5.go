package main

import (
	"errors"
	"fmt"
)

func IntToBinary(num int) (string, error) {
	if num >= 0 {
		return fmt.Sprintf("%b", num), nil

	}
	return "", errors.New("negative numbers are not allowed")
}