package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
)

/*
=== Утилита grep ===

Реализовать утилиту фильтрации (man grep)

Поддержать флаги:
-A - "after" печатать +N строк после совпадения
-B - "before" печатать +N строк до совпадения
-C - "context" (A+B) печатать ±N строк вокруг совпадения
-c - "count" (количество строк)
-i - "ignore-case" (игнорировать регистр)
-v - "invert" (вместо совпадения, исключать)
-F - "fixed", точное совпадение со строкой, не паттерн
-n - "line num", печатать номер строки

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

type grepFlags struct {
	after       int
	before      int
	context     int
	count       bool
	ignoreCase  bool
	invertMatch bool
	fixed       bool
	lineNum     bool
}

func grep(input string, searchTerm string, flags grepFlags, writer io.Writer) {
	// Remove the newline at the end of the input if it exists
	if len(input) > 0 && input[len(input)-1] == '\n' {
		input = input[:len(input)-1]
	}

	lines := strings.Split(input, "\n")
	matches := make(map[int]bool)
	count := 0

	pattern := searchTerm
	if flags.ignoreCase {
		pattern = "(?i)" + pattern
	}
	if !flags.fixed {
		pattern = ".*" + pattern + ".*"
	}
	re, err := regexp.Compile(pattern)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Invalid pattern:", err)
		return
	}

	for i, line := range lines {
		match := re.MatchString(line)
		if flags.invertMatch {
			match = !match
		}
		if match {
			count++
			// Handle context lines
			for j := i - flags.context; j <= i+flags.context; j++ {
				if j >= 0 && j < len(lines) {
					matches[j] = true
				}
			}
			// Handle lines before the match
			for j := i - flags.before; j < i; j++ {
				if j >= 0 {
					matches[j] = true
				}
			}
			// Handle lines after the match
			for j := i + 1; j <= i+flags.after; j++ {
				if j < len(lines) {
					matches[j] = true
				}
			}
		}
	}

	if flags.count {
		fmt.Fprint(writer, count)
	} else {
		for i, line := range lines {
			if matches[i] {
				if flags.lineNum {
					fmt.Fprintf(writer, "%d:%s\n", i+1, line)
				} else {
					fmt.Fprint(writer, line+"\n")
				}
			}
		}
	}
}

func main() {
	var flags grepFlags
	flag.IntVar(&flags.after, "A", 0, "print +N lines after match")
	flag.IntVar(&flags.before, "B", 0, "print +N lines before match")
	flag.IntVar(&flags.context, "C", 0, "print ±N lines around match")
	flag.BoolVar(&flags.count, "c", false, "print count of matching lines")
	flag.BoolVar(&flags.ignoreCase, "i", false, "ignore case")
	flag.BoolVar(&flags.invertMatch, "v", false, "select non-matching lines")
	flag.BoolVar(&flags.fixed, "F", false, "interpret pattern as fixed string")
	flag.BoolVar(&flags.lineNum, "n", false, "print line number with output lines")
	flag.Parse()

	if flag.NArg() < 1 {
		fmt.Fprintln(os.Stderr, "Usage: grep [OPTIONS] PATTERN [FILE]")
		os.Exit(1)
	}
	searchTerm := flag.Arg(0)
	filename := flag.Arg(1)

	var input string
	var err error
	if filename != "" {
		bytes, err := os.ReadFile(filename)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error reading file:", err)
			os.Exit(1)
		}
		input = string(bytes)
	} else {
		reader := bufio.NewReader(os.Stdin)
		input, err = reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error reading standard input:", err)
			os.Exit(1)
		}
	}
	grep(input, searchTerm, flags, os.Stdout)
}
