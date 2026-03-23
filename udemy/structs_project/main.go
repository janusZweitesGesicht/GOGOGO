package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"

	"example.com/note/note"
)

func main () {
	title, content := getNoteData()
	
	userNote, err := note.New(title, content)

	if err != nil {
		fmt.Println(err)
		return
	}
	userNote.Display()
	err = userNote.Save()
	if err != nil {
		panic(errors.New("Saving the note failed"))
	}
	fmt.Println("Saving The Note Succeeded")
}

func getNoteData () (string, string) {
	title := getUserInput("Notes Title:")
	content := getUserInput("Notes content:")
	return title, content
}


func getUserInput(prompt string) (string) {
	fmt.Printf("%v ",prompt)

	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		return ""
	}

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")


	return text
}