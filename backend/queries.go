package backend

import (
	"fmt"

	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)

func GetAllFortunes() ([]string, error) {
	rows, err := DB.Query(`
		SELECT f.title, u.surname
		FROM fortunes f
		JOIN users u ON f.added_by = u.id
		ORDER BY f.created_at DESC
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []string
	for rows.Next() {
		var title, surname string
		if err := rows.Scan(&title, &surname); err != nil {
			return nil, err
		}
		results = append(results, fmt.Sprintf("%s → %s", surname, title))
	}
	return results, nil
}
