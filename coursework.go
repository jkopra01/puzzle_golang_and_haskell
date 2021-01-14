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
	var currentSpace = []int{0, 0}
	fmt.Println(currentSpace)
	endingSpace := []int{9, 9} //9,9
	fmt.Println(endingSpace)
	spacesVisited := [][]int{}
	BadSpaces := [][]int{}
	var spaceToBeAdded = []int{currentSpace[0], currentSpace[1]}
	//spacesVisited = append(spacesVisited, spaceToBeAdded)

	//fmt.Println(spaceHasBeenVisited(spacesVisited, currentSpace))
	//moveRight(boardGrid[currentSpace[0]][currentSpace[1]], currentSpace)

	//fmt.Println(spaceHasBeenVisited(spacesVisited, currentSpace))
	//moveLeft(boardGrid[currentSpace[0]][currentSpace[1]], currentSpace, spacesVisited)

	//moveDown(boardGrid[currentSpace[0]][currentSpace[1]], currentSpace)
	//moveUp(boardGrid[currentSpace[0]][currentSpace[1]], currentSpace)
	//fmt.Println(currentSpace, boardGrid[currentSpace[0]][currentSpace[1]])
	//

	for !(currentSpace[0] == endingSpace[0] && currentSpace[1] == endingSpace[1]) {
		//for i := 0; i < 30; i++ {

		var movementOptionForRightMovement = []int{currentSpace[0], currentSpace[1] + boardGrid[currentSpace[0]][currentSpace[1]]}
		var movementOptionForLeftMovement = []int{currentSpace[0], currentSpace[1] - boardGrid[currentSpace[0]][currentSpace[1]]}
		var movementOptionForDownMovement = []int{currentSpace[0] + boardGrid[currentSpace[0]][currentSpace[1]], currentSpace[1]}
		var movementOptionForUpMovement = []int{currentSpace[0] - boardGrid[currentSpace[0]][currentSpace[1]], currentSpace[1]}

		/* 	fmt.Println("mikä olet left", movementOptionForLeftMovement)
		fmt.Println("mikä olet raitti", movementOptionForRightMovement)
		fmt.Println("mikä olet down", movementOptionForDownMovement)
		fmt.Println("mikä olet uppi", movementOptionForUpMovement)
		*/
		var rightOption = spaceHasBeenVisited(spacesVisited, movementOptionForRightMovement)
		var leftOption = spaceHasBeenVisited(spacesVisited, movementOptionForLeftMovement)
		var downOption = spaceHasBeenVisited(spacesVisited, movementOptionForDownMovement)
		var upOption = spaceHasBeenVisited(spacesVisited, movementOptionForUpMovement)
		var badRightOption = isBadSpace(BadSpaces, movementOptionForRightMovement)
		var badLeftOption = isBadSpace(BadSpaces, movementOptionForLeftMovement)
		var badDownOption = isBadSpace(BadSpaces, movementOptionForDownMovement)
		var badUpOption = isBadSpace(BadSpaces, movementOptionForUpMovement)

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

			//fmt.Println("tescccct")
			//fmt.Println(currentSpace)
			var temp = currentSpace
			BadSpaces = append(BadSpaces, temp)
			//BadSpaces = append(BadSpaces, temp)
			spacesVisited = nil
			var newSpace = []int{0, 0}
			currentSpace = newSpace
			//	fmt.Println(currentSpace)
		}
	}

	fmt.Println(currentSpace)
	fmt.Println("route to exit")
	fmt.Println(spacesVisited)

}

func moveRight(stepsToMove int, currentSpace []int, spacesVisited [][]int) ([]int, [][]int) {
	if currentSpace[1]+stepsToMove < 10 {
		currentSpace[1] = currentSpace[1] + stepsToMove
		fmt.Println("Moved right. Current space is ", currentSpace)
		return currentSpace, spacesVisited
	}
	fmt.Println("Cant go there, move would take you outside the board")
	return currentSpace, spacesVisited
}

func moveLeft(stepsToMove int, currentSpace []int, spacesVisited [][]int) ([]int, [][]int) {
	if currentSpace[1]-stepsToMove > -1 {
		currentSpace[1] = currentSpace[1] - stepsToMove
		fmt.Println("Moved left. Current space is ", currentSpace)
		var spaceToBeAdded = []int{currentSpace[0], currentSpace[1]}
		spacesVisited = append(spacesVisited, spaceToBeAdded)
		return currentSpace, spacesVisited
	}
	fmt.Println("Cant go there, move would take you outside the board")
	return currentSpace, spacesVisited
}

func moveDown(stepsToMove int, currentSpace []int, spacesVisited [][]int) ([]int, [][]int) {
	if currentSpace[0]+stepsToMove < 10 {
		currentSpace[0] = currentSpace[0] + stepsToMove
		fmt.Println("Moved down. Current space is ", currentSpace)
		spacesVisited = append(spacesVisited, currentSpace)
		return currentSpace, spacesVisited
	}
	fmt.Println("Cant go there, move would take you outside the board")
	return currentSpace, spacesVisited
}

func moveUp(stepsToMove int, currentSpace []int, spacesVisited [][]int) ([]int, [][]int) {
	if currentSpace[0]-stepsToMove > -1 {
		currentSpace[0] = currentSpace[0] - stepsToMove
		fmt.Println("Moved up. Current space is ", currentSpace)
		spacesVisited = append(spacesVisited, currentSpace)
		return currentSpace, spacesVisited
	}
	fmt.Println("Cant go there, move would take you outside the board")
	return currentSpace, spacesVisited
}

// Contains tells whether a contains x.
func spaceHasBeenVisited(spacesVisited [][]int, currentSpace []int) bool {
	for _, r := range spacesVisited { //assign

		if currentSpace[0] == r[0] && currentSpace[1] == r[1] {
			return true
		}
	}
	return false
}

// Contains tells whether a contains x.
func isBadSpace(badSpaces [][]int, currentSpace []int) bool {
	for _, r := range badSpaces { //assign
		if currentSpace[0] == r[0] && currentSpace[1] == r[1] {
			return true
		}
	}
	return false
}
