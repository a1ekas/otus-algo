package tests

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"sync"

	"github.com/a1ekas/otus-algo/internal/reader"
)

// StringLenTest - структура для проверки длин строк и системы тестирования
type StringLenTest struct {
	testName     string
	testDataPath string
	inputData    map[int]string
	expectedData map[int]int
	tcasesLen    int
}

// NewStringLenTest коструктор StringLenTest
func NewStringLenTest() *StringLenTest {
	t := StringLenTest{
		testName:     "StringLenTest",
		testDataPath: "/data/0.String",
		inputData:    make(map[int]string),
		expectedData: make(map[int]int),
		tcasesLen:    0,
	}
	t.LoadTestData()
	return &t
}

// LoadTestData - имплементация интерфейса Test
func (t *StringLenTest) LoadTestData() {
	in, out := reader.LoadTestData(t.testDataPath)
	inLen, outLen := len(in), len(out)
	if inLen != outLen || inLen == 0 || outLen == 0 {
		log.Fatalln("Something went wrong. Please check data directory. It should contains data files.")
	}

	for i, tdata := range in {
		t.inputData[i] = tdata
	}

	for j, expData := range out {
		edata, _ := strconv.Atoi(expData)
		t.expectedData[j] = edata
	}
}

// Check - имплементация интерфейса Test
func (t *StringLenTest) Check() {
	log.Println(t.testName + " started asynchronously.")

	var wg sync.WaitGroup
	// Асинхронно запускаем каждый тест в своей горутине
	for i := 0; i < len(t.inputData); i++ {
		wg.Add(1)
		go func(wg *sync.WaitGroup, tcase Case) {
			if len(strings.TrimSpace(tcase.InData.(string))) == tcase.ExpectedData.(int) {
				log.Println(t.testName + ": case " + fmt.Sprint(tcase.ID) + " passed.")
			} else {
				log.Println(t.testName + ": case " + fmt.Sprint(tcase.ID) + " failed.")
			}
			wg.Done()
		}(
			&wg,
			Case{ID: i, InData: t.inputData[i], ExpectedData: t.expectedData[i]},
		)
	}
	wg.Wait()
	log.Println(t.testName + " completed.")
}
