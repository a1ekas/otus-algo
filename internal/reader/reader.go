package reader

import (
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

const (
	// InputFileDataMask - маска для поиска файлов с тестовыми данными
	InputFileDataMask = `test.\d+.in`
	// OutputFileDataMask - маска для поиска файлов с ожидаемым результатом
	OutputFileDataMask = `test.\d+.out`
)

// InputDataType ...
type InputDataType map[int]string

// OutputDataType ...
type OutputDataType map[int]string

// LoadTestData ...
func LoadTestData(path string) (InputDataType, OutputDataType) {
	projectRoot, _ := os.Getwd()
	absPath := projectRoot + path
	f, err := os.Open(absPath)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	filesInfo, err := f.Readdir(-1)
	if err != nil {
		log.Fatal(err)
	}

	in := make(InputDataType)
	out := make(OutputDataType)

	for _, file := range filesInfo {
		if matched, _ := regexp.MatchString(InputFileDataMask, file.Name()); matched {
			tkey := extractTestKey(file.Name())
			in[tkey] = readFile(absPath + "/" + file.Name())
		}

		if matched, _ := regexp.MatchString(OutputFileDataMask, file.Name()); matched {
			tkey := extractTestKey(file.Name())
			out[tkey] = readFile(absPath + "/" + file.Name())
		}

	}

	return in, out
}

func extractTestKey(fileName string) int {
	explodeRes := strings.Split(fileName, ".")
	tkey, _ := strconv.Atoi(explodeRes[1])
	return tkey
}

func readFile(path string) string {
	d, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	return string(d)
}
