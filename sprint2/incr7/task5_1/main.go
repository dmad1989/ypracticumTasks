package main

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/pelletier/go-toml"
	yaml "gopkg.in/yaml.v3"
)

type Data struct {
	ID     int    `toml:"id"`
	Name   string `toml:"name"`
	Values []byte `toml:"values"`
}

const yamlData = `
id: 101
name: Gopher
values:
- 11
- 22
- 33
`

func main() {
	var dataYaml Data
	// вставьте недостающий код
	// 1) десериализуйте yamlData в переменную типа Data
	dec := *yaml.NewDecoder(strings.NewReader(yamlData))
	if err := dec.Decode(&dataYaml); err != nil {
		panic(err)
	}
	buf := new(bytes.Buffer)
	enc := toml.NewEncoder(buf)
	if err := enc.Encode(&dataYaml); err != nil {
		panic(err)
	}

	fmt.Println(buf)
	// 2) преобразуйте полученную переменную в TOML
	// 3) выведите в консоль результат
	// ...
}
