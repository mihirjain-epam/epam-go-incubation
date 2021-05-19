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
	"strings"
)

// ---------------------------------------------------------
// EXERCISE: List
//
//  Now, it's time to add an interface to your program using
//  the bufio.Scanner. So the users can list the games, or
//  search for the games by id.
//
//  1. Scan for the input in a loop (use bufio.Scanner)
//
//  2. Print the available commands.
//
//  3. Implement the quit command: Quits from the loop.
//
//  4. Implement the list command: Lists all the games.
//
//
// EXPECTED OUTPUT
//  Please run the solution and try the program with list and
//  quit commands.
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
	exit_flag := false
	scanner := bufio.NewScanner(os.Stdin)
	for !exit_flag {
		fmt.Print("\nWelcome to go-epam game center. Choose your command:\n")
		fmt.Print("\n> list   : lists all the games\n")
		fmt.Print("> quit   : quits\n\n")
		scanner.Scan()
		text := strings.Trim(scanner.Text(), " ")
		if len(text) != 0 {
			switch text {
			case "list":
				list(games)
			case "quit":
				fmt.Println("bye!")
				exit_flag = exit()
			default:
				fmt.Print("\nPlease select from the provided choice:\n")
				continue
			}
		} else {
			break
		}
	}

}
