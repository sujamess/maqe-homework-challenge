package main

import (
	"fmt"
	"testing"
)

func TestMaqeBotTurn(t *testing.T) {
	testCases := []struct {
		command                Command
		currentDirection       Direction
		expectCurrentDirection Direction
	}{
		{command: Left, currentDirection: North, expectCurrentDirection: West},
		{command: Right, currentDirection: North, expectCurrentDirection: East},
		{command: Left, currentDirection: East, expectCurrentDirection: North},
		{command: Right, currentDirection: East, expectCurrentDirection: South},
		{command: Left, currentDirection: South, expectCurrentDirection: East},
		{command: Right, currentDirection: South, expectCurrentDirection: West},
		{command: Left, currentDirection: West, expectCurrentDirection: South},
		{command: Right, currentDirection: West, expectCurrentDirection: North},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("Current direction: %s, Turn: %s", getDirectionName(tc.currentDirection), tc.command), func(*testing.T) {
			maqeBot := NewMaqeBot(&Position{0, 0}, tc.currentDirection)

			maqeBot.Turn(tc.command)

			if maqeBot.GetCurrentDirection() != getDirectionName(tc.expectCurrentDirection) {
				t.Errorf("got %q, want %q", maqeBot.GetCurrentDirection(), getDirectionName(tc.expectCurrentDirection))
			}
		})
	}
}

