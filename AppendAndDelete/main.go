package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Complete the appendAndDelete function below.
func appendAndDelete(s string, t string, k int32) string {
	i:=0
	for ; i<len(s) && i<len(t); i++ {
		if s[i] != t[i] {
			break
		}
	}
	del := 0
	add := 0
	if i<len(s) {
		del = len(s) - i
	}
	if i<len(t) {
		add = len(t) - i
	}
	atLeast := int32(del + add)
	ok := atLeast == k
	if k > atLeast {
		sub := k - atLeast
		ok = sub % 2 == 0 || sub > 2 * int32(i)
	}
	if ok {
		return "Yes"
	}
	return "No"
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024 * 1024)

	s := readLine(reader)

	t := readLine(reader)

	kTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	k := int32(kTemp)

	result := appendAndDelete(s, t, k)

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
