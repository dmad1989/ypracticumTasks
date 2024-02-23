package main

import (
	"fmt"

	"golang.org/x/sync/errgroup"
)

type Name string

var names = []Name{"Anna", "Ivan", "Fedor", "Katya", "Gleb"}

// Здесь напишите метод для Name
func (n Name) Hello() error {
	fmt.Println("Hello ", n)
	return nil
}

func main() {
	g := &errgroup.Group{}

	// Вставьте ваш код здесь

	for _, name := range names {
		g.Go(name.Hello)
	}

	g.Wait()
}
