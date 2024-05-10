package main

import (
	"fmt"
	"log"
	"strings"

	"golang.org/x/net/html"
)

func GetLink(n *html.Node) {
	// проверка элемента на ссылку
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				fmt.Println(a.Val, n.FirstChild.Data)
				break
			}
		}
	}
	// рекурсивный обход дерева
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		GetLink(c)
	}
}

func main() {
	input := `<p>Ссылки:</p><ul><li><a href="https://ya.ru">Яндекс</a></li>
           <li><a href="https://praktikum.yandex.ru">Яндекс Практикум</a></li></ul>`
	// методы пакета могут токенизировать и парсить любой io.Reader
	// здесь конструируем io.Reader из строки
	page, err := html.Parse(strings.NewReader(input))
	if err != nil {
		log.Fatal(err)
	}
	GetLink(page)
}
