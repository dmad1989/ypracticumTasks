package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"
)

func main() {
	var procAttr os.ProcAttr
	procAttr.Env = os.Environ()
	// первые три элемента указывают на то, куда направлять
	// стандартный ввод, вывод и вывод ошибок
	procAttr.Files = []*os.File{os.Stdin, os.Stdout, os.Stderr}

	// для Windows следует использовать параметр -n вместо -с
	args := []string{"ping", "-n", "3", "yandex.ru"}
	// ищем полный путь к ping
	name, err := exec.LookPath(args[0])
	if err != nil {
		log.Fatal(err)
	}

	// функция для запуска процесса
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
	// запускаем и дожидаемся окончания
	proc := start()
	finish(proc)
	// запускаем и пробуем через секунду убить запущенный процесс
	proc = start()
	time.AfterFunc(time.Second, func() {
		proc.Kill()
	})
	finish(proc)
}
