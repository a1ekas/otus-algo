package cases

import (
	"errors"
	"strconv"

	"github.com/a1ekas/otus-algo/internal/reader"
)

// StringLenTestCase - структура для проверки длин строк и системы тестирования
type StringLenTestCase struct {
	testDataPath string
	inputData    map[int]string
	expectedData map[int]int
}

// NewStringLenTestCase коструктор StringLenTestCase
func NewStringLenTestCase() *StringLenTestCase {
	c := StringLenTestCase{
		testDataPath: reader.ProjectRoot + "/data/0.String",
		inputData:    make(map[int]string),
		expectedData: make(map[int]int),
	}
	c.LoadTestData()
	return &c
}

// LoadTestData - имплементация интерфейса Case
func (c *StringLenTestCase) LoadTestData() error {
	in, out := reader.LoadTestData(c.testDataPath)

	if len(in) != len(out) {
		return errors.New("Something went wrong. The number of files does not match")
	}

	for i, tdata := range in {
		c.inputData[i] = tdata
	}

	for j, expData := range out {
		edata, _ := strconv.Atoi(expData)
		c.expectedData[j] = edata
	}

	return nil
}

// Check - имплементация интерфейса Case
func (c *StringLenTestCase) Check() error {
	return nil
}
