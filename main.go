package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

func appIntro() {
	const greet = "Welcome to VOS Music. Please select an option below to continue\n"

	var options = []string{
		"1. Add new music",
		"2. Search for music",
		"3. Exit the program\n",
	}

	fmt.Println(greet)

	for _, option := range options {
		fmt.Println(option)
	}
}

func main() {
	appIntro()

	var musicLibrary = make(map[string]string)
	for {
		fmt.Print("Enter your option here: ")
		var userOption int 
		_ , userError := fmt.Scan(&userOption)
		if (userError != nil) {
			fmt.Println("Error reading user input:", userError)
		} else {
			switch userOption {
				case 1:
					musicComposerInput := bufio.NewReader(os.Stdin)
					fmt.Print("Enter music composer: ")
					musicComposer, composerError := musicComposerInput.ReadString('\n')
					if composerError != nil {
						fmt.Println("Error reading composer input:", composerError)
						return
					} else {
						/*
						 * Go (and languages support multi-declare variables) will
						 * have pretty funny errors when nesting methods so it's
						 * better to re-assign value to existent variable.
						 */

						// Trim <Enter> input (Go treated as '\n')
						musicComposer = strings.TrimSpace(musicComposer)
					}

					musicTitleInput := bufio.NewReader(os.Stdin)
					fmt.Print("Enter music title: ")
					musicTitle, titleError := musicTitleInput.ReadString('\n')
							
					if titleError != nil {
						fmt.Println("Error reading title input:", titleError)
						return
					} else {
						/*
						 * Go (and languages support multi-declare variables) will
						 * have pretty funny errors when nesting methods so it's
						 * better to re-assign value to existent variable.
						 */

						// Trim <Enter> input (Go treated as '\n')
						musicTitle = strings.TrimSpace(musicTitle)
					}

					musicLibrary[musicComposer] = musicTitle
					fmt.Println(musicLibrary)

				case 2:
					fmt.Println("Coming soon!")

				case 3:
					fmt.Println("Goodbye...")
					return;

				default:
					fmt.Println("Please try again!")
			}
		}
	}
}
