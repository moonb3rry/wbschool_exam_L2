package main

import (
	"reflect"
	"testing"
)

func TestFindAnagrams(t *testing.T) {
	testCases := []struct {
		name     string
		words    []string
		expected map[string][]string
	}{
		{
			name:  "стандартный тест",
			words: []string{"пятак", "пятка", "тяпка", "листок", "слиток", "столик", "кирпич"},
			expected: map[string][]string{
				"пятак":  {"пятак", "пятка", "тяпка"},
				"листок": {"листок", "слиток", "столик"},
			},
		},
		{
			name:     "пустой массив",
			words:    []string{},
			expected: map[string][]string{},
		},
		{
			name:     "нет анаграмм",
			words:    []string{"пятак", "листок", "кирпич"},
			expected: map[string][]string{},
		},
		{
			name:  "слова без анаграмм",
			words: []string{"пятак", "пятка", "тяпка", "абв", "где", "ёжз"},
			expected: map[string][]string{
				"пятак": {"пятак", "пятка", "тяпка"},
			},
		},
		{
			name:  "случай с учетом регистра",
			words: []string{"Пятак", "пЯтка", "Тяпка", "листок", "Слиток", "столИк"},
			expected: map[string][]string{
				"пятак":  {"пятак", "пятка", "тяпка"},
				"листок": {"листок", "слиток", "столик"},
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			result := findAnagrams(testCase.words)
			if !reflect.DeepEqual(result, testCase.expected) {
				t.Errorf("Test '%s' failed: expected %v, but got %v", testCase.name, testCase.expected, result)
			}
		})
	}
}
