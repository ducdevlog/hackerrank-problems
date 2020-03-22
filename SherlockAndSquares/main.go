package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Complete the squares function below.
func squares(a int32, b int32) int32 {
	count := int32(0)
	for i:= 1; ; i++ {
		root := int32(i*i)
		if root > b {
			break
		}
		if root >= a {
			count ++
		}
	}
	return count
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024 * 1024)

	qTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	q := int32(qTemp)

	for qItr := 0; qItr < int(q); qItr++ {
		ab := strings.Split(readLine(reader), " ")

		aTemp, err := strconv.ParseInt(ab[0], 10, 64)
		checkError(err)
		a := int32(aTemp)

		bTemp, err := strconv.ParseInt(ab[1], 10, 64)
		checkError(err)
		b := int32(bTemp)

		result := squares(a, b)

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
