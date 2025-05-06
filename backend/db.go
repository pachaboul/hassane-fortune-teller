package backend

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"time"

	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	connStr := "user=affectify dbname=HassanesFortuneTeller sslmode=disable"

	var err error
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Failed to connect to DB:", err)
	}

	// Load and execute init.sql only once
	sqlBytes, err := ioutil.ReadFile("backend/init.sql")
	if err != nil {
		log.Fatal("Cannot read init.sql:", err)
	}

	_, err = DB.Exec(string(sqlBytes))
	if err != nil {
		log.Fatal("Failed to execute init.sql:", err)
	}

	log.Println("✅ Database initialized successfully")
}

// InsertFortune inserts a new fortune tied to a user (by surname)
func InsertFortune(surname string, fortuneText string) error {
	// 1. Look up the user ID from the surname
	var userID int
	//err := DB.QueryRow("SELECT id FROM users WHERE LOWER(surname) = LOWER($1)", surname).Scan(&userID)
	err := DB.QueryRow("SELECT id FROM users WHERE LOWER(name) = LOWER($1)").Scan(&userID)
	if err != nil {
		if err == sql.ErrNoRows {
			return fmt.Errorf("user '%s' not found", surname)
		}
		return err
	}

	// 2. Insert the fortune
	_, err = DB.Exec(`
		INSERT INTO fortunes (title, category, created_at, added_by)
		VALUES ($1, $2, $3, $4)
	`, fortuneText, "custom", time.Now(), userID)
	if err != nil {
		return err
	}

	return nil
}

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
