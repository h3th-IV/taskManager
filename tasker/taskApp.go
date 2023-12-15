package tasker

import (
	"fmt"
	"log"
	"time"
)

// --TODO REMEMBER TO REMOVE user_ID

var (
	userNameExist string
	passwordExist string
)

func TaskManagerApp() {
	var (
		continum string
	)
	Register()
	auth := authDetails.Auth
	if passwordExist == auth {
		fmt.Printf("Hello, %s\n", userNameExist)
	start:
		option := printMenu()
		switch option {
		case 1:
			//Create Task

			for {
				continum = handleTaskAddition()
				if continum != "A" && continum != "a" {
					goto start
				}
			}
		case 2:
			// List Task
			for {
				continum = Lister()
				if continum != "L" && continum != "l" {
					goto start
				}
			}
		case 3:
			//Mark Completed
			for {
				continum = Marka()
				if continum != "M" && continum != "m" {
					goto start
				}
			}
		case 4:
			//Remove Task
			for {
				continum = Remover()
				if continum != "R" && continum != "r" {
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
		log.Println("Authentication Error")
		return
	}

}
