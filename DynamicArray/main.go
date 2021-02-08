package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the 'dynamicArray' function below.
 *
 * The function is expected to return an INTEGER_ARRAY.
 * The function accepts following parameters:
 *  1. INTEGER n
 *  2. 2D_INTEGER_ARRAY queries
 */

func handleQuery1(idx int32, query []int32, arr [][]int32) [][]int32 {
	arr[idx] = append(arr[idx], query[2])
	return arr
}

func handleQuery2(idx int32, query []int32, arr [][]int32) int32 {
	answerIdx := query[2] % int32(len(arr[idx]))
	return arr[idx][answerIdx]
}

func handleQuery(query, answers []int32, arr [][]int32) ([][]int32, []int32) {
	var answer int32
	if len(answers) == 0 {
		answer = 0
	} else {
		answer = answers[len(answers)-1]
	}

	n := int32(len(arr))
	idx := (query[1] ^ answer) % n

	switch query[0] {
	case 1:
		return handleQuery1(idx, query, arr), answers
	case 2:
		return arr, append(answers, handleQuery2(idx, query, arr))
	default:
		return arr, answers
	}
}

func dynamicArray(n int32, queries [][]int32) []int32 {
	// Write your code here
	var answers []int32
	arr := make([][]int32, n)
	for _, q := range queries {
		arr, answers = handleQuery(q, answers, arr)
	}
	return answers
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	firstMultipleInput := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	nTemp, err := strconv.ParseInt(firstMultipleInput[0], 10, 64)
	checkError(err)
	n := int32(nTemp)

	qTemp, err := strconv.ParseInt(firstMultipleInput[1], 10, 64)
	checkError(err)
	q := int32(qTemp)

	var queries [][]int32
	for i := 0; i < int(q); i++ {
		queriesRowTemp := strings.Split(strings.TrimRight(readLine(reader), " \t\r\n"), " ")

		var queriesRow []int32
		for _, queriesRowItem := range queriesRowTemp {
			queriesItemTemp, err := strconv.ParseInt(queriesRowItem, 10, 64)
			checkError(err)
			queriesItem := int32(queriesItemTemp)
			queriesRow = append(queriesRow, queriesItem)
		}

		if len(queriesRow) != 3 {
			panic("Bad input")
		}

		queries = append(queries, queriesRow)
	}

	result := dynamicArray(n, queries)

	for i, resultItem := range result {
		fmt.Fprintf(writer, "%d", resultItem)

		if i != len(result)-1 {
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
