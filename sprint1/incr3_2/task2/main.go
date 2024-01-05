package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	chi "github.com/go-chi/chi/v5"
)

var cars = map[string]string{
	"id1": "Renault Logan",
	"id2": "Renault Duster",
	"id3": "BMW X6",
	"id4": "BMW M5",
	"id5": "VW Passat",
	"id6": "VW Jetta",
	"id7": "Audi A4",
	"id8": "Audi Q7",
}

// carsListFunc — вспомогательная функция для вывода всех машин.
func carsListFunc() []string {
	var list []string
	for _, c := range cars {
		list = append(list, c)
	}
	return list
}

// carFunc — вспомогательная функция для вывода определённой машины.
func carFunc(id string) string {
	if c, ok := cars[id]; ok {
		return c
	}
	return "unknown identifier " + id
}

func carsHandle(rw http.ResponseWriter, r *http.Request) {
	fmt.Println("carsHandle")
	carsList := carsListFunc()
	fmt.Println("carsList", carsList)
	io.WriteString(rw, strings.Join(carsList, ", "))
}

func carHandle(rw http.ResponseWriter, r *http.Request) {
	carID := chi.URLParam(r, "id")
	if carID == "" {
		http.Error(rw, "id param is missed", http.StatusBadRequest)
		return
	}
	rw.Write([]byte(carFunc(carID)))
}

func main() {
	r := chi.NewRouter()
	// определяем хендлер, который выводит все машины
	r.Get("/cars", carsHandle)
	// определяем хендлер, который выводит определённую машину
	r.Get("/car/{id}", carHandle)
	log.Fatal(http.ListenAndServe(":8080", r))
}
