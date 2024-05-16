package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func main() {
	resM := make(map[string]interface{})
	for _, v := range os.Environ() {
		env := strings.SplitN(v, "=", 2)
		if len(env) == 2 {
			if env[0] == "PATH" {
				resM[env[0]] = strings.Split(env[1], string(os.PathListSeparator))
			} else {
				resM[env[0]] = env[1]
			}
		}
	}
	out, err := json.MarshalIndent(resM, "", "   ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(out))
}
