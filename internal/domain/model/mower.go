package model

// Mower represents the state of a mower.
type Mower struct {
	X, Y        int
	Orientation string
}

// Move the mower based on the instruction.
func (mover *Mower) Move(instruction string) {
	switch instruction {
	case "R":
		mover.TurnRight()
	case "L":
		mover.TurnLeft()
	case "F":
		mover.MoveForward()
	}
}

// TurnRight the mower 90 degrees to the right.
func (mover *Mower) TurnRight() {
	switch mover.Orientation {
	case "N":
		mover.Orientation = "E"
	case "E":
		mover.Orientation = "S"
	case "S":
		mover.Orientation = "W"
	case "W":
		mover.Orientation = "N"
	}
}

// TurnLeft the mower 90 degrees to the left.
func (mover *Mower) TurnLeft() {
	switch mover.Orientation {
	case "N":
		mover.Orientation = "W"
	case "W":
		mover.Orientation = "S"
	case "S":
		mover.Orientation = "E"
	case "E":
		mover.Orientation = "N"
	}
}

// MoveForward the mower forward one space in the direction it faces.
func (mover *Mower) MoveForward() {
	switch mover.Orientation {
	case "N":
		mover.Y++
	case "E":
		mover.X++
	case "S":
		mover.Y--
	case "W":
		mover.X--
	}
}

// MoveBackward the mower backward one space to revert an invalid move.
func (mover *Mower) MoveBackward() {
	switch mover.Orientation {
	case "N":
		mover.Y--
	case "E":
		mover.X--
	case "S":
		mover.Y++
	case "W":
		mover.X++
	}
}
