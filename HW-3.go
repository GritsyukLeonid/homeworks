package main

import (
	"fmt"
)

func printChessboard(size int) {
	board := ""
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if (i+j)%2 == 0 {
				board += "#"
			} else {
				board += " "
			}
		}
		board += "\n"
	}

	fmt.Println(board)
}

func main() {
	var size int
	fmt.Print("Введите размер доски: ")
	fmt.Scan(&size)
	printChessboard(size)
}
