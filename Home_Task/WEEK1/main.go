package main

import (
	"math/rand"
	"time"

	"github.com/kyokomi/emoji/v2"
)

func main() {
	keys_array := []string{
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
	rand.Seed(time.Now().UnixNano())
	emoji.Println(emoji.RevCodeMap()[keys_array[rand.Intn(len(keys_array))]][0])
	emoji.Println(emoji.RevCodeMap()[keys_array[rand.Intn(len(keys_array))]][0])
	emoji.Println(emoji.RevCodeMap()[keys_array[rand.Intn(len(keys_array))]][0])
}
