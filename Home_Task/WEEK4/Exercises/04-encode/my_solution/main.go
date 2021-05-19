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
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// ---------------------------------------------------------
// EXERCISE: Encode
//
//  Add a new command: "save". Encode the games to json, and
//  print it, then terminate the loop.
//
//  1. Create a new struct type with exported fields: ID, Name, Genre and Price.
//
//  2. Create a new slice using the new struct type.
//
//  3. Save the games into the new slice.
//
//  4. Encode the new slice.
//
//
// RESTRICTION
//  Do not export the fields of the game struct.
//
//
// EXPECTED OUTPUT
//  Inanc's game store has 3 games.
//
//    > list   : lists all the games
//    > id N   : queries a game by id
//    > save   : exports the data to json and quits
//    > quit   : quits
//
//  save
//
//  [
//          {
//                  "id": 1,
//                  "name": "god of war",
//                  "genre": "action adventure",
//                  "price": 50
//          },
//          {
//                  "id": 2,
//                  "name": "x-com 2",
//                  "genre": "strategy",
//                  "price": 40
//          },
//          {
//                  "id": 3,
//                  "name": "minecraft",
//                  "genre": "sandbox",
//                  "price": 20
//          }
//  ]
//
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

type ExportedGame struct {
	ID    int
	Name  string
	Genre string
	Price int
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
			item:  item{id: 2, name: "x-com 2", price: 30},
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

	exportedGames := make([]ExportedGame, len(games))
	for i, g := range games {
		var exportedGame ExportedGame
		exportedGame.ID = g.id
		exportedGame.Name = g.name
		exportedGame.Price = g.price
		exportedGame.Genre = g.genre
		exportedGames[i] = exportedGame
	}

	exit_flag := false
	scanner := bufio.NewScanner(os.Stdin)
	for !exit_flag {
		fmt.Print("\nWelcome to go-epam game center. Choose your command:\n")
		fmt.Print("\n> list   : lists all the games\n")
		fmt.Println("> id N   : queries a game by id")
		fmt.Println("> save   : exports the data to json and quits")
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
			case "save":
				encodedGames, err := json.Marshal(exportedGames)
				if err != nil {
					fmt.Println(err)
				} else {
					fmt.Println(string(encodedGames))
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
