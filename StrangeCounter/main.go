package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

// Complete the strangeCounter function below.
func strangeCounter(t int64) int64 {
	tmp := float64(t + 2) / float64(3)
	n := math.Log2(tmp)
	floorN := math.Floor(n)
	return 3 * int64(math.Pow(float64(2), floorN + 1)) - t - 2
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024 * 1024)

	t, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)

	result := strangeCounter(t)

	fmt.Fprintf(writer, "%d\n", result)

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
