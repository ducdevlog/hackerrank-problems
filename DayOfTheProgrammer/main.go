package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func getCalendarType(year int32) string  {
	if year < 1918 {
		return "Julian"
	} else if year >  1918 {
		return "Gregorian"
	}
	return "Transition"
}

func isLeapYear(year int32, calendarType string) bool {
	if "Julian" == calendarType {
		return year % 4 == 0
	} else if "Gregorian" == calendarType {
		return (year %  4 == 0 && year % 100 != 0) || year % 400 == 0
	}
	return false
}

// Complete the dayOfProgrammer function below.
func dayOfProgrammer(year int32) string {
	calendarType := getCalendarType(year)
	if "Transition" == calendarType {
		return "26.09.1918"
	}
	if isLeapYear(year, calendarType)  {
		return "12.09." + fmt.Sprintf("%v", year)
	}
	return "13.09." + fmt.Sprintf("%v", year)
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16 * 1024 * 1024)

	yearTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	year := int32(yearTemp)

	result := dayOfProgrammer(year)

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
