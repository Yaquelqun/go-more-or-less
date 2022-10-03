package main

import (
	"math/rand"
	"sync"
)

var targetNumber int

func runServer(wg *sync.WaitGroup, channel chan GameState) {
	game := GameState{status: "Created"}
	initializeGame(&game)
	channel <- game
Loop:
	for {
		game = <-channel
		switch {
		case game.currentGuess < targetNumber:
			game.status = "Guess too low"
			channel <- game
		case game.currentGuess > targetNumber:
			game.status = "Guess too high"
			channel <- game
		default:
			game.status = "Finished"
			channel <- game
			break Loop
		}
	}
	wg.Done()
}

func initializeGame(game *GameState) {
	game.minimumValue = 0
	game.maximumValue = 100
	targetNumber = rand.Intn(game.maximumValue)
	game.status = "Initialized"
}
