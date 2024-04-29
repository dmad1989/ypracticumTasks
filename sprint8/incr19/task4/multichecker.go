package main

import (
	"golang.org/x/tools/go/analysis/multichecker"
	"golang.org/x/tools/go/analysis/passes/shadow"
)

func main() {
	multichecker.Main(
		shadow.Analyzer,
		ErrCheckAnalyzer)
}
