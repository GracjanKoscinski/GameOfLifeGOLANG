package main

var board [boardSize][boardSize]bool

func howManyNeighbours(x int, y int) int {
	var count int
	if x-1 >= 0 && y-1 >= 0 && board[x-1][y-1] {
		count++
	}
	if y-1 >= 0 && board[x][y-1] {
		count++
	}
	if x+1 < boardSize && y-1 >= 0 && board[x+1][y-1] {
		count++
	}
	if x-1 >= 0 && board[x-1][y] {
		count++
	}
	if x+1 < boardSize && board[x+1][y] {
		count++
	}
	if x-1 >= 0 && y+1 < boardSize && board[x-1][y+1] {
		count++
	}
	if y+1 < boardSize && board[x][y+1] {
		count++
	}
	if x+1 < boardSize && y+1 < boardSize && board[x+1][y+1] {
		count++
	}
	return count
}

func nextGeneration() {
	var newBoard [boardSize][boardSize]bool
	for i := 0; i < boardSize; i++ {
		for j := 0; j < boardSize; j++ {
			neighbors := howManyNeighbours(i, j)
			if board[i][j] {
				newBoard[i][j] = neighbors == 2 || neighbors == 3
			} else {
				newBoard[i][j] = neighbors == 3
			}
		}
	}
	board = newBoard
}

func createGosperGliderGun() {
	// Left square
	board[5][1] = true
	board[5][2] = true
	board[6][1] = true
	board[6][2] = true

	// Left part of the gun
	board[5][11] = true
	board[6][11] = true
	board[7][11] = true
	board[4][12] = true
	board[8][12] = true
	board[3][13] = true
	board[9][13] = true
	board[3][14] = true
	board[9][14] = true
	board[6][15] = true
	board[4][16] = true
	board[8][16] = true
	board[5][17] = true
	board[6][17] = true
	board[7][17] = true
	board[6][18] = true

	// Right part of the gun
	board[3][21] = true
	board[4][21] = true
	board[5][21] = true
	board[3][22] = true
	board[4][22] = true
	board[5][22] = true
	board[2][23] = true
	board[6][23] = true
	board[1][25] = true
	board[2][25] = true
	board[6][25] = true
	board[7][25] = true

	// Right square
	board[3][35] = true
	board[4][35] = true
	board[3][36] = true
	board[4][36] = true
}
