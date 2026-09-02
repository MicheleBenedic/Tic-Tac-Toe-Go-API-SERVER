/* Verifica che la mossa venga registrata correttamente.
Rileva e rifiuta mosse illegali restituendo errore o valore booleano.
Verifica il comportamento in caso di sovrascrittura o indice fuori limite. */

package game

import "testing"

func TestInitBoardVoidCells(t *testing.T) {
	board := [BOARD_SIZE][BOARD_SIZE]rune{}
	init_board(&board)

	for row := 0; row < BOARD_SIZE; row++ {
		for col := 0; col < BOARD_SIZE; col++ {
			if board[row][col] != ' ' {
				t.Fatalf("The cell %d, %d is not empty", row, col)
			}
		}
	}
}

func TestIsValidMove(t *testing.T) {
	board := [BOARD_SIZE][BOARD_SIZE]rune{}
	init_board(&board)

	// caso 1: cella vuota e dentro il bordo -> valida
	if !IsValidMove(board, 0, 0) {
		t.Fatal("expected move (0,0) to be valid on an empty board")
	}

	// caso 2: cella occupata -> non valida
	board[0][0] = 'X'
	if IsValidMove(board, 0, 0) {
		t.Fatal("expected move (0,0) to be invalid when already occupied")
	}

	// caso 3: riga fuori range -> non valida
	if IsValidMove(board, BOARD_SIZE, 0) {
		t.Fatal("expected move outside board row to be invalid")
	}

	// caso 4: colonna fuori range -> non valida
	if IsValidMove(board, 0, BOARD_SIZE) {
		t.Fatal("expected move outside board column to be invalid")
	}
}

func TestVictoryX(t *testing.T) {
	/* Es: "una riga completa di X deve risultare come vittoria di X". */
	board := [BOARD_SIZE][BOARD_SIZE]rune{}
	board[0][0] = 'X'
	board[0][1] = 'X'
	board[0][2] = 'X'

	if winner(board) != 'X' {
		t.Fatal("expected the 'X' to completely occupy the first row")
	}
}

func TestVictoryO(t *testing.T) {
	board := [BOARD_SIZE][BOARD_SIZE]rune{}
	board[0][0] = 'O'
	board[0][1] = 'O'
	board[0][2] = 'O'

	if winner(board) != 'O' {
		t.Fatal("expected the 'O' to completely occupy the first col")
	}
}

func TestDraw(t *testing.T) {
	board := [BOARD_SIZE][BOARD_SIZE]rune{}
	board[0][0] = 'X'
	board[0][1] = 'O'
	board[0][2] = 'X'
	board[1][0] = 'O'
	board[1][1] = 'O'
	board[1][2] = 'X'
	board[2][0] = 'X'
	board[2][1] = 'X'
	board[2][2] = 'O'

	if winner(board) == 'X' {
		t.Fatal("Expected no winner in this game!")
	}
	if winner(board) == 'O' {
		t.Fatal("Expected no winner in ths gfame!")
	}
}

func TestIllegalMoves(t *testing.T) {
	board := [BOARD_SIZE][BOARD_SIZE]rune{}
	board[0][0] = 'X'
	if IsValidMove(board, 0, 0) {
		t.Fatal("expected (0, 0) to be an invalid move since it is already occupied by 'X'")
	}
}
