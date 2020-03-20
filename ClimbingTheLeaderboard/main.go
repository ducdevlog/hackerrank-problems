package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func buildLeaderBoard(scores []int32) ([]int32, int)  {
	leaderBoardIndex := 1
	leaderBoard := make([]int32, len(scores) + 1)
	leaderBoard[leaderBoardIndex] = scores[0]
	for i:=1; i<len(scores); i++ {
		if scores[i] == scores[i-1] {
			continue
		}
		leaderBoardIndex ++
		leaderBoard[leaderBoardIndex] = scores[i]
	}
	return leaderBoard, leaderBoardIndex
}

func climbingLeaderboard(scores []int32, alice []int32) []int32 {
	leaderBoard, leaderBoardSize := buildLeaderBoard(scores)

	aliceRanks := make([]int32, len(alice))
	j:=leaderBoardSize
	for i:=0; i<len(alice); i++ {
		for ; j>0; j-- {
			if alice[i] < leaderBoard[j] {
				aliceRanks[i] = int32(j) + 1
				break
			}
		}
		if aliceRanks[i] == 0 {
			aliceRanks[i] = int32(1)
		}
	}
	return aliceRanks
}

func main() {
	//file, err := os.Open("/Users/ducnt/workspace/learning/go/hackerank-problems/ClimbingTheLeaderboard/input06.txt")
	//if err != nil {
	//	fmt.Println(err)
	//}
	//defer file.Close()
	//reader := bufio.NewReaderSize(file, 2048 * 2048)
	reader := bufio.NewReaderSize(os.Stdin, 2048 * 2048)

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
