package main

import "net/http"

type DBOperation interface {
	do(r *http.Request)
}

var Funcs = make(map[string]func(*http.Request)) // пустой интерфейс может принять любое значение

func DBInsert(r *http.Request) {
	// логика вставки
}

func DBDelete(r *http.Request) {
	// логика удаления
}

func main() {
	Funcs["DBInsert"] = DBInsert
	Funcs["DBDelete"] = DBDelete
	Funcs["DBChange"] = func(r *http.Request) {
		// логика изменения
	}
	Funcs["DBInsert"](new(http.Request))

	// ...
}
