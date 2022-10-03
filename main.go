package main

import (
	"sync"
)

var wg = sync.WaitGroup{}
var channel = make(chan GameState)

func main() {
	wg.Add(2)
	go runClient(&wg, channel)

	go runServer(&wg, channel)
	wg.Wait()
}
