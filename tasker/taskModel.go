package tasker

// []TODO use slice operation to remove task from task list
import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/h3th-IV/taskManager/database"
)

type InitTask struct {
	Marker      string
	TaskID      uint
	Description string
	StartTime   time.Time
	DueDate     time.Time
}

type CompletedTask struct {
	InitTask
	status         string
	CompletionTime time.Time
}

type TaskManager struct {
	TaskList []CompletedTask
}

func (tm *TaskManager) CreateTask() {
	newScanner := bufio.NewScanner(os.Stdin)

	//logic for adding task to TaskList

	var ID uint
	fmt.Println("+------------------------------------------------+")
	fmt.Print("Enter the Task ID,(check last Task get new TaskID): ")
	newScanner.Scan()
	input := newScanner.Text()
	_, err := fmt.Sscan(input, &ID)
	if err != nil {
		log.Printf("%v\n", err)
	}

	// // if newScanner.Scan() {
	// // 	ID := newScanner.Text()
	// // 	fmt.Printf("Task ID: %s\n", ID)
	// // } else {
	// // 	fmt.Println("Error reading input:", newScanner.Err())
	// // }
	var description string
	fmt.Println("+----------------------+")
	fmt.Println("Task description: ")
	if newScanner.Scan() {
		description = newScanner.Text()
	} else {
		log.Printf("Error reading input: %v\n", newScanner.Err())
	}
	description = "(" + description + ")"

	fmt.Println("+---------------------------------------+")
	var startAt string
	fmt.Print("Time to begin Task(format YY-MM-DD HH:MM:SS): ")
	if newScanner.Scan() {
		startAt = newScanner.Text()
	} else {
		log.Printf("Error reading input: %v\n", newScanner.Err())
	}

	fmt.Println("+---------------------------------------+")
	var DueDate string
	fmt.Print("Task will be due(format YY-MM-DD HH:MM:SS): ")
	if newScanner.Scan() {
		DueDate = newScanner.Text()
	} else {
		log.Printf("Error reading input: %v\n", newScanner.Err())
	}
	//parsetime
	parsedStartAt, err := parseTime(startAt)
	if err != nil {
		log.Printf("%v\n", err)
	}
	parsedDuedate, err := parseTime(DueDate)
	if err != nil {
		log.Printf("%v\n", err)
	}

	tersk := CompletedTask{
		InitTask: InitTask{
			Marker:      "[ ]",
			TaskID:      ID,
			Description: description,
			StartTime:   parsedStartAt,
			DueDate:     parsedDuedate,
		},
		status: "Uncomplete",
	}
	tm.TaskList = append(tm.TaskList, tersk)
	fmt.Println("Task Created Succesfully")
}

func parseTime(input string) (time.Time, error) {
	layout := "06-01-02 15:04:05" // YY-MM-DD HH:MM:SS
	parsedTime, err := time.Parse(layout, input)
	if err != nil {
		return time.Time{}, err
	}
	return parsedTime, nil
}

func (tm *TaskManager) ListTask() {
	fmt.Print("\n\n")
	if len(tm.TaskList) == 0 {
		fmt.Println("You have no task Right now. Rest Up")
	}
	fmt.Println("\t--------Your Task---------")
	fmt.Println("+---------------------------------------+")
	for _, tersk := range tm.TaskList {
		fmt.Printf("%sTaskID: %d\nDescription: %s\nDue by: %s\nStatus: %s\n+---------------------------------------+\n", tersk.Marker, tersk.TaskID, tersk.Description, tersk.StartTime, tersk.status)
	}
}

func (tm *TaskManager) MarkComplete() {
	newScanner := bufio.NewScanner(os.Stdin)

	var ID uint
	fmt.Print("Enter the Task ID, to be Marked complete: ")
	newScanner.Scan()
	input := newScanner.Text()
	_, err := fmt.Sscan(input, &ID)
	if err != nil {
		log.Printf("%v\n", err)
		fmt.Printf("error: unable to get TaskID")
	}

	for i, tersk := range tm.TaskList {
		if tersk.InitTask.TaskID == ID {
			tm.TaskList[i].Marker = "[X]"
			tm.TaskList[i].status = "Completed"
			tm.TaskList[i].CompletionTime = time.Now()
		}
	}
	fmt.Println("Task marked as completed! SCRAM")
}

func (tm *TaskManager) RemoveTask() {
	newScanner := bufio.NewScanner(os.Stdin)

	var ID uint
	fmt.Print("Enter the Task ID, to be Removed: ")
	newScanner.Scan()
	input := newScanner.Text()
	_, err := fmt.Sscan(input, &ID)
	if err != nil {
		log.Printf("%v\n", err)
		fmt.Printf("error: unable to get TaskID")
	}

	for i, tersk := range tm.TaskList {
		if tersk.InitTask.TaskID == ID {
			//slicing
			tm.TaskList = append(tm.TaskList[:i], tm.TaskList[i+1:]...)
			fmt.Println("Task removed succesfully")
		}
	}
	tm.ListTask()
}
