package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/jugaleezy/go-cyoa"
)

func main() {
	filename := flag.String("file", "gopher.json", "the JSON file with CYOA story")
	flag.Parse()

	fmt.Printf("using story in file %v\n", *filename)

	f, err := os.Open(*filename)
	if err != nil {
		panic(err)
	}

	story, err := cyoa.JsonStory(f)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v", story)
}