func TestMaqeBotMoveForward(t *testing.T) {
	testCases := []struct {
		step             int
		currentDirection Direction
		currentPosition  *Position
		expectPosition   *Position
	}{
		// Step 1 with (0, 0)
		{
			step:             1,
			currentDirection: North,
			currentPosition:  &Position{0, 0},
			expectPosition:   &Position{0, 1},
		},
		{
			step:             1,
			currentDirection: East,
			currentPosition:  &Position{0, 0},
			expectPosition:   &Position{1, 0},
		},
		{
			step:             1,
			currentDirection: South,
			currentPosition:  &Position{0, 0},
			expectPosition:   &Position{0, -1},
		},
		{
			step:             1,
			currentDirection: West,
			currentPosition:  &Position{0, 0},
			expectPosition:   &Position{-1, 0},
		},
		// Step 0 with (0, 0)
		{
			step:             0,
			currentDirection: North,
			currentPosition:  &Position{0, 0},
			expectPosition:   &Position{0, 0},
		},
		{
			step:             0,
			currentDirection: East,
			currentPosition:  &Position{0, 0},
			expectPosition:   &Position{0, 0},
		},
		{
			step:             0,
			currentDirection: South,
			currentPosition:  &Position{0, 0},
			expectPosition:   &Position{0, 0},
		},
		{
			step:             0,
			currentDirection: West,
			currentPosition:  &Position{0, 0},
			expectPosition:   &Position{0, 0},
		},
		// Step 5 with (0, 0)
		{
			step:             5,
			currentDirection: North,
			currentPosition:  &Position{0, 0},
			expectPosition:   &Position{0, 5},
		},
		{
			step:             5,
			currentDirection: East,
			currentPosition:  &Position{0, 0},
			expectPosition:   &Position{5, 0},
		},
		{
			step:             5,
			currentDirection: South,
			currentPosition:  &Position{0, 0},
			expectPosition:   &Position{0, -5},
		},
		{
			step:             5,
			currentDirection: West,
			currentPosition:  &Position{0, 0},
			expectPosition:   &Position{-5, 0},
		},
		// Step 5 with (1, 0)
		{
			step:             5,
			currentDirection: North,
			currentPosition:  &Position{1, 0},
			expectPosition:   &Position{1, 5},
		},
		{
			step:             5,
			currentDirection: East,
			currentPosition:  &Position{1, 0},
			expectPosition:   &Position{6, 0},
		},
		{
			step:             5,
			currentDirection: South,
			currentPosition:  &Position{1, 0},
			expectPosition:   &Position{1, -5},
		},
		{
			step:             5,
			currentDirection: West,
			currentPosition:  &Position{1, 0},
			expectPosition:   &Position{-4, 0},
		},
		// Step 5 with (0, 1)
		{
			step:             5,
			currentDirection: North,
			currentPosition:  &Position{0, 1},
			expectPosition:   &Position{0, 6},
		},
		{
			step:             5,
			currentDirection: East,
			currentPosition:  &Position{0, 1},
			expectPosition:   &Position{5, 1},
		},
		{
			step:             5,
			currentDirection: South,
			currentPosition:  &Position{0, 1},
			expectPosition:   &Position{0, -4},
		},
		{
			step:             5,
			currentDirection: West,
			currentPosition:  &Position{0, 1},
			expectPosition:   &Position{-5, 1},
		},
		// Step 5 with (1, 1)
		{
			step:             5,
			currentDirection: North,
			currentPosition:  &Position{1, 1},
			expectPosition:   &Position{1, 6},
		},
		{
			step:             5,
			currentDirection: East,
			currentPosition:  &Position{1, 1},
			expectPosition:   &Position{6, 1},
		},
		{
			step:             5,
			currentDirection: South,
			currentPosition:  &Position{1, 1},
			expectPosition:   &Position{1, -4},
		},
		{
			step:             5,
			currentDirection: West,
			currentPosition:  &Position{1, 1},
			expectPosition:   &Position{-4, 1},
		},
		// Step 5 with (-1, 0)
		{
			step:             5,
			currentDirection: North,
			currentPosition:  &Position{-1, 0},
			expectPosition:   &Position{-1, 5},
		},
		{
			step:             5,
			currentDirection: East,
			currentPosition:  &Position{-1, 0},
			expectPosition:   &Position{4, 0},
		},
		{
			step:             5,
			currentDirection: South,
			currentPosition:  &Position{-1, 0},
			expectPosition:   &Position{-1, -5},
		},
		{
			step:             5,
			currentDirection: West,
			currentPosition:  &Position{-1, 0},
			expectPosition:   &Position{-6, 0},
		},
		// Step 5 with (0, -1)
		{
			step:             5,
			currentDirection: North,
			currentPosition:  &Position{0, -1},
			expectPosition:   &Position{0, 4},
		},
		{
			step:             5,
			currentDirection: East,
			currentPosition:  &Position{0, -1},
			expectPosition:   &Position{5, -1},
		},
		{
			step:             5,
			currentDirection: South,
			currentPosition:  &Position{0, -1},
			expectPosition:   &Position{0, -6},
		},
		{
			step:             5,
			currentDirection: West,
			currentPosition:  &Position{0, -1},
			expectPosition:   &Position{-5, -1},
		},
		// Step 5 with (-1, -1)
		{
			step:             5,
			currentDirection: North,
			currentPosition:  &Position{-1, -1},
			expectPosition:   &Position{-1, 4},
		},
		{
			step:             5,
			currentDirection: East,
			currentPosition:  &Position{-1, -1},
			expectPosition:   &Position{4, -1},
		},
		{
			step:             5,
			currentDirection: South,
			currentPosition:  &Position{-1, -1},
			expectPosition:   &Position{-1, -6},
		},
		{
			step:             5,
			currentDirection: West,
			currentPosition:  &Position{-1, -1},
			expectPosition:   &Position{-6, -1},
		},
		// Step 5 with (-1, 1)
		{
			step:             5,
			currentDirection: North,
			currentPosition:  &Position{-1, 1},
			expectPosition:   &Position{-1, 6},
		},
		{
			step:             5,
			currentDirection: East,
			currentPosition:  &Position{-1, 1},
			expectPosition:   &Position{4, 1},
		},
		{
			step:             5,
			currentDirection: South,
			currentPosition:  &Position{-1, 1},
			expectPosition:   &Position{-1, -4},
		},
		{
			step:             5,
			currentDirection: West,
			currentPosition:  &Position{-1, 1},
			expectPosition:   &Position{-6, 1},
		},
		// Step 5 with (1, -1)
		{
			step:             5,
			currentDirection: North,
			currentPosition:  &Position{1, -1},
			expectPosition:   &Position{1, 4},
		},
		{
			step:             5,
			currentDirection: East,
			currentPosition:  &Position{1, -1},
			expectPosition:   &Position{6, -1},
		},
		{
			step:             5,
			currentDirection: South,
			currentPosition:  &Position{1, -1},
			expectPosition:   &Position{1, -6},
		},
		{
			step:             5,
			currentDirection: West,
			currentPosition:  &Position{1, -1},
			expectPosition:   &Position{-4, -1},
		},
	}

	for _, tc := range testCases {
		t.Run(
			fmt.Sprintf(
				`Current direction: %s, current position: (%d, %d), step: %d`,
				getDirectionName(tc.currentDirection),
				tc.currentPosition.x,
				tc.currentPosition.y,
				tc.step,
			),
			func(*testing.T) {
				maqeBot := NewMaqeBot(tc.currentPosition, tc.currentDirection)

				maqeBot.MoveForward(tc.step)

				if maqeBot.GetCurrentPosition().x != tc.expectPosition.x || maqeBot.GetCurrentPosition().y != tc.expectPosition.y {
					t.Errorf("got (%d, %d), want (%d, %d)", maqeBot.GetCurrentPosition().x, maqeBot.GetCurrentPosition().y, tc.expectPosition.x, tc.expectPosition.y)
				}
			},
		)
	}
}

func TestGetDirectionName(t *testing.T) {
	testCases := []struct {
		direction Direction
		expected  string
	}{
		{direction: North, expected: "North"},
		{direction: East, expected: "East"},
		{direction: South, expected: "South"},
		{direction: West, expected: "West"},
		{direction: Direction(-1), expected: ""},
		{direction: Direction(4), expected: ""},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("Getting the readable direction of %v", tc.direction), func(*testing.T) {
			result := getDirectionName(tc.direction)

			if result != tc.expected {
				t.Errorf("got %q, want %q", result, tc.expected)
			}
		})
	}
}

