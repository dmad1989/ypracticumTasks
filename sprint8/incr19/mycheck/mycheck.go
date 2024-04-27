package main

import (
	// "github.com/dmad1989/ypracticumTasks/tree/master/sprint8/incr19/errcheckanalyzer"
	"golang.org/x/tools/go/analysis/multichecker"
	"golang.org/x/tools/go/analysis/passes/printf"
	"golang.org/x/tools/go/analysis/passes/shadow"
	"golang.org/x/tools/go/analysis/passes/shift"
	"golang.org/x/tools/go/analysis/passes/structtag"
)

func main() {
	multichecker.Main(
		printf.Analyzer,
		shadow.Analyzer,
		structtag.Analyzer,
		shift.Analyzer,
		ErrCheckAnalyzer,
	)
}
