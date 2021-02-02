package tests

import (
	"log"
	"strconv"

	"github.com/a1ekas/otus-algo/internal/reader"
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
func (c *LuckyTicketsTest) LoadTestData() {
	in, out := reader.LoadTestData(c.testDataPath)
	inLen, outLen := len(in), len(out)
	if inLen != outLen || inLen == 0 || outLen == 0 {
		log.Fatalln("Something went wrong. Please check data directory. It should contains data files.")
	}

	for i, tdata := range in {
		idata, _ := strconv.Atoi(tdata)
		c.inputData[i] = idata
	}

	for j, expData := range out {
		edata, _ := strconv.ParseInt(expData, 10, 64)
		c.expectedData[j] = edata
	}
}

// Check - имплементация интерфейса Test
func (c *LuckyTicketsTest) Check() {

}
