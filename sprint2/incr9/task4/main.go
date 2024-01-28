package settings

import (
	"encoding/json"
	"os"
)

// Settings содержит настройки приложения.
type Settings struct {
	Port int    `json:"port"`
	Host string `json:"host"`
}

// Save сохраняет настройки в файле fname.
func (settings Settings) Save(fname string) error {
	// сериализуем структуру в JSON формат
	data, err := json.MarshalIndent(settings, "", "   ")
	if err != nil {
		return err
	}
	// сохраняем данные в файл
	return os.WriteFile(fname, data, 0666)
}

// Load читает настройки из файла fname.
func (settings *Settings) Load(fname string) error {
	b, err := os.ReadFile(fname)
	if err != nil {
		return err
	}

	return json.Unmarshal(b, settings)
}
