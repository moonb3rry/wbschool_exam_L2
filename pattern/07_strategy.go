package pattern

import (
	"errors"
	"fmt"
	"strings"
)

/*
	Реализовать паттерн «стратегия».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Strategy_pattern
*/

/* Шаблон «Стратегия» позволяет при выполнении выбирать поведение алгоритма. */

// SearchStrategy - интерфейс для различных алгоритмов поиска
type SearchStrategy interface {
	Search(data []string, term string) (int, error)
}

// LinearSearchStrategy - конкретная реализация SearchStrategy для линейного поиска
type LinearSearchStrategy struct{}

func (l *LinearSearchStrategy) Search(data []string, term string) (int, error) {
	for i, item := range data {
		if item == term {
			return i, nil
		}
	}
	return -1, errors.New("элемент не найден")
}

// BinarySearchStrategy - конкретная реализация SearchStrategy для бинарного поиска
type BinarySearchStrategy struct{}

func (b *BinarySearchStrategy) Search(data []string, term string) (int, error) {
	low, high := 0, len(data)-1
	for low <= high {
		mid := (low + high) / 2
		midValue := data[mid]
		if midValue == term {
			return mid, nil
		} else if strings.Compare(midValue, term) < 0 {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1, errors.New("элемент не найден")
}

// SearchContext - содержит стратегию и предоставляет метод для выполнения поиска
type SearchContext struct {
	strategy SearchStrategy
}

func (sc *SearchContext) SetStrategy(strategy SearchStrategy) {
	sc.strategy = strategy
}

func (sc *SearchContext) ExecuteSearch(data []string, term string) (int, error) {
	return sc.strategy.Search(data, term)
}

func main() {
	data := []string{"apple", "banana", "cherry", "date", "fig", "grape"}
	context := &SearchContext{}

	// Устанавливаем стратегию линейного поиска
	context.SetStrategy(&LinearSearchStrategy{})
	index, err := context.ExecuteSearch(data, "cherry")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Линейный поиск: найдено по индексу %d\n", index)
	}

	// Устанавливаем стратегию бинарного поиска
	context.SetStrategy(&BinarySearchStrategy{})
	//index, err = context.ExecuteSearch(data, "pineapple")
	index, err = context.ExecuteSearch(data, "grape")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Бинарный поиск: найдено по индексу %d\n", index)
	}
}
