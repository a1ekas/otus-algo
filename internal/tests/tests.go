package tests

// Test общий интерфейс для тестовых заданий
type Test interface {
	// LoadTestData ...
	LoadTestData()
	// Check ...
	Check()
}

// Case ...
type Case struct {
	ID           int
	InData       interface{}
	ExpectedData interface{}
	Status       bool
}
