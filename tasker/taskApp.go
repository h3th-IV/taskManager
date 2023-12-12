package tasker

import (
	"fmt"
	"log"
	"regexp"
	"time"

	"github.com/h3th-IV/taskManager/database"
	"github.com/joho/godotenv"
)

// 09020440447
// macbobbychibuzor@gmail.com
func printMenu() {
	fmt.Println("\nTask List CLI")
	fmt.Println("||Select an option for Tasks||\n ----------------------------")
	fmt.Println(" [1]\t||   Create Task")
	fmt.Println(" [2]\t||   List Task")
	fmt.Println(" [3]\t||   Mark-Completed")
	fmt.Println(" [4]\t||   Remove-Task")
	fmt.Println(" [5]\t||   Exit")
}

func TaskManagerApp() {
	var (
		userName string
		option   int
		continum string
		UserList = &TaskManager{}
	)
	usrName := "^[a-zA-Z0-9]{5,10}$"
	userCheker, err := regexp.Compile(usrName)
	if err != nil {
		log.Fatalf("%v", err)
	}
	//load envrons
	err = godotenv.Load()
	if err != nil {
		log.Fatalf("error loading env variables: %v", err)
	}

	err = database.InitDB()
	if err != nil {
		log.Fatal(err)
	}

	defer database.CloseDB()
	fmt.Print("Enter Your Username(Username must be at least 5 characters): ")
	fmt.Scanf("%s", &userName)
	if userCheker.MatchString(userName) {
		// task := NewTask()
		fmt.Printf("Hello, %s\n", userName)
	start:
		printMenu()
		fmt.Print("select an option: ")
		fmt.Scanf("%d", &option)
		switch option {
		case 1:
			//Create Task
			for {
				fmt.Print("enter A or a to add task, or any other key to exit: ")
				fmt.Scanf("%s", &continum)
				if continum == "A" || continum == "a" {
					UserList.CreateTask()
				} else {
					goto start
				}
			}
		case 2:
			// List Task
			for {
				fmt.Print("enter L or l to List task, or any other key to exit: ")
				fmt.Scanf("%s", &continum)
				if continum == "L" || continum == "l" {
					UserList.ListTask()
				} else {
					goto start
				}
			}
		case 3:
			//Mark Completed
			for {
				fmt.Print("enter M or m to mark task as completed, or any other key to exit: ")
				fmt.Scanf("%s", &continum)
				if continum == "M" || continum == "m" {
					UserList.MarkComplete()
				} else {
					goto start
				}
			}
		case 4:
			//Remove Task
			for {
				fmt.Print("enter R or r to remove task, or any other key to exit: ")
				fmt.Scanf("%s", &continum)
				if continum == "R" || continum == "r" {
					UserList.RemoveTask()
				} else {
					goto start
				}
			}
		case 5:
			//SCRAM
			fmt.Println("Exiting the program...")
			time.Sleep(3 * time.Second)
			return
		default:
			fmt.Println("Invalid choice")
		}

		goto start
	} else {
		log.Printf("Authentication Error: %v", err)
		return
	}

}
