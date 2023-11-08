package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

/*
=== Утилита cut ===

Принимает STDIN, разбивает по разделителю (TAB) на колонки, выводит запрошенные

Поддержать флаги:
-f - "fields" - выбрать поля (колонки)
-d - "delimiter" - использовать другой разделитель
-s - "separated" - только строки с разделителем

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func main() {
	fieldsFlag := flag.String("f", "", "Выбрать поля (колонки)")
	delimiterFlag := flag.String("d", "\t", "Использовать другой разделитель")
	separatedFlag := flag.Bool("s", false, "Только строки с разделителем")

	flag.Parse()

	fields := parseFields(*fieldsFlag)
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()
		if *separatedFlag && !strings.Contains(line, *delimiterFlag) {
			continue
		}

		columns := strings.Split(line, *delimiterFlag)
		selectedColumns := selectFields(columns, fields)
		fmt.Println(strings.Join(selectedColumns, *delimiterFlag))
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Ошибка чтения ввода:", err)
	}
}

// parseFields парсит строку с полями, возвращая слайс индексов полей.
func parseFields(fieldsStr string) []int {
	var fields []int
	for _, f := range strings.Split(fieldsStr, ",") {
		var field int
		fmt.Sscanf(f, "%d", &field)
		fields = append(fields, field-1) // Пользователь вводит поля начиная с 1, а не с 0
	}
	return fields
}

// selectFields возвращает только выбранные поля из входного слайса колонок.
func selectFields(columns []string, fields []int) []string {
	var result []string
	for _, field := range fields {
		if field >= 0 && field < len(columns) {
			result = append(result, columns[field])
		}
	}
	return result
}
