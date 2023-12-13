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
	//constraints for username
	usrName := "^[a-zA-Z0-9]{5,10}$"
	userCheker, err := regexp.Compile(usrName)
	if err != nil {
		log.Fatalf("%v", err)
	}
	//constraints for password
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

	//initialize dB
	err = database.InitDB()
	if err != nil {
		log.Fatal(err)
	}
	defer database.CloseDB()

	fmt.Print("Do you have an account(Y/N): ")
	_, err = fmt.Scanf("%s", &newUser)
	if err != nil {
		log.Println("Error reading user response:", err)
		return
	}
	if newUser == "N" || newUser == "n" {
		fmt.Println("Username(must be alphanumeric, 5-10 characters	): ")
		_, err := fmt.Scanf("%s", &user)
		if err != nil {
			log.Println("Error readding username: ", err)
			return
		}
		fmt.Println("Password(must be at least 8 characters): ")
		_, err = fmt.Scanf("%s", &pass)
		if err != nil {
			log.Println("Error reading password: ", err)
			return
		}
		//try to match username and password, if it matches constraints
		if userCheker.MatchString(user) && passChecker.MatchString(pass) {
			userNameNExist = user
			passwordNExist = pass
			//Register user to dB
			database.CreateUser(userNameNExist, passwordNExist)
		}
	}
	//if user alredy registered, login
	fmt.Print("Enter your username: ")
	_, err = fmt.Scanf("%s", &userNameExist)
	if err != nil {
		log.Println("Error reading username: ")
		return
	}

	fmt.Print("Enter password: ")
	_, err = fmt.Scanf("%s", &passwordExist)
	if err != nil {
		log.Println("Error Reading password: ", err)
		return
	}

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
		log.Printf("Authentication Error: %v", err)
		return
	}

}
