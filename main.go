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
		"3. Exit the program",
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
		fmt.Print("\nEnter your option here: ")
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
						musicComposer = strings.TrimSpace(musicComposer)
					}

					musicTitleInput := bufio.NewReader(os.Stdin)
					fmt.Print("Enter music title: ")
					musicTitle, titleError := musicTitleInput.ReadString('\n')
							
					if titleError != nil {
						fmt.Println("Error reading title input:", titleError)
						return
					} else {
						musicTitle = strings.TrimSpace(musicTitle)
					}

					musicLibrary[musicComposer] = musicTitle

				case 2:
					musicSearchInput := bufio.NewReader(os.Stdin)
					fmt.Print("Enter music title or composer here: ")
					musicSearch, searchError := musicSearchInput.ReadString('\n')
							
					if searchError != nil {
						fmt.Println("Error reading search input:", searchError)
						return
					} else {
						musicSearch = strings.TrimSpace(musicSearch)

						foundSearch := false // The easiest way to do comparision logic in a for-loop
						for composer, title := range musicLibrary {
							if musicSearch == composer || musicSearch == title {
								foundSearch = true
								fmt.Printf("Composer: %s\nTitle: %s\n", composer, title)
							}
						}

						if (!foundSearch) {
							fmt.Println("Music seaech result not found. Perhaps is it not yet added to the app?")
						}
					}

				case 3:
					fmt.Println("Goodbye...")
					return;

				default:
					fmt.Println("Invalid choice. Please try again!")
			}
		}
	}
}
