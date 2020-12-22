package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

// Complete the howManyGames function below.
func howManyGames(p int32, d int32, m int32, s int32) int32 {
	// Return the number of games you can buy
	if s < p {
		return 0
	}
	step := (p-m)/d + 1
	// solve equation dx^2 + (d-2p)x + 2s = 0
	x1, x2 := solveEquation(d, d-2*p, 2*(s-p))
	if x1 > x2 {
		x1, x2 = x2, x1
	}
	x1 = math.Floor(x1)
	x2 = math.Ceil(x2)
	candidate1 := int32(x1) + 1
	candidate2 := int32(x2) + 1

	if candidate2 <= step && x2 >= 0 {
		return candidate2
	} else if candidate1 <= step && x1 >= 0 {
		return candidate1
	}

	usedMoney := step * p - step * (step - 1) * d / 2
	leftMoney := s - usedMoney
	leftStep := leftMoney / m
	return step + leftStep
}

func solveEquation(a, b, c int32) (float64, float64) {
	delta := b*b - 4*a*c
	sqrtDelta := math.Sqrt(float64(delta))
	return (-float64(b) - sqrtDelta) / (2 * float64(a)), (-float64(b) + sqrtDelta) / (2 * float64(a))
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024 * 1024)

	pdms := strings.Split(readLine(reader), " ")

	pTemp, err := strconv.ParseInt(pdms[0], 10, 64)
	checkError(err)
	p := int32(pTemp)

	dTemp, err := strconv.ParseInt(pdms[1], 10, 64)
	checkError(err)
	d := int32(dTemp)

	mTemp, err := strconv.ParseInt(pdms[2], 10, 64)
	checkError(err)
	m := int32(mTemp)

	sTemp, err := strconv.ParseInt(pdms[3], 10, 64)
	checkError(err)
	s := int32(sTemp)

	answer := howManyGames(p, d, m, s)

	fmt.Fprintf(writer, "%d\n", answer)

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
