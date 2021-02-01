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

// ProjectRoot ...
var ProjectRoot string

func init() {
	cwd, _ := os.Getwd()
	ProjectRoot = cwd
}

// LoadTestData ...
func LoadTestData(path string) (InputDataType, OutputDataType) {
	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	filesInfo, err := f.Readdir(-1)
	if err != nil {
		log.Fatal(err)
	}

	var filePath string
	in := make(InputDataType)
	out := make(OutputDataType)

	for _, file := range filesInfo {
		filePath = path + "/" + file.Name()
		if matched, _ := regexp.MatchString(InputFileDataMask, file.Name()); matched {
			tkey := extractTestKey(file.Name())
			in[tkey] = readFile(filePath)
		}

		if matched, _ := regexp.MatchString(OutputFileDataMask, file.Name()); matched {
			tkey := extractTestKey(file.Name())
			out[tkey] = readFile(filePath)
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
