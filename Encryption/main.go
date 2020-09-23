package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strings"
)

func findRange(l int) (int, int) {
	sqrt := math.Sqrt(float64(l))
	floor := int(math.Floor(sqrt))
	ceil := int(math.Ceil(sqrt))
	if floor * ceil < l {
		floor ++
	}
	return floor, ceil
}

// Complete the encryption function below.
func encryption(s string) string {
	l := len(s)
	_, columns := findRange(l)
	result := ""
	for i:=0; i<columns; i++ {
		j := i
		for ; j<l; j+=columns {
			result += string(s[j])
		}
		result += " "
	}
	return result
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024 * 1024)

	s := readLine(reader)

	result := encryption(s)

	fmt.Fprintf(writer, "%s\n", result)

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
