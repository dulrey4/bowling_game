package main

import (
	"fmt"

	"github.com/pkg/errors"
)

type game struct {
	frames []frame
}

func NewGame() game {
	return game{}
}

type frame struct {
	totalScore    int
	acceptedRolls int
	finished      bool
	bonus         int
}

func (g *game) getCurrentFrame() (frame, error) {
	if len(g.frames) == 0 {
		return frame{}, errors.New("no current frame")
	}
	return g.frames[len(g.frames)-1], nil
}

func (f *frame) remainingRolls() int {
	if f.finished {
		return 0
	}
	return 2 - f.acceptedRolls
}

func (f *frame) isStrike() bool {
	return f.acceptedRolls == 1 && f.totalScore == 10
}

func (f *frame) isSpare() bool {
	return f.acceptedRolls == 2 && f.totalScore == 10
}

func (g *game) allFramesFinished() bool {
	return len(g.frames) > 9 && g.frames[9].finished
}

func (g *game) isFinished() bool {
	return g.allFramesFinished() && !bonusesRemain(g.frames)
}

func (g *game) acceptRoll(roll int) error {
	if g.isFinished() {
		return errors.New("game over")
	}
	addBonuses(g.frames, roll)

	if !g.allFramesFinished() {
		if len(g.frames) == 0 || g.frames[len(g.frames)-1].finished {
			f := frame{}
			f.acceptedRolls += 1
			f.totalScore += roll
			if f.isStrike() {
				f.finished = true
				f.bonus = 2
			}
			g.frames = append(g.frames, f)
		} else {
			frame := g.frames[len(g.frames)-1]
			frame.finished = true
			frame.acceptedRolls += 1
			frame.totalScore += roll
			if frame.isSpare() {
				frame.bonus = 1
			}
			g.frames[len(g.frames)-1] = frame
		}
	}
	return nil
}

func bonusesRemain(frames []frame) bool {
	for _, frame := range frames {
		if frame.bonus != 0 {
			return true
		}
	}
	return false
}

func addBonuses(frames []frame, roll int) {
	for i, frame := range frames {
		if frame.bonus > 0 {
			frames[i].totalScore = frames[i].totalScore + roll
			frames[i].bonus = frames[i].bonus - 1
		}
	}
}

func (g *game) getScore() int {
	var score int
	for _, frame := range g.frames {
		score += frame.totalScore
	}
	return score
}

func main() {
	g := NewGame()
	for i := 0; i < 22; i++ {
		err := g.acceptRoll(5)
		if err != nil {
			fmt.Printf("acceptRoll failed with error: %s", err.Error())
			continue
		}
		currentFrame, err := g.getCurrentFrame()
		if err != nil {
			fmt.Printf("getCurrentFrame failed with error: %s", err.Error())
			continue
		}
		fmt.Printf("Remaining rolls in frame: %d", currentFrame.remainingRolls())
		fmt.Println("")
		fmt.Printf("Current score: %d", g.getScore())
		fmt.Println("")
		fmt.Printf("Game finished: %t", g.isFinished())
		fmt.Println("")
	}
}
