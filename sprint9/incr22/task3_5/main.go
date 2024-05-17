package main

import (
	"fmt"
	"log"
	"os"
)

const (
	envName = "MYAPP"
)

func main() {
	p := os.Getenv(envName)
	if p == "" {
		name, err := os.Executable()
		if err != nil {
			log.Fatal(err)
		}
		args := make([]string, 0)
		var procAttr os.ProcAttr
		os.Setenv(envName, name)
		procAttr.Env = os.Environ()
		procAttr.Files = []*os.File{os.Stdin, os.Stdout, os.Stderr}
		start := func() *os.Process {
			proc, err := os.StartProcess(name, args, &procAttr)
			if err != nil {
				log.Fatal(err)
			}
			return proc
		}
		// функция для ожидания конца работы и получения статуса процесса
		finish := func(proc *os.Process) {
			state, err := proc.Wait()
			if err != nil {
				log.Fatal(err)
			}
			// выводим pid процесса и его статус
			fmt.Println("STATUS", proc.Pid, state.ExitCode(), state.Exited(),
				state.Success(), state)
		}
		proc := start()
		finish(proc)
		return
	}
	fmt.Println(p)
}
