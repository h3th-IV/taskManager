package main

import (
	"fmt"

	"github.com/h3th-IV/taskManager/database"
	"github.com/h3th-IV/taskManager/tasker"
)

func main() {
	fmt.Println("TASK BULL developed @bool")
	tasker.TaskManagerApp()
	defer database.CloseDB()
}
