package main

import (
	"flag"
	"fmt"
	"go/parser"
	"go/token"
	"strings"
)

func main() {
	flag.Parse()
	fset := token.NewFileSet()
	for _, fn := range flag.Args() {
		f, err := parser.ParseFile(fset, fn, nil, parser.ParseComments)
		if err != nil {
			fmt.Printf("Failed to parse file %v\n", fn)
			continue
		}
		for _, gr := range f.Comments {
			for _, c := range gr.List {
				if strings.HasPrefix(c.Text, "//TODO") {
					fmt.Println(fset.Position(c.Slash).String(), c.Text)
				}
				//TODO: Добавить обработку комментариев NOTICE
			}
		}
	}
	//TODO: Добавить новых возможностей
}
