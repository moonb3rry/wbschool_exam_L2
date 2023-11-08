package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

/*
=== Утилита sort ===

Отсортировать строки (man sort)
Основное

Поддержать ключи

-k — указание колонки для сортировки
-n — сортировать по числовому значению
-r — сортировать в обратном порядке
-u — не выводить повторяющиеся строки

Дополнительное

Поддержать ключи

-M — сортировать по названию месяца
-b — игнорировать хвостовые пробелы
-c — проверять отсортированы ли данные
-h — сортировать по числовому значению с учётом суффиксов

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

type sortOptions struct {
	column  int
	numeric bool
	reverse bool
	unique  bool
}

type line struct {
	data    string
	columns []string
}

type byColumn struct {
	lines []line
	opts  sortOptions
}

func (s byColumn) Len() int {
	return len(s.lines)
}

func (s byColumn) Swap(i, j int) {
	s.lines[i], s.lines[j] = s.lines[j], s.lines[i]
}

func (s byColumn) Less(i, j int) bool {
	less := func(i, j int) bool {
		// Сортировка по числовому значению
		if s.opts.numeric {
			numI, errI := strconv.Atoi(s.lines[i].columns[s.opts.column])
			numJ, errJ := strconv.Atoi(s.lines[j].columns[s.opts.column])
			if errI == nil && errJ == nil {
				return numI < numJ
			}
		}
		return s.lines[i].columns[s.opts.column] < s.lines[j].columns[s.opts.column]
	}

	if s.opts.reverse {
		return less(j, i)
	}
	return less(i, j)
}

func main() {
	// Флаги командной строки
	k := flag.Int("k", 1, "column")
	n := flag.Bool("n", false, "numeric sort")
	r := flag.Bool("r", false, "reverse sort")
	u := flag.Bool("u", false, "unique sort")

	flag.Parse()

	opts := sortOptions{
		column:  *k - 1, // Пользовательские колонки начинаются с 1
		numeric: *n,
		reverse: *r,
		unique:  *u,
	}

	lines, err := readLines(os.Stdin)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	sortLines(lines, opts)
	writeLines(lines, os.Stdout, opts)
}

// Функция для чтения строк из io.Reader
func readLines(r io.Reader) ([]line, error) {
	scanner := bufio.NewScanner(r)
	var lines []line

	for scanner.Scan() {
		text := scanner.Text()
		columns := strings.Fields(text)
		lines = append(lines, line{data: text, columns: columns})
	}
	return lines, scanner.Err()
}

// Функция для сортировки строк
func sortLines(lines []line, opts sortOptions) {
	sorter := byColumn{lines: lines, opts: opts}
	if opts.reverse {
		sort.Sort(sort.Reverse(sorter))
	} else {
		sort.Sort(sorter)
	}
}

// Функция для записи строк в io.Writer
func writeLines(lines []line, w io.Writer, opts sortOptions) {
	seen := make(map[string]bool)
	for _, ln := range lines {
		if opts.unique {
			if _, found := seen[ln.data]; found {
				continue
			}
			seen[ln.data] = true
		}
		fmt.Fprintln(w, ln.data)
	}
}
