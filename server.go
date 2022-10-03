package main

import "sync"

func runServer(wg *sync.WaitGroup, channel chan GameState) {
	game := GameState{status: "Created"}
	initializeGame(&game)
	channel <- game
	wg.Done()
}

func initializeGame(game *GameState) {
	game.minimumValue = 0
	game.maximumValue = 100
	game.status = "Initialized"
}
