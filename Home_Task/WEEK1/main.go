package main

import (
	"math/rand"
	"time"
	"flag"

	"github.com/kyokomi/emoji/v2"
)

// Note: cannot use implicit connversion outside function like emoji_codes:= ...
var emoji_codes = []string{
	"\U0001f600",
	"\U0001f601",
	"\U0001f602",
	"\U0001f603",
	"\U0001f604",
	"\U0001f605",
	"\U0001f606",
	"\U0001f923",
	"\U0001f642",
	"\U0001f643",
	"\U0001f609",
	"\U0001f60A",
	"\U0001f607",
}
/*
	input args: 'n' int : number of emojis to Print
	return nil

	This function prints 'n' smiling emojis at random from the emoji_codes array
*/
func PrintRandomEmojis(n int){
	rand.Seed(time.Now().UnixNano())
	for ;n > 0;n--{
		rand_id := rand.Intn(len(emoji_codes))
		selected_emoji_code := emoji_codes[rand_id]
		selected_emoji := emoji.RevCodeMap()[selected_emoji_code][0]
		emoji.Println(selected_emoji)
	}
}

/*
	main func receives command line arguments using flag package
*/
func main() {
	num_of_emojis := flag.Int("num_of_emojis", 3, "Number of Emojis to print to console")
	flag.Parse()
	PrintRandomEmojis(*num_of_emojis)
}
