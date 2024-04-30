package main

import (
	"fmt"
	"strings"
)

func main() {
	//create tic-tac toe board
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	//the players take turn
	board[0][0] = "X"
	board[2][2] = "0"
	board[1][2] = "X"
	board[1][0] = "X"
	board[0][2] = "0"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s \n", strings.Join(board[i], " "))
	}
}
