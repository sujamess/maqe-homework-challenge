package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

type Command string

const (
	Left  Command = "L"
	Right Command = "R"
	Walk  Command = "W"
)

type Position struct {
	x int
	y int
}

type IBotAction interface {
	Turn(c Command)
	MoveForward(step int)
	GetCurrentPosition() *Position
	GetCurrentDirection() string
	DisplayBotPosition() string
}

type Bot struct {
	steps  []string
	moving bool
	action IBotAction
}

func NewBot(botAction IBotAction) *Bot {
	return &Bot{
		moving: false,
		action: botAction,
	}
}

func (b *Bot) ExecuteCommand(commands string) error {
	formattedCommands := strings.ToUpper(commands)

	for i := 0; i < len(formattedCommands); i++ {
		command := string(formattedCommands[i])

		if Command(command) == Left || Command(command) == Right {
			err := b.moveForwardAndReset()
			if err != nil {
				return err
			}

			b.action.Turn(Command(formattedCommands[i]))
		} else if Command(command) == Walk {
			// For checking the case that the next command is still `W` like RW5W10LW3
			err := b.moveForwardAndReset()
			if err != nil {
				return err
			}
			// Set the moving to `true` after reset the state
			b.moving = true
		} else if unicode.IsNumber(rune(formattedCommands[i])) && b.moving {
			b.steps = append(b.steps, command)
		}

		/**
		Check if the last iteration still have some step,
		then we move forward
		*/
		if i == len(formattedCommands)-1 {
			err := b.moveForwardAndReset()
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (b *Bot) moveForwardAndReset() error {
	if len(b.steps) > 0 && b.moving {
		step, err := strconv.Atoi(strings.Join(b.steps, ""))
		if err != nil {
			return fmt.Errorf("could not parse the string into an integer: %v", err)
		}

		b.action.MoveForward(step)

		b.steps = []string{}
		b.moving = false
	}

	return nil
}
