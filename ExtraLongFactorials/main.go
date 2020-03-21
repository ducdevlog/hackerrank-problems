package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func paddingZero(s string, amount int) string {
	for ; amount > 0; amount-- {
		s = "0" + s
	}
	return s
}

func sum(a, b string) string {
	if len(a) < len(b) {
		diff := len(b) - len(a)
		a = paddingZero(a, diff)
	}
	if len(b) < len(a) {
		diff := len(a) - len(b)
		b = paddingZero(b, diff)
	}
	result := ""
	temp := uint8(0)
	for i:=len(a)-1; i>=0; i-- {
		cumulus := temp + a[i] + b[i] - 2 * '0'
		result = fmt.Sprintf("%v", cumulus % 10) + result
		temp = cumulus / 10
	}
	if temp > 0 {
		result = fmt.Sprintf("%v", temp) + result
	}
	return result
}

func multiply(a, b string) string {
	globalResult := "0"
	for i:=len(b)-1; i>=0; i-- {
		db := b[i] - '0'
		tempA := uint8(0)
		localResult := ""
		for j:=0; j<len(b)-i-1; j++ {
			localResult += "0"
		}
		for j:=len(a)-1; j>=0; j-- {
			da := a[j] - '0'
			p := da * db + tempA
			localResult = fmt.Sprintf("%v", p % 10) + localResult
			tempA = p / 10
		}
		if tempA > 0 {
			localResult = fmt.Sprintf("%v", tempA) + localResult
		}
		globalResult = sum(globalResult, localResult)
	}
	return globalResult
}

func factorial(n int32) string {
	result := "1"
	for i:=2; i<=int(n); i++ {
		result = multiply(result, fmt.Sprintf("%v", i))
	}
	return result
}

// Complete the extraLongFactorials function below.
func extraLongFactorials(n int32) {
	f := factorial(n)
	fmt.Println(f)
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

	nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	n := int32(nTemp)

	extraLongFactorials(n)
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
