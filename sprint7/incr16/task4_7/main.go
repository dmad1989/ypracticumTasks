package main

import (
	"io"
	"os"
	"reflect"
)

func ConditionalWrite(w io.Writer) (int, error) {
	// используем рефлексию для определения конкретного типа
	if wtype := reflect.TypeOf(w); wtype.String() == "*os.File" {
		// если это *os.File, предпринимаем действия
	}
	// действуем по-другому
	return 1, nil
}

func ConditionalWriteNonReflect(w io.Writer) (int, error) {

	if _, ok := w.(*os.File); ok {
	}
	return 1, nil
}
