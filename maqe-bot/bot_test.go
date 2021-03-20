package main

import (
	"fmt"
	"testing"
)

func TestBotExecuteCommandWithMaqeBot(t *testing.T) {
	testCases := []struct {
		command         string
		isFakeError     bool
		expectPosition  *Position
		expectDirection string
	}{
		{command: "RW15RW1", isFakeError: false, expectPosition: &Position{15, -1}, expectDirection: "South"},
		{command: "W5RW5RW2RW1R", isFakeError: false, expectPosition: &Position{4, 3}, expectDirection: "North"},
		{command: "RRW11RLLW19RRW12LW1", isFakeError: false, expectPosition: &Position{7, -12}, expectDirection: "South"},
		{command: "LLW100W50RW200W10", isFakeError: false, expectPosition: &Position{-210, -150}, expectDirection: "West"},
		{command: "LLLLLW99RRRRRW88LLLRL", isFakeError: false, expectPosition: &Position{-99, 88}, expectDirection: "East"},
		{command: "W55555RW555555W444444W1", isFakeError: false, expectPosition: &Position{1000000, 55555}, expectDirection: "East"},
		{command: "W55555W555555W444444W1", isFakeError: false, expectPosition: &Position{0, 1055555}, expectDirection: "North"},
		{command: "W55555", isFakeError: false, expectPosition: &Position{0, 55555}, expectDirection: "North"},
		{command: "W55555RW555555W444444W1", isFakeError: true, expectPosition: &Position{0, 0}, expectDirection: "East"},
		{command: "W55555W555555W444444W1", isFakeError: true, expectPosition: &Position{0, 0}, expectDirection: "East"},
		{command: "W55555", isFakeError: true, expectPosition: &Position{0, 0}, expectDirection: "East"},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("Executing command: %s", tc.command), func(*testing.T) {
			bot := NewBot(NewMaqeBot(&Position{0, 0}, North))

			if tc.isFakeError {
				bot.steps = append(bot.steps, "fake_error_string")

				err := bot.ExecuteCommand(tc.command)
				if err == nil {
					t.Error("got nil on error, expect not nil")
				}
			} else {
				err := bot.ExecuteCommand(tc.command)
				if err != nil {
					t.Error("got not nil on error, expect nil")
				}

				if bot.action.GetCurrentPosition().x != tc.expectPosition.x ||
					bot.action.GetCurrentPosition().y != tc.expectPosition.y ||
					bot.action.GetCurrentDirection() != tc.expectDirection {
					t.Errorf(
						`got (%d, %d) %s, expect (%d, %d) %s`,
						bot.action.GetCurrentPosition().x,
						bot.action.GetCurrentPosition().y,
						bot.action.GetCurrentDirection(),
						tc.expectPosition.x,
						tc.expectPosition.y,
						tc.expectDirection,
					)
				}
			}
		})
	}
}

func BenchmarkBotExecuteCommandWithMaqeBot(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bot := NewBot(NewMaqeBot(&Position{0, 0}, North))

		bot.ExecuteCommand("RW15RW1")
	}
}
