package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Complete the organizingContainers function below.
func organizingContainers(container [][]int32) string {
	types := make([]int64, 100)
	containers := make([]int64, 100)
	n := len(container)
	for i := 0; i < n; i++ {
		var ballsPerType int64 = 0
		var ballsPerContainer int64 = 0
		for j := 0; j < n; j++ {
			ballsPerType += int64(container[j][i])
			ballsPerContainer += int64(container[i][j])
		}
		types[i] = ballsPerType
		containers[i] = ballsPerContainer
	}
	for i := 0; i < n; i++ {
		check := false
		for j := 0; j < n; j++ {
			if types[i] == containers[j] {
				check = true
				break
			}
		}
		if !check {
			return "Impossible"
		}
	}
	return "Possible"
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	qTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	q := int32(qTemp)

	for qItr := 0; qItr < int(q); qItr++ {
		nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
		checkError(err)
		n := int32(nTemp)

		var container [][]int32
		for i := 0; i < int(n); i++ {
			containerRowTemp := strings.Split(readLine(reader), " ")

			var containerRow []int32
			for _, containerRowItem := range containerRowTemp {
				containerItemTemp, err := strconv.ParseInt(containerRowItem, 10, 64)
				checkError(err)
				containerItem := int32(containerItemTemp)
				containerRow = append(containerRow, containerItem)
			}

			if len(containerRow) != int(int(n)) {
				panic("Bad input")
			}

			container = append(container, containerRow)
		}

		result := organizingContainers(container)

		fmt.Fprintf(writer, "%s\n", result)
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
