package bowling_game

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
			testName:           "Normal game with strike at end",
			rolls:              []int{2, 3, 4, 1, 5, 0, 0, 5, 3, 2, 1, 4, 0, 5, 5, 0, 3, 2, 10, 5, 5},
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
		t.Logf("Test: %s", tc.testName)
		g := NewBowlingGame()
		for i, roll := range tc.rolls {
			err := g.AcceptRoll(roll)
			if err != nil && i != tc.expectedErrorRollIndex && tc.expectedErr == nil {
				t.Errorf("Got unexpected error %s", err.Error())
			} else if err == nil && i == tc.expectedErrorRollIndex && tc.expectedErr != nil {
				t.Error("expected error and got nil")
			}
		}
		finalScore := g.GetScore()
		if tc.expectedFinalScore != finalScore {
			t.Errorf("Expected final score (%d) actual score (%d)", tc.expectedFinalScore, finalScore)
		}
	}
}
