package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"reflect"
	"strconv"
	"strings"
)

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
	fmt.Println(reflect.TypeOf(boardGrid))
}
