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

const (
	data = `
[
        {
                "id": 1,
                "name": "god of war",
                "genre": "action adventure",
                "price": 50
        },
        {
                "id": 2,
                "name": "x-com 2",
                "genre": "strategy",
                "price": 40
        },
        {
                "id": 3,
                "name": "minecraft",
                "genre": "sandbox",
                "price": 20
        }
]`

	commands = `
Welcome to go-epam game center. Choose your command:

	> list   : lists all the games
	> id N   : queries a game by id	
	> save   : exports the data to json and quits
	> quit   : quits

`
)

type item struct {
	id    int
	name  string
	price int
}

type game struct {
	item
	genre string
}

type jsonGame struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
	Genre string `json:"genre"`
}

/*
 * Input arguments -
 *                nil
 *
 * Return value - string -> string representation of item struct
 */
func (i item) String() string {
	return fmt.Sprintf("id:%d, name:%s, price:%d", i.id, i.name, i.price)
}

/*
 * Input arguments -
 *                nil
 *
 * Return value - string -> string representation of game struct
 */
func (g game) String() string {
	return fmt.Sprintf("#%d: %-15q %-20s $%d",
		g.id, g.name, "("+g.genre+")", g.price)
}

/*
 * Input arguments -
 *                games []game - list of games
 *
 * Return value - nil
 */
func list(games []game) {
	fmt.Print("\nList of games in my store is-\n\n")
	for _, g := range games {
		fmt.Println(g)
	}
	fmt.Println()
}

/*
 * Input arguments -
 *                nil
 *
 * Return value - bool -> flag value to quit
 */
func quit() bool {
	fmt.Println("bye!")
	return true
}

/*
 * Input arguments -
 *              nil
 *
 * Return value -
 				[]game -> slice of games
				error -> error occured in decoding
*/
func decodeGames() ([]game, error) {
	var decodedGames []jsonGame
	err := json.Unmarshal([]byte(data), &decodedGames)
	if err != nil {
		fmt.Println("Decoding failed due to error:", err)
		return nil, err
	}
	games := make([]game, len(decodedGames))
	for i, dg := range decodedGames {
		games[i] = game{item{dg.ID, dg.Name, dg.Price}, dg.Genre}
	}
	return games, err
}

/*
 * Input arguments -
 *                games []game ->  slice of games
 *
 * Return value - nil
 */
func encodeGames(games []game) {
	exportedGames := make([]jsonGame, len(games))
	for i, g := range games {
		exportedGames[i] = jsonGame{g.id, g.name, g.price, g.genre}
	}
	encodedGames, err := json.Marshal(exportedGames)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(encodedGames))
	}
}

/*
 * Input arguments -
 *                gamesMap map[int]game - map with game id's as key and game as value
 *                commandArgs []string  - slice of strings containing commands and their argument
 *
 * Return value - nil
 */
func findById(gamesMap map[int]game, commandArgs []string) {
	gameId, err := strconv.Atoi(commandArgs[1])
	if err != nil {
		fmt.Print("wrong id\n")
	} else if val, ok := gamesMap[gameId]; ok {
		fmt.Println(val)
	} else {
		fmt.Print("sorry. i don't have the game\n")
	}
}

/*
 * Input arguments -
 *                games game[] - list of games
 *
 * Return value - map[int]game -> map with game id's as key and game as value
 */
func mapGamesById(games []game) map[int]game {
	gamesMap := make(map[int]game, len(games))
	for _, g := range games {
		gamesMap[g.id] = g
	}
	return gamesMap
}

func main() {
	games, err := decodeGames()
	if err != nil {
		return
	}

	exitFlag := false
	scanner := bufio.NewScanner(os.Stdin)
	for !exitFlag {
		fmt.Print(commands)
		scanner.Scan()
		command := strings.Trim(scanner.Text(), " ")
		commandArgs := strings.Fields(command)
		if len(command) != 0 {
			switch commandArgs[0] {
			case "list":
				list(games)
			case "quit":
				exitFlag = quit()
			case "id":
				gamesMap := mapGamesById(games)
				findById(gamesMap, commandArgs)
			case "save":
				encodeGames(games)
				exitFlag = quit()
			default:
				fmt.Print("\nPlease select from the provided choice:\n")
				continue
			}
		} else {
			break
		}
	}

}
