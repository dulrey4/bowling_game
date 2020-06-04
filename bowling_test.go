package main

import (
	"testing"

	"github.com/pkg/errors"
)

//TODO add more tests for individual functions/methods & more test cases for entire game. Also need to test isFinished.

type testCase struct {
	testName               string
	rolls                  []int
	expectedFinalScore     int
	expectedErr            error
	expectedErrorRollIndex int
}

func TestVariousGames(t *testing.T) {
	testCases := []testCase{
		{
			testName:           "Single roll",
			rolls:              []int{5},
			expectedFinalScore: 5,
		},
		{
			testName:           "Perfect game",
			rolls:              []int{10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10}, //12 rolls, 2 bonus at the end.
			expectedFinalScore: 300,
		},
		{
			testName:           "All spares",
			rolls:              []int{5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5}, //21 rolls, 1 bonus at the end.
			expectedFinalScore: 150,
		},
		{
			testName:           "All spares",
			rolls:              []int{2, 3, 4, 1, 5, 0, 0, 5, 3, 2, 1, 4, 0, 5, 5, 0, 3, 2, 10, 5, 5}, //Average game ending with a strike.
			expectedFinalScore: 65,
		},
		{
			testName:               "Error game over",
			rolls:                  []int{10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10}, //13 rolls, last one game is over
			expectedFinalScore:     300,
			expectedErrorRollIndex: 12,
			expectedErr:            errors.New("game over"),
		},
	}

	for _, tc := range testCases {
		g := NewGame()
		for i, roll := range tc.rolls {
			err := g.acceptRoll(roll)
			if err != nil {
				if tc.expectedErr.Error() != err.Error() || tc.expectedErrorRollIndex != i {
					t.Logf("Test: %s", tc.testName)
					t.Errorf("Unexpected error: %s", err.Error())
				}
			}
		}
		finalScore := g.getScore()
		if tc.expectedFinalScore != finalScore {
			t.Errorf("Expected final score (%d) actual score (%d)", tc.expectedErr, finalScore)
		}
	}
}
