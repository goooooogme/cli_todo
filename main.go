package main

import (
	"bufio"
	"cli_todo/text"
	"fmt"
	"log"
	"os"
	"strings"
)

func enterInput() (string, error) {
	reader := bufio.NewReader(os.Stdin)

	input, err := reader.ReadString('\n')

	if err != nil {
		return "", err
	}

	input = strings.TrimSpace(input)

	return input, err
}

func addTask() {
	text.TextForAddTask()

	input, err := enterInput()

	if err != nil {
		log.Fatal(err)
	}

	file, err := os.OpenFile("tasks.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	file.WriteString(input + "\n")
	text.TextForAddTaskDone()
}

func listTasks() {
	file, err := os.Open("tasks.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func main() {
	for {
		text.TextForMenu()

		input, err := enterInput()

		if err != nil {
			log.Fatal(err)
		}

		switch input {
		case "1":
			listTasks()
		case "2":
			addTask()
		default:
			return
		}
	}

}
