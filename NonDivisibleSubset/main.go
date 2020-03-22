package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func modulusFrequency(k int32, s []int32) []int32 {
	modulus := make([]int32, k)
	for _, i:= range s {
		modulus[i%k]  ++
	}
	return modulus
}

/*
 * Complete the 'nonDivisibleSubset' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER k
 *  2. INTEGER_ARRAY s
 */

func nonDivisibleSubset(k int32, s []int32) int32 {
	// Write your code here
	modulusFreq := modulusFrequency(k, s)
	if k % 2 == 0 && modulusFreq[k/2] > 1 {
		modulusFreq[k/2] = 1
	}
	if modulusFreq[0] > 1 {
		modulusFreq[0] = 1
	}
	subsetCount := modulusFreq[0]
	for i:=int32(1); i<=k/2; i++ {
		if modulusFreq[i] > modulusFreq[k-i] {
			subsetCount += modulusFreq[i]
		} else {
			subsetCount += modulusFreq[k-i]
		}
	}
	return subsetCount
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

	kTemp, err := strconv.ParseInt(firstMultipleInput[1], 10, 64)
	checkError(err)
	k := int32(kTemp)

	sTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var s []int32

	for i := 0; i < int(n); i++ {
		sItemTemp, err := strconv.ParseInt(sTemp[i], 10, 64)
		checkError(err)
		sItem := int32(sItemTemp)
		s = append(s, sItem)
	}

	result := nonDivisibleSubset(k, s)

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
