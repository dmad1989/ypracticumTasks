// Пакет check предназначен для проверки знаний студентов.
package check

import "unicode"

// Student хранит информацию о студенте.
type Student struct {
	Name  string // имя студента
	Year  int    // год рождения
	state int
}

// SetName устанавливает имя студента.
func (s *Student) SetName(name string) {
	r := []rune(name)
	r[0] = unicode.ToUpper(r[0])
	nameCapitalized := string(r)

	s.Name = nameCapitalized
}

// GetName получает имя студента.
func (s *Student) GetName() string {
	return s.Name
}
