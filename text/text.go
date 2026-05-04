package text

import "fmt"

func TextForMenu() {
	fmt.Println("1 - list tasks")
	fmt.Println("2 - add task")

	fmt.Println("select operation:")
}

func TextForAddTask() {
	fmt.Println("input task")
}

func TextForAddTaskDone() {
	fmt.Println("done")
}
