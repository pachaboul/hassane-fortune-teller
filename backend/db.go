package backend

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

func SaveFortune(surname string, fortune string) {
	db, err := sql.Open("sqlite3", "fortunes.db")
	if err != nil {
		fmt.Println("Error opening database:", err)
		return
	}
	defer db.Close()

	// Create the table if it doesn't exist
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS history (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		surname TEXT,
		fortune TEXT,
		timestamp TEXT
	)`)
	if err != nil {
		fmt.Println("Error creating table:", err)
		return
	}

	// Insert the new fortune
	_, err = db.Exec(`INSERT INTO history (surname, fortune, timestamp) VALUES (?, ?, ?)`,
		surname, fortune, time.Now().Format(time.RFC3339))
	if err != nil {
		fmt.Println("Error inserting into database:", err)
		return
	}
}
