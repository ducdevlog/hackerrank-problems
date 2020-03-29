package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func isInCol(rq, cq int32, r, c int32) int32 {
	if c != cq {
		return 0
	} else if rq < r {
		return 1
	}
	return -1
}

func isInRow(rq, cq int32, r, c int32) int32 {
	if r != rq {
		return 0
	} else if cq < c {
		return 1
	}
	return -1
}

func isInDiagonal(rq, cq int32, r, c int32) int32 {
	// Diagonal equation
	if r == c + rq - cq {
		if r > rq {
			return 1
		}
		return -1
	}
	return 0
}

func isInSubDiagonal(rq, cq int32, r, c int32) int32 {
	// Sub-Diagonal equation
	if r == rq + cq - c {
		if r > rq {
			return 1
		}
		return -1
	}
	return 0
}

func minOf(a, b int32) int32 {
	if a > b {
		return b
	}
	return a
}

func getCountByDirectionWithoutObstacles(n int32, rq int32, cq int32) []int32 {
	return []int32{
		n - rq,
		minOf(n-rq, n-cq),
		n-cq, minOf(rq-1, n-cq),
		rq-1,
		minOf(rq-1, cq-1),
		cq-1,
		minOf(n-rq, cq-1)}
}

// Complete the queensAttack function below.
func queensAttack(n int32, k int32, r_q int32, c_q int32, obstacles [][]int32) int32 {
	t, tr, r, dr, d, dl, l, tl := 0, 1, 2, 3, 4, 5, 6, 7
	minValues := getCountByDirectionWithoutObstacles(n, r_q, c_q)
	for _, obstacle := range obstacles {
		or, oc  := obstacle[0], obstacle[1]
		inRow := isInRow(r_q, c_q, or, oc)
		if inRow == -1 {
			minValues[l] = minOf(minValues[l], c_q - oc - 1)
			continue
		} else if inRow == 1 {
			minValues[r] = minOf(minValues[r], oc - c_q - 1)
			continue
		}
		inColumn := isInCol(r_q, c_q, or, oc)
		if inColumn == -1 {
			minValues[d] = minOf(minValues[d], r_q - or - 1)
			continue
		} else if inColumn == 1 {
			minValues[t] = minOf(minValues[t], or - r_q - 1)
			continue
		}
		inDiagonal := isInDiagonal(r_q, c_q, or, oc)
		if inDiagonal == -1 {
			minValues[dl] = minOf(minValues[dl], r_q - or - 1)
			continue
		} else if inDiagonal == 1 {
			minValues[tr] = minOf(minValues[tr], or - r_q - 1)
			continue
		}
		inSubDiagonal := isInSubDiagonal(r_q, c_q, or, oc)
		if inSubDiagonal == -1 {
			minValues[dr] = minOf(minValues[dr], r_q - or - 1)
			continue
		} else if inSubDiagonal == 1 {
			minValues[tl] = minOf(minValues[tl], or - r_q - 1)
			continue
		}
	}
	sum := int32(0)
	for _, v := range minValues {
		sum += v
	}
	return sum
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 2048 * 2048)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024 * 1024)

	nk := strings.Split(readLine(reader), " ")

	nTemp, err := strconv.ParseInt(nk[0], 10, 64)
	checkError(err)
	n := int32(nTemp)

	kTemp, err := strconv.ParseInt(nk[1], 10, 64)
	checkError(err)
	k := int32(kTemp)

	r_qC_q := strings.Split(readLine(reader), " ")

	r_qTemp, err := strconv.ParseInt(r_qC_q[0], 10, 64)
	checkError(err)
	r_q := int32(r_qTemp)

	c_qTemp, err := strconv.ParseInt(r_qC_q[1], 10, 64)
	checkError(err)
	c_q := int32(c_qTemp)

	var obstacles [][]int32
	for i := 0; i < int(k); i++ {
		obstaclesRowTemp := strings.Split(readLine(reader), " ")

		var obstaclesRow []int32
		for _, obstaclesRowItem := range obstaclesRowTemp {
			obstaclesItemTemp, err := strconv.ParseInt(obstaclesRowItem, 10, 64)
			checkError(err)
			obstaclesItem := int32(obstaclesItemTemp)
			obstaclesRow = append(obstaclesRow, obstaclesItem)
		}

		if len(obstaclesRow) != int(2) {
			panic("Bad input")
		}

		obstacles = append(obstacles, obstaclesRow)
	}

	result := queensAttack(n, k, r_q, c_q, obstacles)

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
