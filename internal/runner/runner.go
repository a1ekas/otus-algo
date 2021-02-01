package runner

import (
	"log"

	"github.com/a1ekas/otus-algo/internal/cases"
)

// Runner структура для хранения тест кейсов.
type Runner struct {
	cases []cases.Case
}

// NewRunner конструктор Runner. Возвращает ссылку на структуру
func NewRunner(testsTable []cases.Case) *Runner {
	return &Runner{cases: testsTable}
}

// Run метод запускает проверку тест кейсов
func (r *Runner) Run() {
	log.Println("Runner is ran...")
	for _, tcase := range r.cases {
		tcase.Check()
	}
}
