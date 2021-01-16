package main

import (
	"testing"
)

func Test_Board(t *testing.T) {
	boardInPicture := [][]int{{7, 4, 4, 6, 6, 3, 2, 2, 6, 8}, {3, 3, 6, 5, 4, 3, 7, 2, 8, 3}, {4, 1, 6, 6, 2, 4, 4, 4, 7, 4}, {4, 5, 3, 4, 3, 5, 4, 4, 8, 5}, {5, 1, 4, 6, 6, 5, 0, 7, 1, 4}, {2, 6, 9, 4, 9, 7, 7, 9, 1, 4}, {3, 5, 4, 0, 6, 4, 5, 5, 5, 6}, {6, 6, 2, 3, 4, 7, 1, 2, 3, 3}, {3, 5, 4, 3, 6, 5, 4, 5, 2, 6}, {3, 9, 3, 5, 1, 1, 5, 4, 6, 0}}
	createdBoard := createBoard()

	for r := range boardInPicture {
		for i := range createdBoard {
			println(createdBoard[r][i])
			println("ffff")
			println(boardInPicture[r][i])
			if createdBoard[r][i] != boardInPicture[r][i] {
				t.Fatalf("\nBoard created was incorrect.\nReceived board: %v \nCorrect board: %v", createdBoard, boardInPicture)
			}
		}
	}
}

func TestMovement(t *testing.T) {
	currentSpace := []int{0, 0}
	spacesVisited := [][]int{}
	correctSpace := []int{5, 3}

	moveRight(7, currentSpace, spacesVisited)
	moveDown(2, currentSpace, spacesVisited)
	moveLeft(4, currentSpace, spacesVisited)
	moveDown(6, currentSpace, spacesVisited)
	moveUp(3, currentSpace, spacesVisited)
	if !(currentSpace[0] == correctSpace[0] && currentSpace[1] == correctSpace[1]) {
		t.Fatalf("\nEnded up on a wrong space.\nReceived space: %v \nCorrect space: %v", currentSpace, correctSpace)
	}

}
func Test_correctExitSpace(t *testing.T) {
	createdBoard := createBoard()
	spacesVisited := solve(createdBoard)
	lastSpace := spacesVisited[len(spacesVisited)-1]
	correctExitSpace := []int{9, 9}
	if !(lastSpace[0] == correctExitSpace[0] && lastSpace[1] == correctExitSpace[1]) {
		t.Fatalf("\nExited on a wrong space.\nReceived ending space: %v \nCorrect exit space: %v", lastSpace, correctExitSpace)
	}
}
