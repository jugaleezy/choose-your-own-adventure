package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"

	cyoa "github.com/jugaleezy/choose-your-own-adventure"
)

func main() {
	// port := flag.Int("port", 3000, "port to start the CYOA web app on")
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

	currentStory := "intro"
	showChapter(story[currentStory])

	optionChannel := make(chan string)

	for {
		go func() {
			var option string
			fmt.Scanf("%s\n", &option)
			optionChannel <- strings.TrimSpace(option)
		}()

		option, ok := <-optionChannel
		if !ok {
			fmt.Println("Something went wrong")
			break
		}

		if option == "q" {
			os.Exit(0)
		}

		num, err := strconv.Atoi(option)
		if err != nil {
			fmt.Println("\nYou must choose from given options")
			continue
		}

		if num-1 >= 0 && num-1 < len(story[currentStory].Options) {
			currentStory = story[currentStory].Options[num-1].Chapter
			showChapter(story[currentStory])
			continue
		} else {
			fmt.Println("\nYou must choose from given options")
			continue
		}
	}
}

func showChapter(c cyoa.Chapter) {
	clearConsole()
	fmt.Println("\n~~~~~~~~~~~~ CHOOSE YOUR OWN ADVENTURE ~~~~~~~~~~~~\n\n")
	fmt.Println("\n" + c.Title + "\n")
	for _, paragraph := range c.Paragraph {
		fmt.Println("\n" + paragraph)
	}
	fmt.Println("\n\n")
	for index, nextChapter := range c.Options {
		fmt.Printf("%d. %s\n", index+1, nextChapter.Text)
	}
	fmt.Println("\n\nPress \"q\" to exit the adventure game")
	return
}

func clearConsole() {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}

	cmd.Stdout = os.Stdout
	cmd.Run()
}
