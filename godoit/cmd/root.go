package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/juanjuanzero/BuildingAnAwesomeCliApp/godoit/domain"
	"github.com/spf13/cobra"
)

var (
	file    string = "todo.json"
	rootCmd        = &cobra.Command{
		Use:   "godoit",
		Short: "this is my first cli-app so please generous with your feedback!",
		Long:  "a todo list application build for the love of learning",
	}
	todoList domain.ToDoList = domain.ToDoList{}
)

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	// TODO: read file
	fi, err := os.Open(file)
	if err != nil {
		panic(err)
	}
	// defer a Close so we always close the file at the end of this
	defer func() {
		if err := fi.Close(); err != nil {
			panic(err)
		}
	}()

	// you can read the full file, but wanted to challenge myself to read 1 byte at a time.
	var fullfile []byte
	for {
		buf := make([]byte, 1)
		_, err := fi.Read(buf)
		if err != nil {
			if err == io.EOF {
				// Read the end, breaking out
				break
			}
		}
		fullfile = append(fullfile, buf...)
	}

	// once you have a byte array you can use the encoding/json lib to read the file
	err = json.Unmarshal(fullfile, &todoList)
	if err != nil {
		fmt.Println("marshall error:", err)
	}

	// managing in-memory at the moment. Yes, its not scalable but its a to-do list :shrug:
	fmt.Printf("Successfully initialized with %v tasks from %v\n", len(todoList.ToDoItems), fi.Name())
}
