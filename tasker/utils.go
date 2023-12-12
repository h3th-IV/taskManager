package tasker

import "fmt"

var (
	option   int
	UserList = &TaskManager{}
)

func printMenu() int {
	fmt.Println("\nTask List CLI")
	fmt.Println("||Select an option for Tasks||\n ----------------------------")
	fmt.Println(" [1]\t||   Create Task")
	fmt.Println(" [2]\t||   List Task")
	fmt.Println(" [3]\t||   Mark-Completed")
	fmt.Println(" [4]\t||   Remove-Task")
	fmt.Println(" [5]\t||   Exit")

	fmt.Print("Select an option: ")
	fmt.Scanf("%d", &option)
	return option
}

// Helper function to handle task addition
func handleTaskAddition() string {
	var continum string
	for {
		fmt.Print("Enter A or a to add task, or any other key to exit: ")
		fmt.Scanf("%s", &continum)
		if continum == "A" || continum == "a" {
			UserList.CreateTask()
		} else {
			return continum
		}
	}
}

// Helper function to handle task listig
func Lister() string {
	var continum string
	for {
		fmt.Print("Enter L or l to list tasks, or any other key to exit: ")
		fmt.Scanf("%s", &continum)
		if continum == "L" || continum == "l" {
			UserList.ListTask()
		} else {
			return continum
		}
	}
}

// Helper function to handle task marking
func Marka() string {
	var continum string
	for {
		fmt.Print("Enter M or m to mark task as completed, or any other key to exit: ")
		fmt.Scanf("%s", &continum)
		if continum == "M" || continum == "m" {
			UserList.MarkComplete()
		} else {
			return continum
		}
	}
}

// Helper function to handle task removal
func Remover() string {
	var continum string
	for {
		fmt.Print("Enter R or r to remove task, or any other key to exit: ")
		fmt.Scanf("%s", &continum)
		if continum == "R" || continum == "r" {
			UserList.RemoveTask()
		} else {
			return continum
		}
	}
}
