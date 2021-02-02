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
	c := StringLenTest{
		testName:     "StringLenTest",
		testDataPath: "/data/0.String",
		inputData:    make(map[int]string),
		expectedData: make(map[int]int),
		tcasesLen:    0,
	}
	c.LoadTestData()
	return &c
}

// LoadTestData - имплементация интерфейса Test
func (c *StringLenTest) LoadTestData() {
	in, out := reader.LoadTestData(c.testDataPath)
	inLen, outLen := len(in), len(out)
	if inLen != outLen || inLen == 0 || outLen == 0 {
		log.Fatalln("Something went wrong. Please check data directory. It should contains data files.")
	}

	for i, tdata := range in {
		c.inputData[i] = tdata
	}

	for j, expData := range out {
		edata, _ := strconv.Atoi(expData)
		c.expectedData[j] = edata
	}
}

// Check - имплементация интерфейса Test
func (c *StringLenTest) Check() {
	tcaseCh := make(chan Case, len(c.inputData))
	log.Println(c.testName + " started asynchronously.")
	var wg sync.WaitGroup
	for i := 0; i < len(c.inputData); i++ {
		wg.Add(1)
		go runCase(
			&wg,
			Case{ID: i, InData: c.inputData[i], ExpectedData: c.expectedData[i]},
			tcaseCh,
		)
	}
	wg.Wait()
	close(tcaseCh)

	for tcase := range tcaseCh {
		if tcase.Status {
			log.Println(c.testName + ": case " + fmt.Sprint(tcase.ID) + " passed.")
		} else {
			log.Println(c.testName + ": case " + fmt.Sprint(tcase.ID) + " failed.")
		}
	}

	log.Println(c.testName + " completed.")
}

func runCase(wg *sync.WaitGroup, tcase Case, tcaseCh chan Case) {
	if len(strings.TrimSpace(tcase.InData.(string))) == tcase.ExpectedData.(int) {
		tcase.Status = true
	} else {
		tcase.Status = false
	}
	tcaseCh <- tcase
	wg.Done()
}
