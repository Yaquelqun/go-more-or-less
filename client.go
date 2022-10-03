package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
)

func runClient(wg *sync.WaitGroup, channel chan GameState) {
	println("Waiting for the game to be initialized")
	game := <-channel
	fmt.Printf("Welcome to the More or Less game\n")
Loop:
	for {
		game.status = "Guessing"
		game.currentGuess, game.status = acceptNumber(game)
		if game.status != "Guessed" {
			continue
		}

		channel <- game
		game = <-channel
		switch game.status {
		case "Guess too high":
			fmt.Printf("%d is too high, try again\n", game.currentGuess)
		case "Guess too low":
			fmt.Printf("%d is too low, try again\n", game.currentGuess)
		case "Finished":
			fmt.Printf("The number was indeed %d ! Well Played !!!\n", game.currentGuess)
			break Loop
		default:
			println("Something went wrong, try again")
		}
	}

	println("Thank you, come again")
	close(channel)
	wg.Done()
}

func acceptNumber(game GameState) (int, string) {
	fmt.Printf("Please input a number between %d and %d\n", game.minimumValue, game.maximumValue)

	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("An error occured while reading input. Please try again", err)
		return -1, "error"
	}

	input = strings.TrimSuffix(input, "\n")
	guess, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("That was not a number. Please try again", err)
		return -1, "not a number"
	} else if guess < game.minimumValue || guess > game.maximumValue {
		fmt.Println("Number is out of bound. Please try again", err)
		return -1, "out of bound"
	}

	return guess, "Guessed"
}
