package main

import (
	"encoding/json"
	"os"
	"path/filepath"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/multichecker"
	"golang.org/x/tools/go/analysis/passes/printf"
	"golang.org/x/tools/go/analysis/passes/shadow"
	"golang.org/x/tools/go/analysis/passes/structtag"
	"honnef.co/go/tools/staticcheck"
)

const confFileName = "config.json"

type ConfigData struct {
	Staticcheck []string
}

func main() {
	appfile, err := os.Executable()
	if err != nil {
		panic(err)
	}
	data, err := os.ReadFile(filepath.Join(filepath.Dir(appfile), confFileName))
	if err != nil {
		panic(err)
	}
	var cfg ConfigData
	if err = json.Unmarshal(data, &cfg); err != nil {
		panic(err)
	}
	mychecks := []*analysis.Analyzer{
		ErrCheckAnalyzer,
		printf.Analyzer,
		shadow.Analyzer,
		structtag.Analyzer,
	}

	checks := make(map[string]bool)
	for _, v := range cfg.Staticcheck {
		checks[v] = true
	}
	// добавляем анализаторы из staticcheck, которые указаны в файле конфигурации
	for _, v := range staticcheck.Analyzers {
		if checks[v.Analyzer.Name] {
			mychecks = append(mychecks, v.Analyzer)
		}
	}
	multichecker.Main(
		mychecks...)
}
