package model

import "testing"

// TestMowerMovement test the mower movement logic
func TestMowerMovement(t *testing.T) {
	// Test the rotation to the right
	m := Mower{0, 0, "N"}
	m.TurnRight()
	if m.Orientation != "E" {
		t.Errorf("Expected orientation: E, Got: %s", m.Orientation)
	}

	// Test the rotation to the left
	m.TurnLeft()
	if m.Orientation != "N" {
		t.Errorf("Expected orientation: N, Got: %s", m.Orientation)
	}

	// Test moving forward
	m.MoveForward()
	if m.X != 0 || m.Y != 1 {
		t.Errorf("Expected position: (0, 1), Got: (%d, %d)", m.X, m.Y)
	}

	// Test moving backward
	m.MoveBackward()
	if m.X != 0 || m.Y != 0 {
		t.Errorf("Expected position: (0, 0), Got: (%d, %d)", m.X, m.Y)
	}
}
