package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func minOf(a, b int32) int32 {
	if a < b {
		return a
	}
	return b
}

// Complete the acmTeam function below.
func acmTeam(topic []string) []int32 {
	max := int32(-1)
	count := int32(0)

	for i:=0; i<len(topic) - 1; i++ {
		personA := topic[i]
		l := len(personA)
		for j:=i+1; j<len(topic); j++ {
			personB := topic[j]
			knowledge := int32(0)
			for k:=0; k<l; k++ {
				knowledge += minOf(1, int32(personA[k] + personB[k] - 2 * '0'))
			}
			if knowledge > max {
				max = knowledge
				count = 1
			} else if knowledge == max {
				count ++
			}
		}
	}
	return []int32{max, count}
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024 * 1024)

	nm := strings.Split(readLine(reader), " ")

	nTemp, err := strconv.ParseInt(nm[0], 10, 64)
	checkError(err)
	n := int32(nTemp)

	_, err = strconv.ParseInt(nm[1], 10, 64)
	checkError(err)
	//_ := int32(mTemp)

	var topic []string

	for i := 0; i < int(n); i++ {
		topicItem := readLine(reader)
		topic = append(topic, topicItem)
	}

	result := acmTeam(topic)

	for i, resultItem := range result {
		fmt.Fprintf(writer, "%d", resultItem)

		if i != len(result) - 1 {
			fmt.Fprintf(writer, "\n")
		}
	}

	fmt.Fprintf(writer, "\n")

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