func TestDisplayBotPosition(t *testing.T) {
	testCases := []struct {
		currentPosition  *Position
		currentDirection Direction
		expect           string
	}{
		// North
		{currentPosition: &Position{0, 0}, currentDirection: North, expect: "X: 0 Y: 0 Direction: North"},
		{currentPosition: &Position{0, 1}, currentDirection: North, expect: "X: 0 Y: 1 Direction: North"},
		{currentPosition: &Position{1, 0}, currentDirection: North, expect: "X: 1 Y: 0 Direction: North"},
		{currentPosition: &Position{1, 1}, currentDirection: North, expect: "X: 1 Y: 1 Direction: North"},
		{currentPosition: &Position{0, -1}, currentDirection: North, expect: "X: 0 Y: -1 Direction: North"},
		{currentPosition: &Position{-1, 0}, currentDirection: North, expect: "X: -1 Y: 0 Direction: North"},
		{currentPosition: &Position{-1, -1}, currentDirection: North, expect: "X: -1 Y: -1 Direction: North"},
		{currentPosition: &Position{1, -1}, currentDirection: North, expect: "X: 1 Y: -1 Direction: North"},
		{currentPosition: &Position{-1, 1}, currentDirection: North, expect: "X: -1 Y: 1 Direction: North"},
		// East
		{currentPosition: &Position{0, 0}, currentDirection: East, expect: "X: 0 Y: 0 Direction: East"},
		{currentPosition: &Position{0, 1}, currentDirection: East, expect: "X: 0 Y: 1 Direction: East"},
		{currentPosition: &Position{1, 0}, currentDirection: East, expect: "X: 1 Y: 0 Direction: East"},
		{currentPosition: &Position{1, 1}, currentDirection: East, expect: "X: 1 Y: 1 Direction: East"},
		{currentPosition: &Position{0, -1}, currentDirection: East, expect: "X: 0 Y: -1 Direction: East"},
		{currentPosition: &Position{-1, 0}, currentDirection: East, expect: "X: -1 Y: 0 Direction: East"},
		{currentPosition: &Position{-1, -1}, currentDirection: East, expect: "X: -1 Y: -1 Direction: East"},
		{currentPosition: &Position{1, -1}, currentDirection: East, expect: "X: 1 Y: -1 Direction: East"},
		{currentPosition: &Position{-1, 1}, currentDirection: East, expect: "X: -1 Y: 1 Direction: East"},
		// South
		{currentPosition: &Position{0, 0}, currentDirection: South, expect: "X: 0 Y: 0 Direction: South"},
		{currentPosition: &Position{0, 1}, currentDirection: South, expect: "X: 0 Y: 1 Direction: South"},
		{currentPosition: &Position{1, 0}, currentDirection: South, expect: "X: 1 Y: 0 Direction: South"},
		{currentPosition: &Position{1, 1}, currentDirection: South, expect: "X: 1 Y: 1 Direction: South"},
		{currentPosition: &Position{0, -1}, currentDirection: South, expect: "X: 0 Y: -1 Direction: South"},
		{currentPosition: &Position{-1, 0}, currentDirection: South, expect: "X: -1 Y: 0 Direction: South"},
		{currentPosition: &Position{-1, -1}, currentDirection: South, expect: "X: -1 Y: -1 Direction: South"},
		{currentPosition: &Position{1, -1}, currentDirection: South, expect: "X: 1 Y: -1 Direction: South"},
		{currentPosition: &Position{-1, 1}, currentDirection: South, expect: "X: -1 Y: 1 Direction: South"},
		// West
		{currentPosition: &Position{0, 0}, currentDirection: West, expect: "X: 0 Y: 0 Direction: West"},
		{currentPosition: &Position{0, 1}, currentDirection: West, expect: "X: 0 Y: 1 Direction: West"},
		{currentPosition: &Position{1, 0}, currentDirection: West, expect: "X: 1 Y: 0 Direction: West"},
		{currentPosition: &Position{1, 1}, currentDirection: West, expect: "X: 1 Y: 1 Direction: West"},
		{currentPosition: &Position{0, -1}, currentDirection: West, expect: "X: 0 Y: -1 Direction: West"},
		{currentPosition: &Position{-1, 0}, currentDirection: West, expect: "X: -1 Y: 0 Direction: West"},
		{currentPosition: &Position{-1, -1}, currentDirection: West, expect: "X: -1 Y: -1 Direction: West"},
		{currentPosition: &Position{1, -1}, currentDirection: West, expect: "X: 1 Y: -1 Direction: West"},
		{currentPosition: &Position{-1, 1}, currentDirection: West, expect: "X: -1 Y: 1 Direction: West"},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("Current position: (%d, %d), direction: %s", tc.currentPosition.x, tc.currentPosition.y, getDirectionName(tc.currentDirection)), func(*testing.T) {
			maqeBot := NewMaqeBot(tc.currentPosition, tc.currentDirection)

			if maqeBot.DisplayBotPosition() != tc.expect {
				t.Errorf("got %q, want %q", maqeBot.DisplayBotPosition(), tc.expect)
			}
		})
	}
}
