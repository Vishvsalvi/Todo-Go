package main

import (
	"fmt"
)

var storage[] string 

func addTask(){
	
	var task string
	fmt.Printf("Enter your task: ")
	fmt.Scan(&task)
	if task == ""{
		for {
			
			fmt.Println("Please enter a task: ")
			fmt.Scan(&task)

			if task != "" {
				storage = append(storage, task)
				fmt.Println("Task added successfully")
				break
			}
			
		}
		} else {
			storage = append(storage, task)
			fmt.Println("Task added successfully")
	}

}

func displayTask(){

	if(len(storage) == 0) {
		fmt.Printf("Your task list is empty\n")
		return;
	} else {
	
		for index, task := range storage{
			fmt.Printf("\n%d. %s\n", index + 1, task)
		}
	}
}

func deleteTask(){
	displayTask()
	var deleteKey int
	fmt.Printf("Enter the serial no of the task you want to delete\n")
	fmt.Scan(&deleteKey)

	if deleteKey < 1 || deleteKey > len(storage) {
		fmt.Printf("Invalid serial no\n")
		} else {
			storage = append(storage[:deleteKey - 1],storage[deleteKey:]... )
			fmt.Printf("Task deleted successfully\n")
			return
	}
	
}

func main(){
	for {
		fmt.Printf("Hit 1 to add task\nHit 2 to see the list\nHit 3 to delete a specific task\n")
		var command string
		fmt.Scan(&command)

		if command == "1" {
			addTask();
		} else if command == "2" {
			displayTask()
		} else if command == "3" {
			deleteTask()
		} else {
			fmt.Printf("Invalid command!! Try again\n")
		}
	}
}