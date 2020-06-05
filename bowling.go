package main

import (
	"fmt"
	"github.com/pkg/errors"
)

const framesPerGame = 10
const pinCount = 10

type BowlingGame interface{
	getRemainingRollsForCurrentFrame() int
	acceptRoll(int) error
	isFinished() bool
	getScore() int

}


func NewBowlingGame() BowlingGame {
	return &defaultBowlingGame{}
}

type defaultBowlingGame struct {
	frames []frame
}

/*
	frame stores data associated with a given set of pins. A frame is considered finished when either two rolls have been made, or a strike has occurred.
	Evaluation of the bonus score is not required to mark a frame finished.
*/
type frame struct {
	rollScore     int
	bonusScore    int
	acceptedRolls int
	finished      bool
	bonusRolls    int
}


func newFrame(roll int) frame{
	f := frame{}
	f.acceptedRolls += 1
	f.rollScore += roll
	if f.isStrike() {
		f.finished = true
		f.bonusRolls = 2
	}
	return f
}

func (f *frame) finishFrame(roll int){
	f.finished = true
	f.acceptedRolls += 1
	f.rollScore += roll
	if f.isSpare() {
		f.bonusRolls = 1
	}
}

func (f *frame) getTotalScore() int {
	return f.rollScore + f.bonusScore
}

func (f *frame) isStrike() bool {
	return f.acceptedRolls == 1 && f.rollScore == pinCount
}

func (f *frame) isSpare() bool {
	return f.acceptedRolls == 2 && f.rollScore == pinCount
}

func (g *defaultBowlingGame) getRemainingRollsForCurrentFrame() int {
	currentFrame := g.frames[len(g.frames)-1]
	//For the last frame, consider bonus rolls to be remaining rolls in that frame.
	if len(g.frames) == framesPerGame && (currentFrame.isStrike() || currentFrame.isSpare()) {
		return currentFrame.bonusRolls
	}
	return 2 - currentFrame.acceptedRolls
}

func (g *defaultBowlingGame) getScore() int {
	var score int
	for _, frame := range g.frames {
		score += frame.getTotalScore()
	}
	return score
}

func (g *defaultBowlingGame) isFinished() bool {
	return allFramesFinished(g.frames) && !bonusesRemain(g.frames)
}

func (g *defaultBowlingGame) acceptRoll(roll int) error {
	if g.isFinished() {
		return errors.New("game over")
	}

	addBonuses(g.frames, roll)

	if !allFramesFinished(g.frames) {
		if currentFrameFinished(g.frames) {
			g.frames = append(g.frames, newFrame(roll))
		} else {
			g.frames[len(g.frames)-1].finishFrame(roll)
		}
	}
	return nil
}

func currentFrameFinished(frames []frame) bool{
	return len(frames) == 0 || frames[len(frames)-1].finished
}

func allFramesFinished(frames []frame) bool {
	return len(frames) == framesPerGame && frames[framesPerGame-1].finished
}

func bonusesRemain(frames []frame) bool {
	for _, frame := range frames {
		if frame.bonusRolls != 0 {
			return true
		}
	}
	return false
}

func addBonuses(frames []frame, roll int) {
	for i, frame := range frames {
		if frame.bonusRolls > 0 {
			frames[i].bonusScore += roll
			frames[i].bonusRolls = frames[i].bonusRolls - 1
		}
	}
}

func main() {
	g := NewBowlingGame()
	for i := 0; i < 13; i++ {
		err := g.acceptRoll(10)
		if err != nil {
			fmt.Printf("acceptRoll failed with error: %s", err.Error())
			fmt.Println("")
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
