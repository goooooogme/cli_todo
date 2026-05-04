package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	fmt.Println("1 - list tasks")
	fmt.Println("2 - add task")

	fmt.Println("select operation:")

	reader := bufio.NewReader(os.Stdin)

	input, err := reader.ReadString('\n')

	input = strings.TrimSpace(input)

	if err != nil {
		log.Fatal(err)
	}

	if input == "1" {
		fmt.Println("list tasks")
	} else if input == "2" {
		fmt.Println("add task")
	} else {
		fmt.Println("invalid command")
	}
}
