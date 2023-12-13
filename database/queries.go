package database

import (
	"time"
)

// tables won't be created everytime
func CreateUserTable() error {
	query := `CREATE TABLE IF NOT EXISTS users (
		user_id SERIAL PRIMARY KEY,
		username VARCHAR(50) NOT NULL,
		password_hash VARCHAR(100) NOT NULL
	);`
	_, err := dB.Exec(query)
	if err != nil {
		return err
	}
	return nil
}

// Creates User and Update the database
func CreateUser(username, password string) error {
	query := `INSERT INTO users (username, password_hash)
	VALUES ($1, $2);`
	_, err := dB.Exec(query, username, password)
	if err != nil {
		return err
	}
	return nil
}

// retreive userId
func GetUserById(username string) (int, error) {
	query := `SELECT user_id FROM users WHERE username = $1;`
	var userID int
	row := dB.QueryRow(query, username)
	err := row.Scan(&userID)
	if err != nil {
		return 0, err
	}
	return userID, nil
}

// some kinda login mechanism
func SelectDetails(username string) (string, error) {
	query := `SELECT password_hash FROM users where username = $1;`
	var password string
	row := dB.QueryRow(query, username)
	err := row.Scan(&password)
	if err != nil {
		return "", err
	}
	return password, nil
}

// actually tables won't be created everytime so won't be used
func CreateTaskTable() error {
	query := `CREATE TABLE IF NOT EXISTS tasks (
        task_id SERIAL PRIMARY KEY,
        user_id INT,
        FOREIGN KEY (user_id) REFERENCES users(user_id),
        description VARCHAR(255) NOT NULL,
        status VARCHAR(20) NOT NULL,
        start_time TIMESTAMP NOT NULL,
        due_date TIMESTAMP NOT NULL,
        completion_time TIMESTAMP
    );`
	_, err := dB.Exec(query)
	if err != nil {
		return err
	}
	return nil
}

// insert task with user_id
func InsertTask(description, status string, startAt, dueAt time.Time) error {
	query := `INSERT INTO tasks (description, status, start_time, due_date, completion_time)
	VALUES ($1, $2,$3, $4, $5);`
	_, err := dB.Exec(query, description, status, startAt, dueAt, nil)
	if err != nil {
		return err
	}
	return nil
}

// mark task as completed with ....
func MarkTask(status string, completedAt time.Time) error {
	query := `INSERT INTO tasks (status, completion_time)
	VALUES ($1, $2);`
	_, err := dB.Exec(query, status, completedAt)
	if err != nil {
		return err
	}
	return nil
}

func DeleteTask(task_ID int) error {
	query := `DELETE FROM tasks WHERE task_id = $1;`
	_, err := dB.Exec(query, task_ID)
	if err != nil {
		return err
	}
	return nil
}

// --todo hash password omooooo :)
