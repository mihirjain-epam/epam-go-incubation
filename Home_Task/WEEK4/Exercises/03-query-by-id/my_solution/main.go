// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// ---------------------------------------------------------
// EXERCISE: Query By Id
//
//  Add a new command: "id". So the users can query the games
//  by id.
//
//  1. Before the loop, index the games by id (use a map).
//
//  2. Add the "id" command.
//     When a user types: id 2
//     It should print only the game with id: 2.
//
//  3. Handle the errors:
//
//     id
//     wrong id
//
//     id HEY
//     wrong id
//
//     id 10
//     sorry. i don't have the game
//
//     id 1
//     #1: "god of war" (action adventure) $50
//
//     id 2
//     #2: "x-com 2" (strategy) $40
//
//
// EXPECTED OUTPUT
//  Please also run the solution and try the program with
//  list, quit, and id commands to see it in action.
// ---------------------------------------------------------

type item struct {
	id    int
	name  string
	price int
}

type game struct {
	item
	genre string
}

func (i item) String() string {
	return fmt.Sprintf("id:%d, name:%s, price:%d", i.id, i.name, i.price)
}

func (g game) String() string {
	return fmt.Sprintf("#%d: %-15q %-20s $%d",
		g.id, g.name, "("+g.genre+")", g.price)
}

func list(games []game) {
	fmt.Print("\nList of games in my store is-\n\n")
	for _, g := range games {
		fmt.Println(g)
	}
	fmt.Println()
}

func exit() bool {
	return true
}

func main() {
	games := []game{
		{
			item:  item{id: 1, name: "god of war", price: 50},
			genre: "action adventure",
		},
		{
			item:  item{id: 2, name: "x-com 2 ", price: 30},
			genre: "strategy",
		},
		{
			item:  item{id: 3, name: "minecraft", price: 20},
			genre: "sandbox",
		},
	}

	gamesMap := make(map[int]game, len(games))
	for _, g := range games {
		gamesMap[g.id] = g
	}

	exit_flag := false
	scanner := bufio.NewScanner(os.Stdin)
	for !exit_flag {
		fmt.Print("\nWelcome to go-epam game center. Choose your command:\n")
		fmt.Print("\n> list   : lists all the games\n")
		fmt.Println("> id N   : queries a game by id")
		fmt.Print("> quit   : quits\n\n")
		scanner.Scan()
		command := strings.Trim(scanner.Text(), " ")
		commandArgs := strings.Fields(command)
		if len(command) != 0 {
			switch commandArgs[0] {
			case "list":
				list(games)
			case "quit":
				fmt.Println("bye!")
				exit_flag = exit()
			case "id":
				gameId, err := strconv.Atoi(commandArgs[1])
				if err != nil {
					fmt.Print("wrong id\n")
				} else if val, ok := gamesMap[gameId]; ok {
					fmt.Println(val)
				} else {
					fmt.Print("sorry. i don't have the game\n")
				}
			default:
				fmt.Print("\nPlease select from the provided choice:\n")
				continue
			}
		} else {
			break
		}
	}

}
