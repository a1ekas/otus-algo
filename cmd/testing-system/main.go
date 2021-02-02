package main

import (
	"github.com/a1ekas/otus-algo/internal/runner"
	"github.com/a1ekas/otus-algo/internal/tests"
)

func main() {
	runner := runner.NewRunner(initTestsTable())
	runner.Run()
}

func initTestsTable() []tests.Test {
	table := []tests.Test{
		tests.NewStringLenTest(),
		tests.NewLuckyTicketsTest(),
	}

	return table
}
