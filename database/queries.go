package database

func CreateTable() error {
	query := `CREATE TABLE IF NOT EXISTS tasks (
		task_id SERIAL PRIMARY KEY,
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
