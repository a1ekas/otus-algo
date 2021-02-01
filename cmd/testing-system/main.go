package main

import (
	"github.com/a1ekas/otus-algo/internal/cases"
	"github.com/a1ekas/otus-algo/internal/runner"
)

func main() {
	runner := runner.NewRunner(initTestsTable())
	runner.Run()
}

func initTestsTable() []cases.Case {
	table := []cases.Case{
		cases.NewStringLenTestCase(),
	}

	return table
}
