package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)


var twelveNumbers = []string{
	"one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "ten", "eleven", "twelve",
}

var otherNumbers = []string{
	"thirteen", "fourteen", "quarter", "sixteen", "seventeen", "eighteen", "nineteen", "twenty", "twenty one",
	"twenty two", "twenty three", "twenty four", "twenty five", "twenty six", "twenty seven", "twenty eight",
	"twenty nine",
}

func timeInWords(h int32, m int32) string {
	linkWord := "past "
	if m > 30 {
		linkWord = "to "
		m = 60 - m
		h += 1
		h %= 12
	}
	hStr := twelveNumbers[h-1]
	prefix := ""
	suffix := ""
	if m == 0 {
		suffix = "o' clock"
		linkWord = ""
	} else if m < 13 {
		prefix = twelveNumbers[m-1] + " "
	} else if m == 30 {
		prefix = "half "
	} else {
		prefix = otherNumbers[m-13] + " "
	}
	if m > 1 && m%15 != 0 {
		prefix += "minutes "
	} else if m%15 != 0 {
		prefix += "minute "
	}
	return prefix + linkWord + hStr + " " + suffix
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024 * 1024)

	hTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	h := int32(hTemp)

	mTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	m := int32(mTemp)

	result := timeInWords(h, m)

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
