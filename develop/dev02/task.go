package main

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

/*
=== Задача на распаковку ===

Создать Go функцию, осуществляющую примитивную распаковку строки, содержащую повторяющиеся символы / руны, например:
	- "a4bc2d5e" => "aaaabccddddde"
	- "abcd" => "abcd"
	- "45" => "" (некорректная строка)
	- "" => ""
Дополнительное задание: поддержка escape - последовательностей
	- qwe\4\5 => qwe45 (*)
	- qwe\45 => qwe44444 (*)
	- qwe\\5 => qwe\\\\\ (*)

В случае если была передана некорректная строка функция должна возвращать ошибку. Написать unit-тесты.

Функция должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func unpack(s string) (string, error) {
	var result strings.Builder
	var prevRune rune
	escaped := false

	for i, r := range s {
		if unicode.IsDigit(r) && !escaped {
			if prevRune == 0 || unicode.IsDigit(prevRune) {
				return "", errors.New("invalid string")
			}
			count, _ := strconv.Atoi(string(r))
			result.WriteString(strings.Repeat(string(prevRune), count-1))
		} else {
			if r != '\\' || escaped {
				result.WriteRune(r)
				escaped = false
			} else if i+1 < len(s) && s[i+1] == '\\' {
				escaped = true
			} else {
				escaped = !escaped
				continue
			}
		}
		prevRune = r
	}

	if escaped {
		return "", errors.New("invalid string")
	}

	return result.String(), nil
}

func main() {
}
