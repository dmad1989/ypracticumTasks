package main

import "strings"

type Modifier interface {
	Modify() string
}

type Original struct {
	Value string
}

func (o *Original) Modify() string {
	return o.Value
}

// Upper возвращает строку в верхнем регистре.
type Upper struct {
	modifier Modifier
}

func (u *Upper) Modify() string {
	return strings.ToUpper(u.modifier.Modify())
}

// Replace заменяет строки old на new.
type Replace struct {
	modifier Modifier
	old      string
	new      string
}

// добавьте метод Modify для *Replace
// он должен заменять old на new
// ...

func (r *Replace) Modify() string {
	return strings.ReplaceAll(r.modifier.Modify(), r.old, r.new)
}
