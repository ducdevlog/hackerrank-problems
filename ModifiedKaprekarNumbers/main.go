package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func isKaprekarNumber(n int64) bool {
	nl := len(strconv.Itoa(int(n)))
	sqr := n * n
	divisor := int64(1)
	for i:=0; i<nl; i++ {
		divisor *= 10
	}
	r := sqr % divisor
	l := sqr / divisor
	return n == r + l
}

// Complete the kaprekarNumbers function below.
func kaprekarNumbers(p int32, q int32) {
	var results []string
	for n := p; n <= q; n++ {
		if isKaprekarNumber(int64(n)) {
			results = append(results, strconv.Itoa(int(n)))
		}
	}
	if len(results) == 0 {
		fmt.Println("INVALID RANGE")
		return
	}
	result := strings.Join(results, " ")
	fmt.Println(result)
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

	pTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	p := int32(pTemp)

	qTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	q := int32(qTemp)

	kaprekarNumbers(p, q)
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
