package main

import (
	"sort"
	"strings"
)

/*
=== Поиск анаграмм по словарю ===

Напишите функцию поиска всех множеств анаграмм по словарю.
Например:
'пятак', 'пятка' и 'тяпка' - принадлежат одному множеству,
'листок', 'слиток' и 'столик' - другому.

Входные данные для функции: ссылка на массив - каждый элемент которого - слово на русском языке в кодировке utf8.
Выходные данные: Ссылка на мапу множеств анаграмм.
Ключ - первое встретившееся в словаре слово из множества
Значение - ссылка на массив, каждый элемент которого, слово из множества. Массив должен быть отсортирован по возрастанию.
Множества из одного элемента не должны попасть в результат.
Все слова должны быть приведены к нижнему регистру.
В результате каждое слово должно встречаться только один раз.

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

// sortRunes - вспомогательная функция для сортировки рун в строке.
func sortRunes(str string) string {
	r := []rune(str)
	sort.Slice(r, func(i, j int) bool {
		return r[i] < r[j]
	})
	return string(r)
}

// findAnagrams - функция для поиска анаграмм.
func findAnagrams(words []string) map[string][]string {
	anagrams := make(map[string][]string)
	seen := make(map[string]string)

	for _, word := range words {
		wordLower := strings.ToLower(word) // Приведение слова к нижнему регистру
		sortedWord := sortRunes(wordLower) // Сортировка рун в слове

		if key, found := seen[sortedWord]; found {
			// Если слово-анаграмма уже было найдено, добавляем его в массив по ключу
			anagrams[key] = append(anagrams[key], wordLower)
		} else {
			// Иначе, создаем новый ключ для этой анаграммы
			seen[sortedWord] = wordLower
			anagrams[wordLower] = []string{wordLower}
		}
	}

	// Удаление множеств анаграмм, содержащих только одно слово
	for key, group := range anagrams {
		if len(group) < 2 {
			delete(anagrams, key)
		} else {
			sort.Strings(anagrams[key]) // Сортировка множеств анаграмм
		}
	}

	return anagrams
}

func main() {
	// Пример использования функции findAnagrams.
	dictionary := []string{"пятак", "пятка", "тяпка", "листок", "слиток", "столик", "кирпич"}
	anagrams := findAnagrams(dictionary)

	for key, group := range anagrams {
		println(key, ":", strings.Join(group, ", "))
	}
}
