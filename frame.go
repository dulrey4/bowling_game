package bowling_game

const pinCount = 10

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

func newFrame(roll int) frame {
	f := frame{}
	f.acceptedRolls += 1
	f.rollScore += roll
	if f.isStrike() {
		f.finished = true
		f.bonusRolls = 2
	}
	return f
}

func (f *frame) finishFrame(roll int) {
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
