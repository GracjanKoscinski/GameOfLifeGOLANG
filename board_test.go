package main

import (
	"testing"
)

func TestHowManyNeighbours(t *testing.T) {
	board = [boardSize][boardSize]bool{
		{false, true, false},
		{true, false, true},
		{false, true, false},
		{false, false, true},
	}

	tests := []struct {
		name     string
		x        int
		y        int
		expected int
	}{
		{"Center cell", 1, 1, 4},
		{"Top left cell", 0, 0, 2},
		{"Bottom right cell", 2, 2, 3},
		{"Bottom left cell", 0, 3, 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := howManyNeighbours(tt.x, tt.y); got != tt.expected {
				t.Errorf("howManyNeighbours() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func boardsAreEqual(a, b [boardSize][boardSize]bool) bool {
	for i := 0; i < boardSize; i++ {
		for j := 0; j < boardSize; j++ {
			if a[i][j] != b[i][j] {
				return false
			}
		}
	}
	return true
}

func TestNextGeneration(t *testing.T) {
	tests := []struct {
		name     string
		initial  [boardSize][boardSize]bool
		expected [boardSize][boardSize]bool
	}{
		{
			"Test 1",
			[boardSize][boardSize]bool{
				{false, false, false},
				{true, true, true},
				{false, false, false},
				{false, false, false},
			},
			[boardSize][boardSize]bool{
				{false, true, false},
				{false, true, false},
				{false, true, false},
				{false, false, false},
			},
		},
		{
			"Test 2",
			[boardSize][boardSize]bool{
				{true, true, false},
				{true, true, true},
				{false, true, true},
			},
			[boardSize][boardSize]bool{
				{true, false, true},
				{false, false, false},
				{true, false, true},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			board = tt.initial
			nextGeneration()
			if !boardsAreEqual(board, tt.expected) {
				t.Errorf("Unexpected board state, got: %v, want: %v", board, tt.expected)
			}
		})
	}
}
