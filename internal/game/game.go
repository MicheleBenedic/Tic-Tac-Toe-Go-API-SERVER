// qui metti la logica del tris (quella che hai in game.c)
package game

import "fmt"

const BOARD_SIZE int = 3

var board [BOARD_SIZE][BOARD_SIZE]rune

type action struct {
	row int
	col int
}

func IsValidMove(board [BOARD_SIZE][BOARD_SIZE]rune, row int, col int) bool {
	// Check if row and col are within bounds
	if row < 0 || row >= BOARD_SIZE || col < 0 || col >= BOARD_SIZE {
		return false
	}
	// Check if cell is empty
	return board[row][col] == ' '
}

func init_board(board *[BOARD_SIZE][BOARD_SIZE]rune) {
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

func read_index(label string, out *int) bool {
	var input string
	fmt.Print(label + ": ")
	_, err := fmt.Scanln(&input)
	if err != nil {
		return false
	}

	var value int
	_, err = fmt.Sscanf(input, "%d", &value)
	if err != nil {
		return false
	}
	if value < 0 || value >= BOARD_SIZE {
		return false
	}

	*out = value
	return true
}

func player_number(marker rune) int {
	if marker == 'X' {
		return 1
	} else {
		return 2
	}
}

func player_move(marker rune, board [BOARD_SIZE][BOARD_SIZE]rune) {
	print_board(board)
	fmt.Printf("Player %d (%c), insert the coordinates of your move:\n",
		player_number(marker), marker)

	for {
		var move action

		if !read_index("row (0-2)", &move.row) ||
			!read_index("column (0-2)", &move.col) {
			fmt.Println("Coordinates not available, try again.")
			continue
		}
		if !IsValidMove(board, move.row, move.col) {
			fmt.Println("Square already occupied or invalid coordinates, insert new coordinates.")
			continue
		}

		board[move.row][move.col] = marker
		return
	}
}

func winner(board [BOARD_SIZE][BOARD_SIZE]rune) rune {
	for i := 0; i < BOARD_SIZE; i++ {
		if board[i][0] != ' ' && board[i][0] == board[i][1] && board[i][1] == board[i][2] {
			return board[i][0]
		}
		if board[0][i] != ' ' && board[0][i] == board[1][i] && board[1][i] == board[2][i] {
			return board[0][i]
		}
	}
	if board[0][0] != ' ' && board[0][0] == board[1][1] && board[1][1] == board[2][2] {
		return board[0][0]
	}
	if board[0][2] != ' ' && board[0][2] == board[1][1] && board[1][1] == board[2][0] {
		return board[0][2]
	}

	return ' '
}

func board_full(board [BOARD_SIZE][BOARD_SIZE]rune) bool {
	for row := 0; row < BOARD_SIZE; row++ {
		for col := 0; col < BOARD_SIZE; col++ {
			if board[row][col] == ' ' {
				return false
			}
		}
	}
	return true
}

func IsGameFinished(board [BOARD_SIZE][BOARD_SIZE]rune) bool {
	// Verifica se c'è un vincitore
	if winner(board) != ' ' {
		return true
	}
	// Verifica se la board è piena (pareggio)
	return board_full(board)
}
