package main

import (
	"errors"
)

func GetCharacterAtPosition(str string, position int) (rune, error){
	if position > len(str)/2 || position < 0{
		return rune(0), errors.New("position out of range")
	}
	return []rune(str)[position], nil
}