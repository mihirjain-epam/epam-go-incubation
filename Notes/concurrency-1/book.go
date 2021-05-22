package main

import "fmt"

type Book struct {
	ID            int
	Title         string
	Author        string
	YearPublished int
}

func (b Book) String() string {
	return fmt.Sprintf(
		"Title:\t\t%q\n"+
			"Author:\t\t%q\n"+
			"Published:\t%v\n",
		b.Title, b.Author, b.YearPublished)
}

var books = []Book{
	{
		ID:            1,
		Title:         "1 - Think and grow rich",
		Author:        "",
		YearPublished: 1937,
	},
	{
		ID:            2,
		Title:         "2 - Think and grow rich",
		Author:        "",
		YearPublished: 1937,
	},
	{
		ID:            3,
		Title:         "3 - Think and grow rich",
		Author:        "",
		YearPublished: 1937,
	},
	{
		ID:            4,
		Title:         "4 - Think and grow rich",
		Author:        "",
		YearPublished: 1937,
	},
	{
		ID:            5,
		Title:         "5 - Think and grow rich",
		Author:        "",
		YearPublished: 1937,
	},
	{
		ID:            6,
		Title:         "6 - Think and grow rich",
		Author:        "",
		YearPublished: 1937,
	},
	{
		ID:            7,
		Title:         "7 - Think and grow rich",
		Author:        "",
		YearPublished: 1937,
	},
	{
		ID:            8,
		Title:         "8 - Think and grow rich",
		Author:        "",
		YearPublished: 1937,
	},
	{
		ID:            9,
		Title:         "9 - Think and grow rich",
		Author:        "",
		YearPublished: 1937,
	},
	{
		ID:            10,
		Title:         "10 - Think and grow rich",
		Author:        "",
		YearPublished: 1937,
	},
}
