package runner

import (
	"log"

	"github.com/a1ekas/otus-algo/internal/tests"
)

// Runner структура для хранения тест кейсов.
type Runner struct {
	tests []tests.Test
}

// NewRunner конструктор Runner. Возвращает ссылку на структуру
func NewRunner(testsTable []tests.Test) *Runner {
	return &Runner{tests: testsTable}
}

// Run метод запускает проверку тест кейсов
func (r *Runner) Run() {
	log.Println("Runner is ran...")
	for _, t := range r.tests {
		t.Check()
	}
}
