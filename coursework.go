package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

//var currentSpace []int
var boardGrid [][]int

func main() {
	textFileLocation := "./puzzleBoard.txt"
	board, errorGiven := ioutil.ReadFile(textFileLocation)
	if errorGiven != nil {
		log.Fatal(errorGiven)
	}
	boardString := string(board)
	lineByLine := strings.Fields(boardString)
	SizeOfBoard := len(lineByLine)
	boardGrid := make([][]int, SizeOfBoard)
	for i := range boardGrid {
		boardGrid[i] = make([]int, SizeOfBoard)
	}
	//convert strings to int and create the correct board
	for i := range lineByLine {
		a := strings.Split(lineByLine[i], ",")
		for j := range lineByLine {
			number, err := strconv.Atoi(a[j])
			if err != nil {
				log.Fatal(err)
			}
			boardGrid[i][j] = number
		}
	}
	fmt.Println(boardGrid)
	currentSpace := []int{0, 0}
	fmt.Println(currentSpace)
	endingSpace := []int{9, 9} //9,9
	fmt.Println(endingSpace)
	//moveRight(boardGrid[currentSpace[0]][currentSpace[1]], currentSpace)
	//	moveLeft(boardGrid[currentSpace[0]][currentSpace[1]], currentSpace)
	//moveDown(boardGrid[currentSpace[0]][currentSpace[1]], currentSpace)
	//moveUp(boardGrid[currentSpace[0]][currentSpace[1]], currentSpace)
	fmt.Println(currentSpace, boardGrid[currentSpace[0]][currentSpace[1]])

	for !(currentSpace[0] == endingSpace[0] && currentSpace[1] == endingSpace[1]) {
		if currentSpace[1]+boardGrid[currentSpace[0]][currentSpace[1]] < 10 {
			moveRight(boardGrid[currentSpace[0]][currentSpace[1]], currentSpace)
		} else if currentSpace[0]+boardGrid[currentSpace[0]][currentSpace[1]] < 10 {
			moveDown(boardGrid[currentSpace[0]][currentSpace[1]], currentSpace)
		} else if currentSpace[1]-boardGrid[currentSpace[0]][currentSpace[1]] > -1 {
			moveLeft(boardGrid[currentSpace[0]][currentSpace[1]], currentSpace)
		} else if currentSpace[0]-boardGrid[currentSpace[0]][currentSpace[1]] > -1 {
			moveUp(boardGrid[currentSpace[0]][currentSpace[1]], currentSpace)
		} else {
			fmt.Println("test")
		}
		fmt.Println(boardGrid[currentSpace[0]][currentSpace[1]])
	}
	fmt.Println(currentSpace)
}

func moveRight(stepsToMove int, currentSpace []int) []int {
	if currentSpace[1]+stepsToMove < 10 {
		currentSpace[1] = currentSpace[1] + stepsToMove
		fmt.Println("Moved right. Current space is ", currentSpace)
		return currentSpace
	}
	fmt.Println("Cant go there, move would take you outside the board")
	return currentSpace
}

func moveLeft(stepsToMove int, currentSpace []int) []int {
	if currentSpace[1]-stepsToMove > -1 {
		currentSpace[1] = currentSpace[1] - stepsToMove
		fmt.Println("Moved left. Current space is ", currentSpace)
		return currentSpace
	}
	fmt.Println("Cant go there, move would take you outside the board")
	return currentSpace
}

func moveDown(stepsToMove int, currentSpace []int) []int {
	if currentSpace[0]+stepsToMove < 10 {
		currentSpace[0] = currentSpace[0] + stepsToMove
		fmt.Println("Moved down. Current space is ", currentSpace)
		return currentSpace
	}
	fmt.Println("Cant go there, move would take you outside the board")
	return currentSpace
}

func moveUp(stepsToMove int, currentSpace []int) []int {
	if currentSpace[0]-stepsToMove > -1 {
		currentSpace[0] = currentSpace[0] - stepsToMove
		fmt.Println("Moved up. Current space is ", currentSpace)
		return currentSpace
	}
	fmt.Println("Cant go there, move would take you outside the board")
	return currentSpace
}
