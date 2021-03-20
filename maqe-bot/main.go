package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	walkingCode := os.Args[1:]

	if len(walkingCode) > 1 {
		log.Fatal("Please pass only one command line argument")
	} else if len(walkingCode) == 0 {
		log.Fatal("Please input the command")
	}

	maqeBot := NewMaqeBot(&Position{0, 0}, North)
	bot := NewBot(maqeBot)

	err := bot.ExecuteCommand(walkingCode[0])
	if err != nil {
		log.Fatalf("An error occurred while trying to parse a string into an integer: %v", err)
	}

	fmt.Println(bot.action.DisplayBotPosition())
}
