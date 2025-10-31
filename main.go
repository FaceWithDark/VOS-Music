package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

func main() {
    // fmt.Println("TUI Music Library, made in Go!");

    musicComposerInput := bufio.NewReader(os.Stdin)
    fmt.Print("Enter music composer: ")
	musicComposer, composerError := musicComposerInput.ReadString('\n')
	if composerError != nil {
		fmt.Println("Error reading composer input:", composerError)

		// exit for current condition, not the program because that is how Go
		// works with early return codes.
		return
	} else {
		/*
         * Go (and languages support multi-declare variables) will have pretty
		 * funny errors when nesting methods so it's better to re-assign value
		 * to existen variable.
         */

		// Trim <Enter> input (Go treated as '\n')
        musicComposer = strings.TrimSpace(musicComposer)
    }

    musicTitleInput := bufio.NewReader(os.Stdin)
    fmt.Print("Enter music title: ")
	musicTitle, titleError := musicTitleInput.ReadString('\n')
    
	if titleError != nil {
		fmt.Println("Error reading title input:", titleError)

		// exit for current condition, not the program because that is how Go
		// works with early return codes.
		return
	} else {
		/*
         * Go (and languages support multi-declare variables) will have pretty
		 * funny errors when nesting methods so it's better to re-assign value
		 * to existen variable.
         */

		// Trim <Enter> input (Go treated as '\n')
        musicTitle = strings.TrimSpace(musicTitle)
    }

	var musicLibrary = make(map[string]string)
	musicLibrary[musicComposer] = musicTitle

	fmt.Println(musicLibrary)
}
