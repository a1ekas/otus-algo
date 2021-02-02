package tests

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"sync"

	"github.com/a1ekas/otus-algo/internal/reader"
	"github.com/a1ekas/otus-algo/internal/tests/lkytkts"
)

// LuckyTicketsTest ...
type LuckyTicketsTest struct {
	testName     string
	testDataPath string
	inputData    map[int]int
	expectedData map[int]int64
	tcasesLen    int
}

// NewLuckyTicketsTest коструктор LuckyTicketsTest
func NewLuckyTicketsTest() *LuckyTicketsTest {
	t := LuckyTicketsTest{
		testName:     "LuckyTicketsTest",
		testDataPath: "/data/1.Tickets",
		inputData:    make(map[int]int),
		expectedData: make(map[int]int64),
		tcasesLen:    0,
	}
	t.LoadTestData()
	return &t
}

// LoadTestData - имплементация интерфейса Test
func (t *LuckyTicketsTest) LoadTestData() {
	in, out := reader.LoadTestData(t.testDataPath)
	inLen, outLen := len(in), len(out)
	if inLen != outLen || inLen == 0 || outLen == 0 {
		log.Fatalln("Something went wrong. Please check data directory. It should contains data files.")
	}

	for i, tdata := range in {
		idata, _ := strconv.Atoi(strings.TrimSpace(tdata))
		t.inputData[i] = idata
	}

	for j, expData := range out {
		edata, _ := strconv.ParseInt(strings.TrimSpace(expData), 10, 64)
		t.expectedData[j] = edata
	}
}

// Check - имплементация интерфейса Test
func (t *LuckyTicketsTest) Check() {
	tcaseCh := make(chan Case, len(t.inputData))
	log.Println(t.testName + " started asynchronously.")

	var wg sync.WaitGroup
	// Асинхронно запускаем каждый тест в своей горутине
	for i := 0; i < len(t.inputData); i++ {
		wg.Add(1)
		go func(wg *sync.WaitGroup, tcase Case, tcaseCh chan Case) {
			a := lkytkts.NewLuckyTicketsAlgo(tcase.InData.(int))
			res := a.Count()
			expD := tcase.ExpectedData.(int64)
			tcase.Status = res == expD
			tcaseCh <- tcase
			wg.Done()
		}(
			&wg,
			Case{ID: i, InData: t.inputData[i], ExpectedData: t.expectedData[i]},
			tcaseCh,
		)
	}
	wg.Wait()
	close(tcaseCh)
	// Проверяем результат тестов и выводим результат
	for tcase := range tcaseCh {
		if tcase.Status {
			log.Println(t.testName + ": case " + fmt.Sprint(tcase.ID) + " passed.")
		} else {
			log.Println(t.testName + ": case " + fmt.Sprint(tcase.ID) + " failed.")
		}
	}

	log.Println(t.testName + " completed.")
}
