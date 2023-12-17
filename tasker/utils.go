package tasker

import (
	"fmt"
	"log"
	"regexp"
	"time"

	"github.com/h3th-IV/taskManager/database"
	"github.com/joho/godotenv"
)

// authentiction details
type AuthDetails struct {
	Auth     string
	Username string
}

var authDetails AuthDetails

var (
	option   int
	UserList = &TaskManager{}
)
var (
	user             string
	pass             string
	userNameNExist   string
	passwordNExist   string
	checkIfUserExist rune
)

func Register() {
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

	fmt.Print("Do you have an account(y/n)->  ")
	_, err = fmt.Scanf("%c", &checkIfUserExist)
	if err != nil {
		log.Println("Error reading user response:", err)
		return
	}
	fmt.Scanln()
	if checkIfUserExist == 'n' {
		fmt.Print("\t\t\t\tREGISTER  \t\t\t\t\t\n")
		fmt.Print("Username(must be alphanumeric, 5-10 characters	)-> ")
		_, err := fmt.Scanf("%s", &user)
		if err != nil {
			log.Println("Error readding username: ", err)
			return
		}
		fmt.Print("Password(must be at least 8 characters)->  ")
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
		} else {
			log.Println("Invalid Username or password charctrers")
			return
		}
	}
	invalidNewuser := checkIfUserExist >= 'a' && checkIfUserExist <= 'z' && checkIfUserExist != 'y' && checkIfUserExist != 'n'
	if invalidNewuser {
		log.Println("Unrecognized input. Please enter 'N' or 'Y' to login. Exiting...")
		time.Sleep(2 * time.Second)
		return
	}
	//if user alredy registered, login
	fmt.Print("\t\t\t\t\tLOGIN  \t\t\t\t\t\t\n")
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
	fmt.Println("Logging in...")
	time.Sleep(2 * time.Second)
	auth, err := database.SelectDetails(userNameExist)
	if err != nil {
		log.Fatal("Unable to retrieve user deatils: Incorrect user details")
		return
	}
	authDetails.Auth = auth
	authDetails.Username = userNameExist
}

func parseTime(input string) (time.Time, error) {
	//time is collected as string then parsed as time.Time
	layout := "06-01-02 15:04:05" // YY-MM-DD HH:MM:SS
	parsedTime, err := time.Parse(layout, input)
	if err != nil {
		return time.Time{}, err
	}
	return parsedTime, nil
}
func printMenu() int {
	fmt.Println("\nTask Manager CLI")
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
