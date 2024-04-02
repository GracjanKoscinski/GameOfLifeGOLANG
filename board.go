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
