package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Generate permutation of 9
func permutation(i int, values *[]int, checked *[]bool, results *[][][]int32) {
	if i == 9 {
		count := 0
		square := make([][]int32, 3)
		for x:=0; x<3; x++ {
			row := make([]int32, 3)
			for y:=0; y<3; y++ {
				row[y] = int32((*values)[count])
				count ++
			}
			square[x] = row
		}
		if isMagicSquare(square) {
			*results = append(*results, square)
		}
		return
	}
	for j:=1; j<=9; j++ {
		if !(*checked)[j] {
			(*checked)[j] = true
			(*values)[i] = j
			permutation(i+1, values, checked, results)
			(*values)[i] = -1
			(*checked)[j] = false
		}
	}
}

// Check whether is magic square matrix
func isMagicSquare(square [][]int32) bool {
	sum := int32(0)
	for i:=0; i<3; i++ {
		sum += square[0][i]
	}

	for i:= 1; i<3; i++ {
		otherSum := int32(0)
		for _, v := range square[i] {
			otherSum += v
		}
		if sum != otherSum {
			return false
		}
	}

	for j:= 0; j<3; j++ {
		otherSum := int32(0)
		for i:=0; i<3; i++ {
			otherSum += square[i][j]
		}
		if sum != otherSum {
			return false
		}
	}

	otherSum := int32(0)
	for i:=0; i<3; i++ {
		otherSum += square[i][i]
	}
	if sum != otherSum {
		return false
	}

	otherSum = 0
	for i:=0; i<3; i++ {
		otherSum += square[i][2-i]
	}
	if sum != otherSum {
		return false
	}

	return true
}

func calculateDifferent(a [][]int32, b [][]int32) int32 {
	different := int32(0)

	for i:=0; i<len(a); i++ {
		for j:=0; j<len(b); j++ {
			localDiff := a[i][j] - b[i][j]
			if localDiff < 0 {
				localDiff = -localDiff
			}
			different += localDiff
		}
	}

	return different
}

// Complete the formingMagicSquare function below.
func formingMagicSquare(s [][]int32) int32 {
	checked := make([]bool, 10)
	values := make([]int, 9)
	var results [][][]int32
	permutation(0, &values, &checked, &results)

	min := int32(1 << 31 - 1)
	for _, result := range results {
		diff := calculateDifferent(result, s)
		if diff < min {
			min = diff
		}
	}

	return min
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024 * 1024)

	var s [][]int32
	for i := 0; i < 3; i++ {
		sRowTemp := strings.Split(readLine(reader), " ")

		var sRow []int32
		for _, sRowItem := range sRowTemp {
			sItemTemp, err := strconv.ParseInt(sRowItem, 10, 64)
			checkError(err)
			sItem := int32(sItemTemp)
			sRow = append(sRow, sItem)
		}

		if len(sRow) != 3 {
			panic("Bad input")
		}

		s = append(s, sRow)
	}

	result := formingMagicSquare(s)

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
