package main

import "fmt"

type Direction int

const (
	North Direction = iota
	East
	South
	West
)

type MaqeBot struct {
	CurrentPosition  *Position
	CurrentDirection Direction
}

func NewMaqeBot(currentPosition *Position, currentDirection Direction) IBotAction {
	return &MaqeBot{
		CurrentPosition:  currentPosition,
		CurrentDirection: currentDirection,
	}
}

func (b *MaqeBot) Turn(c Command) {
	switch c {
	case Left:
		b.CurrentDirection -= 1

		if b.CurrentDirection < North {
			b.CurrentDirection = West
		}
	case Right:
		b.CurrentDirection += 1

		if b.CurrentDirection > West {
			b.CurrentDirection = North
		}
	}
}

func (b *MaqeBot) MoveForward(step int) {
	switch b.CurrentDirection {
	case North:
		b.CurrentPosition.y += step
	case East:
		b.CurrentPosition.x += step
	case South:
		b.CurrentPosition.y -= step
	case West:
		b.CurrentPosition.x -= step
	}
}

func (b *MaqeBot) GetCurrentPosition() *Position {
	return b.CurrentPosition
}

func (b *MaqeBot) GetCurrentDirection() string {
	return getDirectionName(b.CurrentDirection)
}

func (b *MaqeBot) DisplayBotPosition() string {
	return fmt.Sprintf("X: %d Y: %d Direction: %s", b.CurrentPosition.x, b.CurrentPosition.y, getDirectionName(b.CurrentDirection))
}

func getDirectionName(d Direction) string {
	switch d {
	case North:
		return "North"
	case East:
		return "East"
	case South:
		return "South"
	case West:
		return "West"
	default:
		return ""
	}
}
