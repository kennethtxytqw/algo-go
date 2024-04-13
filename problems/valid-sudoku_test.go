package problems

import "testing"

func TestIsValidSudoku(t *testing.T) {
	testCases := []struct {
		name   string
		board  [][]byte
		expect bool
	}{
		{
			name: "Test 1: Valid Sudoku",
			board: [][]byte{
				{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
				{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
				{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
				{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
				{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
				{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
				{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
				{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
				{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
			},
			expect: true,
		},
		{
			name: "Test 2: Invalid Sudoku",
			board: [][]byte{
				{'8', '3', '.', '.', '7', '.', '.', '.', '.'},
				{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
				{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
				{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
				{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
				{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
				{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
				{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
				{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
			},
			expect: false,
		},
		{
			name: "Test 3: Invalid Sudoku",
			board: [][]byte{
				{'.', '.', '4', '.', '.', '.', '6', '3', '.'},
				{'.', '.', '.', '.', '.', '.', '.', '.', '.'},
				{'5', '.', '.', '.', '.', '.', '.', '9', '.'},
				{'.', '.', '.', '5', '6', '.', '.', '.', '.'},
				{'4', '.', '3', '.', '.', '.', '.', '.', '1'},
				{'.', '.', '.', '7', '.', '.', '.', '.', '.'},
				{'.', '.', '.', '5', '.', '.', '.', '.', '.'},
				{'.', '.', '.', '.', '.', '.', '.', '.', '.'},
				{'.', '.', '.', '.', '.', '.', '.', '.', '.'},
			},
			expect: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if got := isValidSudoku(tc.board); got != tc.expect {
				t.Errorf("isValidSudoku() = %v, want %v", got, tc.expect)
			}
		})
	}
}
