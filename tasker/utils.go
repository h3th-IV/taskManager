package tasker

import (
	"fmt"
	"log"
	"regexp"

	"github.com/h3th-IV/taskManager/database"
	"github.com/joho/godotenv"
)

type AuthDetails struct {
	Auth     string
	Username string
}

var authDetails AuthDetails

var (
	option   int
	UserList = &TaskManager{}
)

func Login() {
	var (
		user           string
		pass           string
		userNameNExist string
		passwordNExist string
		newUser        string
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

	fmt.Print("Do you have an account(Y/N): ")
	_, err = fmt.Scanf("%s", &newUser)
	if err != nil {
		log.Println("Error reading user response:", err)
		return
	}
	if newUser == "N" || newUser == "n" {
		fmt.Print("|\t\t\t\tREGISTER  \t\t\t\t\t|\n")
		fmt.Print("Username(must be alphanumeric, 5-10 characters	): ")
		_, err := fmt.Scanf("%s", &user)
		if err != nil {
			log.Println("Error readding username: ", err)
			return
		}
		fmt.Print("Password(must be at least 8 characters): ")
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
	fmt.Print("|\t\t\t\t\tLOGIN  \t\t\t\t\t\t|\n")
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
		return
	}
	authDetails.Auth = auth
	authDetails.Username = userNameExist

}

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
