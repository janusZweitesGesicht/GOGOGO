package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"

	"example.com/interfacec/note"
	"example.com/interfacec/todo"
)


type saver interface {
	Save() error
}

// type displayer interface {
// 	Display()
// }


type outputtable interface {
	saver
	Display()
}

// type outputtable interface {
// 	Save() error
// 	Display()
// }

func main () {
	title, content := getNoteData()
	todoText := getTodoData()
	todo, err := todo.New(todoText)
	
	if err != nil {
		fmt.Println(err)
		return
	}

	userNote, err := note.New(title, content)

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("________________________________")

	err = outputData(todo)
	if err != nil {
		fmt.Println(err)
		return
	}
	//_________________
	err = outputData(userNote)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func getNoteData () (string, string) {
	title := getUserInput("Notes Title:")
	content := getUserInput("Notes content:")
	return title, content
}

func getTodoData() (string) {
	text := getUserInput("Todo Text:")
	return text
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

func saveData (data saver) error {
	err := data.Save()
	if err != nil {
		fmt.Println(errors.New("Saving the note failed"))
		return err
	}
	fmt.Println("Saving The Note Succeeded")
	return nil
}

func outputData(data outputtable) error{
	data.Display()
	return saveData(data)
	
}


// func printSomething(value any){
// fmt.Println(value)
// }

// func printSomething(value interface{}){
// aInt, alsInt := a.(int)
// bInt, bIsInt := b.(int)
// if alsInt && bIsInt {
// return aInt + bInt
// }
// fmt.Println(value)
// }


//generics
// func add[T int|float64|string](a, b T) T {
// return a + b 
// }