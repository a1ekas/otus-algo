package main

import (
	"github.com/a1ekas/otus-algo/internal/runner"
	"github.com/a1ekas/otus-algo/internal/tests"
	// "github.com/a1ekas/otus-algo/internal/tests"
	// "github.com/a1ekas/otus-algo/internal/tests/strlen"
)

func main() {
	runner := runner.NewRunner(initTestsTable())
	runner.Run()
}

// Uncomment any test, make build and run it
func initTestsTable() []tests.Test {
	table := []tests.Test{
		// strlen.NewStringLenTest(),
		// tests.NewLuckyTicketsTest(),
	}

	return table
}
