package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func climbingLeaderboard(scores []int32, alice []int32) []int32 {
	var ranks []int32
	var scoreRanks []int32
	leaderBoardRank := int32(1)
	for i:=0; i<len(scores); i++ {
		if i > 0 && scores[i] < scores[i-1] {
			leaderBoardRank ++
		}
		scoreRanks = append(scoreRanks, leaderBoardRank)
	}

	for i:=0; i<len(alice); i++ {
		for j:=len(scores) - 1; j>=0; j-- {
			if scores[j] > alice[i] {
				ranks = append(ranks, scoreRanks[j]+1)
				break
			} else if scores[j] == alice[i] {
				ranks = append(ranks, scoreRanks[j])
				break
			}
		}
		if i == len(ranks) {
			ranks = append(ranks, 1)
		}
	}
	return ranks
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024 * 1024)

	scoresCount, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)

	scoresTemp := strings.Split(readLine(reader), " ")

	var scores []int32

	for i := 0; i < int(scoresCount); i++ {
		scoresItemTemp, err := strconv.ParseInt(scoresTemp[i], 10, 64)
		checkError(err)
		scoresItem := int32(scoresItemTemp)
		scores = append(scores, scoresItem)
	}

	aliceCount, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)

	aliceTemp := strings.Split(readLine(reader), " ")

	var alice []int32

	for i := 0; i < int(aliceCount); i++ {
		aliceItemTemp, err := strconv.ParseInt(aliceTemp[i], 10, 64)
		checkError(err)
		aliceItem := int32(aliceItemTemp)
		alice = append(alice, aliceItem)
	}

	result := climbingLeaderboard(scores, alice)

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
