package main

import (
	"fmt"
	"golang.org/x/net/html"
	"io"
	"net/http"
	"net/url"
	"os"
	"path"
	"path/filepath"
	"strings"
)

/*
=== Утилита wget ===

Реализовать утилиту wget с возможностью скачивать сайты целиком

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func main() {
	args := os.Args[1:]
	if len(args) < 1 {
		fmt.Println("Usage: wget <URL>")
		os.Exit(1)
	}
	websiteUrl := args[0]
	err := downloadWebsite(websiteUrl)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Ошибка при скачивании сайта: %s\n", err)
		os.Exit(1)
	}
}

// downloadWebsite скачивает страницы начиная с заданного URL.
func downloadWebsite(websiteUrl string) error {
	baseUrl, err := url.Parse(websiteUrl)
	if err != nil {
		return err
	}

	page, err := downloadPage(baseUrl.String())
	if err != nil {
		return err
	}

	links := getAllLinks(page)

	for _, link := range links {
		absUrl, err := baseUrl.Parse(link)
		if err != nil {
			fmt.Println("Невозможно обработать ссылку:", link)
			continue
		}
		_, err = downloadPage(absUrl.String())
		if err != nil {
			fmt.Println("Ошибка при скачивании ресурса:", absUrl.String())
		}
	}

	return nil
}

// downloadPage скачивает страницу по URL и возвращает её содержимое.
func downloadPage(pageUrl string) (string, error) {
	resp, err := http.Get(pageUrl)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	content := string(body)
	savePath := getSavePath(pageUrl)
	err = os.WriteFile(savePath, body, 0644)
	if err != nil {
		return "", err
	}

	return content, nil
}

// getAllLinks возвращает все ссылки найденные в HTML странице.
func getAllLinks(htmlContent string) []string {
	var links []string
	z := html.NewTokenizer(strings.NewReader(htmlContent))
	for {
		tt := z.Next()
		switch tt {
		case html.ErrorToken:
			return links
		case html.StartTagToken, html.SelfClosingTagToken:
			token := z.Token()
			if token.Data == "a" {
				for _, attr := range token.Attr {
					if attr.Key == "href" {
						links = append(links, attr.Val)
					}
				}
			}
			// Добавить обработку других тегов, таких как img, script и link
		}
	}
}

// getSavePath генерирует путь для сохранения файла на основе его URL.
func getSavePath(pageUrl string) string {
	parsedUrl, err := url.Parse(pageUrl)
	if err != nil {
		panic(err)
	}
	domain := parsedUrl.Host
	filePath := parsedUrl.Path
	if filePath == "" || strings.HasSuffix(filePath, "/") {
		filePath = path.Join(filePath, "index.html")
	}

	saveDir := filepath.Join("downloads", domain)
	os.MkdirAll(saveDir, os.ModePerm)

	savePath := filepath.Join(saveDir, filePath)
	os.MkdirAll(filepath.Dir(savePath), os.ModePerm)
	return savePath
}
