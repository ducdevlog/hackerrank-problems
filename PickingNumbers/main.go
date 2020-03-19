package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func sort(array *[]int32)  {
	for i:=0; i<len(*array)-1; i++ {
		for j:=i+1; j<len(*array); j++ {
			if (*array)[i] > (*array)[j] {
				temp := (*array)[i]
				(*array)[i] = (*array)[j]
				(*array)[j] = temp
			}
		}
	}
}

func getMaxValidSubsetLength(array []int32) int32  {
	max := int32(0)
	var current []int32
	j := 0
	for i:=0; i<len(array); i++ {
		for ; j<len(array); j++ {
			subtract := array[j] - array[i]
			if subtract <= 1 && subtract >= -1 {
				current = append(current, array[j])
			} else {
				break
			}
		}
		if len(current) > int(max) {
			max = int32(len(current))
		}
		if current != nil && len(current) > 0 {
			current = current[1:]
		}
	}
	return max
}

/*
 * Complete the 'pickingNumbers' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts INTEGER_ARRAY a as parameter.
 */

func pickingNumbers(a []int32) int32 {
	// Write your code here
	sort(&a)
	return getMaxValidSubsetLength(a)
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16 * 1024 * 1024)

	nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	n := int32(nTemp)

	aTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var a []int32

	for i := 0; i < int(n); i++ {
		aItemTemp, err := strconv.ParseInt(aTemp[i], 10, 64)
		checkError(err)
		aItem := int32(aItemTemp)
		a = append(a, aItem)
	}

	result := pickingNumbers(a)

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
