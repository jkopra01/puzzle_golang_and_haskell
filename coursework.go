package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {
	var boardGrid [][]int
	boardGrid = createBoard()
	printBoard(boardGrid)
	solve(boardGrid)
}

func createBoard() [][]int {
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
	return boardGrid
}

func printBoard(boardGrid [][]int) {
	fmt.Println("Board:")
	for _, r := range boardGrid {
		fmt.Println(r)
	}
}
func solve(boardGrid [][]int) [][]int {
	var currentSpace = []int{0, 0}
	endingSpace := []int{9, 9}
	spacesVisited := [][]int{}
	BadSpaces := [][]int{}
	var spaceToBeAdded = []int{currentSpace[0], currentSpace[1]}
	for !(currentSpace[0] == endingSpace[0] && currentSpace[1] == endingSpace[1]) {
		var movementOptionForRightMovement = []int{currentSpace[0], currentSpace[1] + boardGrid[currentSpace[0]][currentSpace[1]]}
		var movementOptionForLeftMovement = []int{currentSpace[0], currentSpace[1] - boardGrid[currentSpace[0]][currentSpace[1]]}
		var movementOptionForDownMovement = []int{currentSpace[0] + boardGrid[currentSpace[0]][currentSpace[1]], currentSpace[1]}
		var movementOptionForUpMovement = []int{currentSpace[0] - boardGrid[currentSpace[0]][currentSpace[1]], currentSpace[1]}
		var rightOption = spaceHasBeenVisited(spacesVisited, movementOptionForRightMovement)
		var leftOption = spaceHasBeenVisited(spacesVisited, movementOptionForLeftMovement)
		var downOption = spaceHasBeenVisited(spacesVisited, movementOptionForDownMovement)
		var upOption = spaceHasBeenVisited(spacesVisited, movementOptionForUpMovement)
		var badRightOption = isBadSpace(BadSpaces, movementOptionForRightMovement)
		var badLeftOption = isBadSpace(BadSpaces, movementOptionForLeftMovement)
		var badDownOption = isBadSpace(BadSpaces, movementOptionForDownMovement)
		var badUpOption = isBadSpace(BadSpaces, movementOptionForUpMovement)

		//check if movement stays in the board, has not been marked as a bad space and if it has been visited this "run"
		if currentSpace[1]+boardGrid[currentSpace[0]][currentSpace[1]] < 10 && !badRightOption && !rightOption {
			moveRight(boardGrid[currentSpace[0]][currentSpace[1]], currentSpace, spacesVisited)
			spaceToBeAdded = []int{currentSpace[0], currentSpace[1]}
			spacesVisited = append(spacesVisited, spaceToBeAdded)
		} else if currentSpace[0]+boardGrid[currentSpace[0]][currentSpace[1]] < 10 && !badDownOption && !downOption {
			moveDown(boardGrid[currentSpace[0]][currentSpace[1]], currentSpace, spacesVisited)
			spaceToBeAdded = []int{currentSpace[0], currentSpace[1]}
			spacesVisited = append(spacesVisited, spaceToBeAdded)
		} else if currentSpace[0]-boardGrid[currentSpace[0]][currentSpace[1]] > -1 && !badUpOption && !upOption {
			moveUp(boardGrid[currentSpace[0]][currentSpace[1]], currentSpace, spacesVisited)
			spaceToBeAdded = []int{currentSpace[0], currentSpace[1]}
			spacesVisited = append(spacesVisited, spaceToBeAdded)
		} else if currentSpace[1]-boardGrid[currentSpace[0]][currentSpace[1]] > -1 && !badLeftOption && !leftOption {
			moveLeft(boardGrid[currentSpace[0]][currentSpace[1]], currentSpace, spacesVisited)
			spaceToBeAdded = []int{currentSpace[0], currentSpace[1]}
			spacesVisited = append(spacesVisited, spaceToBeAdded)
		} else {
			//no good spaces to move to, mark current space as a bad space, empty visited spaces and go back to the starting space.
			BadSpaces = append(BadSpaces, currentSpace)
			spacesVisited = nil
			currentSpace = []int{0, 0}
		}
	}
	fmt.Println("Route to the exit was:\n", spacesVisited)
	return spacesVisited
}

func moveRight(stepsToMove int, currentSpace []int, spacesVisited [][]int) ([]int, [][]int) {
	currentSpace[1] = currentSpace[1] + stepsToMove
	fmt.Println("Moved right. Current space is ", currentSpace)
	spacesVisited = append(spacesVisited, currentSpace)
	return currentSpace, spacesVisited
}

func moveLeft(stepsToMove int, currentSpace []int, spacesVisited [][]int) ([]int, [][]int) {
	currentSpace[1] = currentSpace[1] - stepsToMove
	fmt.Println("Moved left. Current space is ", currentSpace)
	spacesVisited = append(spacesVisited, currentSpace)
	return currentSpace, spacesVisited
}

func moveDown(stepsToMove int, currentSpace []int, spacesVisited [][]int) ([]int, [][]int) {
	currentSpace[0] = currentSpace[0] + stepsToMove
	fmt.Println("Moved down. Current space is ", currentSpace)
	spacesVisited = append(spacesVisited, currentSpace)
	return currentSpace, spacesVisited
}

func moveUp(stepsToMove int, currentSpace []int, spacesVisited [][]int) ([]int, [][]int) {
	currentSpace[0] = currentSpace[0] - stepsToMove
	fmt.Println("Moved up. Current space is ", currentSpace)
	spacesVisited = append(spacesVisited, currentSpace)
	return currentSpace, spacesVisited
}

func spaceHasBeenVisited(spacesVisited [][]int, currentSpace []int) bool {
	for _, r := range spacesVisited {
		if currentSpace[0] == r[0] && currentSpace[1] == r[1] {
			return true
		}
	}
	return false
}

func isBadSpace(badSpaces [][]int, currentSpace []int) bool {
	for _, r := range badSpaces {
		if currentSpace[0] == r[0] && currentSpace[1] == r[1] {
			return true
		}
	}
	return false
}
