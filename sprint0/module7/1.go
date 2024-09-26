package main

import (
	"errors"
)

func UnderLimit(nums []int, limit int, n int) ([]int, error) {
	// Проверка входных данных
	if nums == nil {
		return nil, errors.New("входной слайс не может быть nil")
	}
	if n <= 0 {
		return nil, errors.New("n должно быть больше 0")
	}

	// Инициализация результирующего слайса
	var result []int

	// Итерация по входному слайсу
	for _, num := range nums {
		// Проверка, является ли текущий элемент меньше limit
		if num < limit {
			// Добавление элемента в результирующий слайс
			result = append(result, num)
			// Проверка, достигнут ли предел n элементов
			if len(result) == n {
				break
			}
		}
	}

	return result, nil
}