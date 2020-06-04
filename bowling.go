package main

import (
	"fmt"
	"github.com/pkg/errors"
)

const framesPerGame = 10
const pinCount = 10

/*
	game is a bowling game. It is composed of a series of frames, of which there are ten.
*/
type game struct {
	frames []frame
}

func NewGame() game {
	return game{}
}

/*
	frame stores data associated with a given set of fresh pins.	A frame is considered finished when either two rolls have been made, or a strike has occurred.
	Evaluation of the bonus score is not required to mark a frame finished.
*/
type frame struct {
	rollScore     int
	bonusScore    int
	acceptedRolls int
	finished      bool
	bonus         int
}

func (f *frame) getTotalScore() int {
	return f.rollScore + f.bonusScore
}

func (g *game) getCurrentFrame() (frame, error) {
	if len(g.frames) == 0 {
		return frame{}, errors.New("no current frame")
	}
	return g.frames[len(g.frames)-1], nil
}

func (g *game) getRemainingRollsForCurrentFrame() int {
	currentFrame := g.frames[len(g.frames)-1]
	//For the last frame, consider bonus rolls to be remaining rolls in that frame.
	if len(g.frames) == framesPerGame && (currentFrame.isStrike() || currentFrame.isSpare()) {
		return currentFrame.bonus
	}
	return 2 - currentFrame.acceptedRolls
}

func (f *frame) isStrike() bool {
	return f.acceptedRolls == 1 && f.rollScore == pinCount
}

func (f *frame) isSpare() bool {
	return f.acceptedRolls == 2 && f.rollScore == pinCount
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
			f.rollScore += roll
			if f.isStrike() {
				f.finished = true
				f.bonus = 2
			}
			g.frames = append(g.frames, f)
		} else {
			frame := g.frames[len(g.frames)-1]
			frame.finished = true
			frame.acceptedRolls += 1
			frame.rollScore += roll
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
			frames[i].bonusScore += roll
			frames[i].bonus = frames[i].bonus - 1
		}
	}
}

func (g *game) getScore() int {
	var score int
	for _, frame := range g.frames {
		score += frame.getTotalScore()
	}
	return score
}

func main() {
	g := NewGame()
	for i := 0; i < 12; i++ {
		err := g.acceptRoll(10)
		if err != nil {
			fmt.Printf("acceptRoll failed with error: %s", err.Error())
			continue
		}
		fmt.Printf("Remaining rolls in current frame: %d", g.getRemainingRollsForCurrentFrame())
		fmt.Println("")
		fmt.Printf("Current score: %d", g.getScore())
		fmt.Println("")
		fmt.Printf("Game finished: %t", g.isFinished())
		fmt.Println("")
	}
}
