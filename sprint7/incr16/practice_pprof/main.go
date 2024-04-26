package main

import (
	"net/http"
	_ "net/http/pprof" // подключаем пакет pprof
)

const (
	addr    = ":8080"  // адрес сервера
	maxSize = 10000000 // будем растить слайс до 10 миллионов элементов
)

func foo() {
	// полезная нагрузка
	for {
		var s []int
		for i := 0; i < maxSize; i++ {
			s = append(s, i)
		}
	}
}

func main() {
	// создаём файл журнала профилирования cpu
	// fcpu, err := os.Create(`cpu.profile`)
	// if err != nil {
	// 	panic(err)
	// }
	// defer fcpu.Close()
	// if err := pprof.StartCPUProfile(fcpu); err != nil {
	// 	panic(err)
	// }
	// defer pprof.StopCPUProfile()

	go foo()
	http.ListenAndServe(addr, nil)
	// time.Sleep(10 * time.Second)

	// создаём файл журнала профилирования памяти
	// fmem, err := os.Create(`mem.profile`)
	// if err != nil {
	// 	panic(err)
	// }
	// defer fmem.Close()
	// runtime.GC() // получаем статистику по использованию памяти
	// if err := pprof.WriteHeapProfile(fmem); err != nil {
	// 	panic(err)
	// }
}
