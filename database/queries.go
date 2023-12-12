package database

func CreateUserTable() error {
	query := `CREATE TABLE IF NOT EXISTS users (
		user_id SERIAL PRIMARY KEY,
		username VARCHAR(50) NOT NULL,
		password_hash VARCHAR(100) NOT NULL, --todo hash password omooooo :)
		-- Add other user-related fields as needed
	);`
	_, err := dB.Exec(query)
	if err != nil {
		return err
	}
	return nil
}
func CreateTaskTable() error {
	query := `CREATE TABLE IF NOT EXISTS tasks (
		task_id SERIAL PRIMARY KEY,
		user_id INT REFERENCES users(user_id),
		description VARCHAR(255) NOT NULL,
		status VARCHAR(20) NOT NULL,
		start_time TIMESTAMP NOT NULL,
		due_date TIMESTAMP NOT NULL,
		completion_time TIMESTAMP,
	);`
	_, err := dB.Exec(query)
	if err != nil {
		return err
	}
	return nil
}

func CreateUser(username, password string) error {
	query := `INSERT INTO users (username, password_hash)
	VALUES ($1, $2, $3);`
	_, err := dB.Exec(query, username, password)
	if err != nil {
		return err
	}
	return nil
}

// --todo hash password omooooo :)
