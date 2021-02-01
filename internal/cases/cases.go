package cases

// Case общий интерфейс для тестовых заданий
type Case interface {
	// LoadTestData
	LoadTestData() error
	Check() error
}
