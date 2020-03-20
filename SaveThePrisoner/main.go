
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Complete the saveThePrisoner function below.
func saveThePrisoner(n int32, m int32, s int32) int32 {
	if n == 1 {
		return 1
	}
	mod := m % n - 1
	result := mod + s
	if result > n {
		result -= n
	} else if result == 0 {
		return n
	}
	return result
}


func main() {
	//file, err := os.Open("/Users/ducnt/workspace/learning/go/hackerank-problems/SaveThePrisoner/test.txt")
	//checkError(err)
	reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	//stdout, err := os.Create("/Users/ducnt/workspace/learning/go/hackerank-problems/SaveThePrisoner/log.txt")
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024 * 1024)

	tTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	t := int32(tTemp)

	for tItr := 0; tItr < int(t); tItr++ {
		nms := strings.Split(readLine(reader), " ")

		nTemp, err := strconv.ParseInt(nms[0], 10, 64)
		checkError(err)
		n := int32(nTemp)

		mTemp, err := strconv.ParseInt(nms[1], 10, 64)
		checkError(err)
		m := int32(mTemp)

		sTemp, err := strconv.ParseInt(nms[2], 10, 64)
		checkError(err)
		s := int32(sTemp)

		result := saveThePrisoner(n, m, s)

		fmt.Fprintf(writer, "%d\n", result)
	}

	writer.Flush()
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
