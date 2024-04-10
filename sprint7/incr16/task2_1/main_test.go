package main

import "testing"

func TestModifier(t *testing.T) {
	original := &Original{Value: "Привет, гофер!"}
	replace := &Replace{modifier: original, old: "гофер", new: "мир"} // инициализируйте поля нужными значениями
	// ...

	upper := &Upper{modifier: replace} // инициализируйте поле нужным значением
	// ...

	if upper.Modify() != "ПРИВЕТ, МИР!" {
		t.Errorf(`get %s`, upper.Modify())
	}
}
