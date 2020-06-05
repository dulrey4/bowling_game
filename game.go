package bowling_game

import (
	"github.com/pkg/errors"
)

const framesPerGame = 10

type BowlingGame interface {
	GetRemainingRollsForCurrentFrame() int
	AcceptRoll(int) error
	Finished() bool
	GetScore() int
}

func NewBowlingGame() BowlingGame {
	return &defaultBowlingGame{}
}

type defaultBowlingGame struct {
	frames []frame
}

func (g *defaultBowlingGame) GetRemainingRollsForCurrentFrame() int {
	currentFrame := g.frames[len(g.frames)-1]
	//For the last frame, consider bonus rolls to be remaining rolls in that frame.
	if len(g.frames) == framesPerGame && (currentFrame.isStrike() || currentFrame.isSpare()) {
		return currentFrame.bonusRolls
	}
	return 2 - currentFrame.acceptedRolls
}

func (g *defaultBowlingGame) GetScore() int {
	var score int
	for _, frame := range g.frames {
		score += frame.getTotalScore()
	}
	return score
}

func (g *defaultBowlingGame) Finished() bool {
	return allFramesFinished(g.frames) && !bonusesRemain(g.frames)
}

func (g *defaultBowlingGame) AcceptRoll(roll int) error {
	if g.Finished() {
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

func currentFrameFinished(frames []frame) bool {
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
