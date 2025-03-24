package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/jugaleezy/go-cyoa"
)

func main() {
	port := flag.Int("port", 3000, "port to start the CYOA web app on")
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

	// fmt.Printf("%+v", story)

	//handler
	h := cyoa.NewHandler(story, nil)
	fmt.Printf("Starting the server on port: %d\n", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), h))
}
