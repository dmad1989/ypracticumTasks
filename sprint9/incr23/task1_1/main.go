package main

import (
	"log"
	"os"
	"os/exec"
)

func main() {
	cmdout := exec.Command(`echo`, "Hello, world!")
	stdout, err := cmdout.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}

	cmdin := exec.Command(`cat`)
	// указываем текущую консоль для стандартного вывода
	cmdin.Stdout = os.Stdout
	cmdin.Stdin = stdout

	if err = cmdout.Start(); err != nil {
		log.Fatal(err)
	}
	if err = cmdin.Start(); err != nil {
		log.Fatal(err)
	}
	cmdin.Wait()
}
