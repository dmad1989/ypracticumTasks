package main

import (
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
	carsList := carsListFunc()
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

func brandHandle(rw http.ResponseWriter, r *http.Request) {
	carsList := carsByBrand(rw, r)
	rw.Write([]byte(strings.Join(carsList, ", ")))
}

func carsByBrand(rw http.ResponseWriter, r *http.Request) (cars []string) {
	carBrand := chi.URLParam(r, "brand")
	if carBrand == "" {
		http.Error(rw, "brand param is missed", http.StatusBadRequest)
		return
	}
	for _, car := range carsListFunc() {
		if strings.ToUpper(carBrand) == strings.ToUpper(strings.Split(car, " ")[0]) {
			cars = append(cars, car)
		}
	}
	return
}

func modelHandle(rw http.ResponseWriter, r *http.Request) {
	var cars []string
	carModel := chi.URLParam(r, "model")
	for _, car := range carsByBrand(rw, r) {
		if strings.ToUpper(carModel) == strings.ToUpper(strings.Split(car, " ")[1]) {
			cars = append(cars, car)
		}
	}
	rw.Write([]byte(strings.Join(cars, ", ")))
}

func main() {
	r := chi.NewRouter()
	r.Route("/cars", func(r chi.Router) {
		r.Get("/", carsHandle)
		r.Route("/{brand}", func(r chi.Router) {
			r.Get("/", brandHandle)
			r.Get("/{model}", modelHandle)
		})
	})

	r.Get("/car/{id}", carHandle)

	log.Fatal(http.ListenAndServe(":8080", r))
}
