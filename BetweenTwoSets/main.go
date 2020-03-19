package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func isCommonMultipleOf(consider int32, a []int32) bool {
	for _, i := range a {
		if consider % i != 0 {
			return false
		}
	}
	return true
}

func isCommonDivisorOf(consider int32, a []int32) bool {
	for _, i := range a {
		if i % consider != 0 {
			return false
		}
	}
	return true
}

func findMaxValue(a []int32) int32 {
	max := int32(-1 << 31)
	for _, i := range a {
		if max < i {
			max = i
		}
	}
	return max
}

func findMinValue(a []int32) int32 {
	min := int32(1 << 31 - 1)
	for _, i := range a {
		if min > i {
			min = i
		}
	}
	return min
}

/*
 * Complete the 'getTotalX' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER_ARRAY a
 *  2. INTEGER_ARRAY b
 */

func getTotalX(a []int32, b []int32) int32 {
	// Write your code here
	maxA := findMaxValue(a)
	minB := findMinValue(b)
	count := 0
	for i:=maxA; i<=minB; i+=maxA {
		if isCommonMultipleOf(i, a) && isCommonDivisorOf(i, b) {
			count ++
		}
	}
	return int32(count)
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16 * 1024 * 1024)

	firstMultipleInput := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	nTemp, err := strconv.ParseInt(firstMultipleInput[0], 10, 64)
	checkError(err)
	n := int32(nTemp)

	mTemp, err := strconv.ParseInt(firstMultipleInput[1], 10, 64)
	checkError(err)
	m := int32(mTemp)

	arrTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var arr []int32

	for i := 0; i < int(n); i++ {
		arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
		checkError(err)
		arrItem := int32(arrItemTemp)
		arr = append(arr, arrItem)
	}

	brrTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var brr []int32

	for i := 0; i < int(m); i++ {
		brrItemTemp, err := strconv.ParseInt(brrTemp[i], 10, 64)
		checkError(err)
		brrItem := int32(brrItemTemp)
		brr = append(brr, brrItem)
	}

	total := getTotalX(arr, brr)

	fmt.Fprintf(writer, "%d\n", total)

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
