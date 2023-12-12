package tasker

import (
	"fmt"
	"log"
	"regexp"
	"time"

	"github.com/h3th-IV/taskManager/database"
	"github.com/joho/godotenv"
)

func TaskManagerApp() {
	var (
		user           string
		pass           string
		userNameNExist string
		passwordNExist string
		newUser        string
		userNameExist  string
		passwordExist  string
		continum       string
	)
	usrName := "^[a-zA-Z0-9]{5,10}$"
	userCheker, err := regexp.Compile(usrName)
	if err != nil {
		log.Fatalf("%v", err)
	}
	userSecret := "^[a-zA-Z0-9!@#$%?/\\<>.,;:]{8,15}"
	passChecker, err := regexp.Compile(userSecret)
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

	fmt.Println("Do you have an account(Y/N): ")
	fmt.Scanf("%s", newUser)
	if newUser == "Y" || newUser == "y" {
		fmt.Println("Username(must be alphanumeric, 5-10 characters	): ")
		fmt.Scanf("%s", &user)
		fmt.Println("Password(must be at least 8 characters): ")
		fmt.Scanf("%s", &pass)
		//try to match username and password, if it matches specification
		if userCheker.MatchString(user) && passChecker.MatchString(pass) {
			userNameNExist = user
			passwordNExist = pass
			database.CreateUser(userNameNExist, passwordNExist)
			database.CreateTaskTable()
		}
	}
	fmt.Print("Username: ")
	fmt.Scanf("%s", &userNameExist)
	fmt.Print("Password: ")
	fmt.Scanf("%s", &passwordExist)
	auth, err := database.SelectDetails(userNameExist)
	if err != nil {
		log.Fatalf("Unable to retrieve user deatils: %v", err)
	}
	if passwordExist == auth {
		// task := NewTask()
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
				fmt.Print("enter R or r to remove task, or any other key to exit: ")
				fmt.Scanf("%s", &continum)
				if continum == "R" || continum == "r" {
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
