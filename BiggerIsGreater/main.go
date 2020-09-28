package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func bestSuiteCandidate(currentCandidate, targetCandidate int, w string) int  {
	for i:=currentCandidate; i>targetCandidate; i-- {
		if w[i] > w[targetCandidate] {
			return i
		}
	}
	return targetCandidate + 1
}

func reorderSuffix(from int, s string, insertRune rune) string {
	subStr := s[from:]
	l := len(subStr)
	suffix := ""
	for i:=l-1; i>=0; i-- {
		if subStr[i] > uint8(insertRune) {
			suffix += string(insertRune)
			insertRune = 0
		}
		suffix += string(subStr[i])
	}
	if insertRune > 0 {
		suffix += string(insertRune)
	}
	return suffix
}

func swapChars(first, second int, w string) string {
	r := []rune(w)
	tmp := r[second]
	r[second] = r[first]
	if first == len(r) - 1 {
		r = r[:first]
	} else {
		r = append(r[:first], r[first+1:]...)
	}

	afterReplace := string(r)
	prefix := afterReplace[:second+1]

	suffix := reorderSuffix(second+1, afterReplace, tmp)
	return prefix + suffix
}

// Complete the biggerIsGreater function below.
func biggerIsGreater(w string) string {
	l := len(w)
	if l < 2 {
		return "no answer"
	}
	firstCandidate := l-1
	secondCandidate := -1
	for i:=l-2; i>=0; i-- {
		if w[i] < w[i+1] {
			secondCandidate = i
			firstCandidate = bestSuiteCandidate(firstCandidate, secondCandidate, w)
			break
		}
	}
	if secondCandidate == -1 {
		return "no answer"
	}
	return swapChars(firstCandidate, secondCandidate, w)
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024 * 1024)

	TTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	T := int32(TTemp)

	for TItr := 0; TItr < int(T); TItr++ {
		w := readLine(reader)

		result := biggerIsGreater(w)

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
