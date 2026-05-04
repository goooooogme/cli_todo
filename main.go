package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func enterInput() (string, error) {
	reader := bufio.NewReader(os.Stdin)

	input, err := reader.ReadString('\n')

	input = strings.TrimSpace(input)

	return input, err
}

func addTask() {
	textForAddTask()

	input, err := enterInput()

	if err != nil {
		log.Fatal(err)
	}

	file, err := os.OpenFile("tasks.txt", os.O_APPEND|os.O_WRONLY, 0644)

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	file.WriteString(input + "\n")
	fmt.Println("done")
}

func listTasks() {
	file, err := os.Open("tasks.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func textForMenu() {
	fmt.Println("1 - list tasks")
	fmt.Println("2 - add task")

	fmt.Println("select operation:")
}

func textForAddTask() {
	fmt.Println("input")
}

func main() {
	textForMenu()

	input, err := enterInput()

	if err != nil {
		log.Fatal(err)
	}

	if input == "1" {
		listTasks()
	} else if input == "2" {
		addTask()
	} else {
		fmt.Println("invalid command")
	}
}
