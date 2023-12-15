package tasker

// []TODO use slice operation to remove task from task list
import (
	"bufio"
	"database/sql"
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

func NewTask() *CompletedTask {

	newScanner := bufio.NewScanner(os.Stdin)
	// var ID uint
	// fmt.Println("+------------------------------------------------+")
	// fmt.Print("Enter the Task ID,(check last Task get new TaskID): ")
	// newScanner.Scan()
	// input := newScanner.Text()
	// _, err := fmt.Sscan(input, &ID)
	// if err != nil {
	// 	log.Printf("%v\n", err)
	// }

	// // if newScanner.Scan() {
	// // 	ID := newScanner.Text()
	// // 	fmt.Printf("Task ID: %s\n", ID)
	// // } else {
	// // 	fmt.Println("Error reading input:", newScanner.Err())
	// // }
	var description string
	fmt.Println("+----------------------+")
	fmt.Println("New Task description: ")
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
			Description: description,
			StartTime:   parsedStartAt,
			DueDate:     parsedDuedate,
		},
		status: "Uncomplete",
	}
	return &tersk
}

func (tm *TaskManager) CreateTask() {
	userName := authDetails.Username
	tersk := NewTask()
	tm.TaskList = append(tm.TaskList, *tersk)
	database.InsertTask(userName, tersk.Description, tersk.status, tersk.StartTime, tersk.DueDate)
	fmt.Println("Task Created Succesfully")
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

func (tm *TaskManager) ListTask() string {
	username := authDetails.Username
	//get tasks for user with user_Id
	taskRow := database.GetTaskList(username)
	fmt.Print("\n\n")
	fmt.Println("\t--------Your Task---------")
	fmt.Println("+---------------------------------------+")
	//iterate through the rows returned
	if taskRow.Next() {
		for {
			var task_ID int
			var user_ID int
			var Description string
			var status string
			var start_time time.Time
			var due_date time.Time
			var completedAt sql.NullTime

			//coollect input from that task row instance retured
			if err := taskRow.Scan(&task_ID, &user_ID, &Description, &status, &start_time, &due_date, &completedAt); err != nil {
				log.Fatal(err)
			}
			if status == "Uncomplete" {
				fmt.Printf("[ ]TaskID: %d\nDescription: %s\nDue by: %s\nStatus: %s\n+---------------------------------------+\n", task_ID, Description, due_date, status)
			} else {
				fmt.Printf("[X]TaskID: %d\nDescription: %s\nDue by: %s\nStatus: %s\n+---------------------------------------+\n", task_ID, Description, due_date, status)
			}
			if !taskRow.Next() {
				break
			}
		}

	} else {
		fmt.Print("you have no tasks\n\n")
		return "false"
	}
	return "true"
}

func (tm *TaskManager) MarkComplete() {
	checker := tm.ListTask()
	if checker == "true" {
		newScanner := bufio.NewScanner(os.Stdin)

		var ID uint
		fmt.Print("Enter the Task ID, to be Marked complete: ")
		newScanner.Scan()
		input := newScanner.Text()
		_, err := fmt.Sscan(input, &ID)
		if err != nil {
			log.Printf("%v\n", err)
			fmt.Printf("error: unable to get TaskID")
			return
		}
		database.MarkTask(int(ID))
	} else {
		fmt.Println("No task to mark as complete")
	}
}

func (tm *TaskManager) RemoveTask() {
	checker := tm.ListTask()
	if checker == "true" {
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
		database.DeleteTask(int(ID))
	}

}
