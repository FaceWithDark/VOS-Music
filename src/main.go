package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

func main() {
    // fmt.Println("TUI Music Library, made in Go!");

    musicTitle := bufio.NewReader(os.Stdin)
    fmt.Print("Enter music title that you want to play: ")

    userInput, error := musicTitle.ReadString('\n')
    if error != nil {
        fmt.Println("Error reading input", error)
        return // exit with code 1 for current conditon, not the program
    } else {
        /*
         * Go (and languages support multi-declare variables) will have pretty funny
         * errors when nesting methods so it's better to re-assign value to existen
         * variable.
         */
        userInput = strings.TrimSpace(userInput) // Trim "Enter" input (we treated as '\n')
        fmt.Println("Music Title:", userInput)
    }
}

// vim: ts=4 sw=4 et

