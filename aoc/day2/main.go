package main

import (
	"bytes"
	"fmt"
	"os"
)

type GameElement string
type GameResult string

const (
	Rock     GameElement = "Rock"
	Paper    GameElement = "Paper"
	Scissors GameElement = "Scissors"
)

const (
	W GameResult = "Win"
	D GameResult = "Draw"
	L GameResult = "Lost"
)

type GameElementInfo struct {
	Player1Move string
	Player2Move string
	Score       int
}

var gameMapping = map[GameElement]GameElement{
	Rock:     Scissors,
	Paper:    Rock,
	Scissors: Paper,
}

var gameScores = map[GameResult]int{
	W: 6,
	D: 3,
	L: 0,
}

var preCalculatedResults map[string]int

var elementMapping = map[GameElement]GameElementInfo{
	Rock:     {Player1Move: "A", Player2Move: "X", Score: 1},
	Paper:    {Player1Move: "B", Player2Move: "Y", Score: 2},
	Scissors: {Player1Move: "C", Player2Move: "Z", Score: 3},
}

func getRoundResult(player1Move, player2Move GameElement) GameResult {
	if player1Move == player2Move {
		return D
	}

	if gameMapping[player2Move] == player1Move {
		return W
	}
	return L
}

func init() {
	preCalculatedResults = make(map[string]int)
	for p1move, p1info := range elementMapping {
		for p2move, p2info := range elementMapping {
			result := getRoundResult(p1move, p2move)
			key := p1info.Player1Move + p2info.Player2Move
			preCalculatedResults[key] = gameScores[result] + p2info.Score
		}
	}
}

func getPlayer2MoveResultWithScore(player1move, player2move string) (int, bool) {
	key := player1move + player2move
	score, found := preCalculatedResults[key]
	return score, found
}

func getFinalResult(file string) int {
	total := 0
	fileContent, err := os.ReadFile(file)

	if err != nil {
		fmt.Println("Error reading file: ", file, err)
		return 0
	}

	lines := bytes.Split(fileContent, []byte("\n"))

	for i, line := range lines {
		input := bytes.Split(line, []byte(" "))
		player1Move := string(input[0])
		player2Move := string(input[1])

		score, found := getPlayer2MoveResultWithScore(player1Move, player2Move)
		if found {
			total += score
		} else {
			fmt.Printf("Unable to find score at line %d with moves: %s %s\n", i+1, player1Move, player2Move)
		}
	}

	return total
}

func main() {
	file := "day2-input.txt"
	result := getFinalResult(file)
	fmt.Printf("The final score is %d\n", result)
}
