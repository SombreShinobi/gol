package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"time"
)

func main() {
	game()
}

func game() {
	board := gen_board()

	for {
		clear_screen()
		advance_gen(&board)
		time.Sleep(time.Second / 10)
	}
}

func clear_screen() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func gen_board() [20][20]int {
	board := [20][20]int{}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			board[i][j] = rand.Intn(2)
		}
	}

	return board
}

func advance_gen(board *[20][20]int) {
	prev := *board
	for i := 0; i < len(prev); i++ {
		for j := 0; j < len(prev[i]); j++ {
			fmt.Print(prev[i][j], " ")

			neighbours := count_neighbours(&prev, i, j)

			if cell(&prev, i, j) == 1 && (neighbours < 2 || neighbours > 3) {
				board[i][j] = 0
			} else if cell(&prev, i, j) == 0 && neighbours == 3 {
				board[i][j] = 1
			}
		}
		fmt.Print("\n")
	}
}

func count_neighbours(board *[20][20]int, i, j int) int {
	return cell(board, i-1, j-1) + cell(board, i-1, j) + cell(board, i-1, j+1) +
		cell(board, i, j-1) + cell(board, i, j+1) +
		cell(board, i+1, j-1) + cell(board, i+1, j) + cell(board, i+1, j+1)
}

func cell(board *[20][20]int, i, j int) int {
	i = mod(i, len(board))
	j = mod(j, len(board[i]))

	return board[i][j]
}

// Workaround since go doesn't support the modulus operation
func mod(x, y int) int {
	return (x%y + y) % y
}
