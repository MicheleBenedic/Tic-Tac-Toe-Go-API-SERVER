// qui metti la logica del tris (quella che hai in game.c)
package game

import "fmt"

const BOARD_SIZE int = 3

var board [BOARD_SIZE][BOARD_SIZE]rune

type action struct {
	row int
	col int
}

func init_board(board [BOARD_SIZE][BOARD_SIZE]rune) {
	for row := 0; row < BOARD_SIZE; row++ {
		for col := 0; col < BOARD_SIZE; col++ {
			board[row][col] = ' '
		}
	}
}

func print_board(board [BOARD_SIZE][BOARD_SIZE]rune) {
	for row := 0; row < BOARD_SIZE; row++ {
		for col := 0; col < BOARD_SIZE; col++ {
			fmt.Println(" %r ", board[row][col])
			if col < BOARD_SIZE-1 {
				fmt.Print("-----------\n")
			}
		}
		if row < BOARD_SIZE-1 {
			fmt.Print("-----------\n")
		}
	}
	fmt.Println()
}
