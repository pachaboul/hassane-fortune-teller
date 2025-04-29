package backend

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

var fortuneWeights = map[string]int{
	"millionaire":    5,
	"gift":           4,
	"money":          3,
	"compliment":     2,
	"unexpected":     2,
	"journey":        1,
	"super strength": 2,
	"invisible":      2,
	"mysterious":     1,
	"singing voice":  2,
	"betray you":     -2,
	"sneaky cats":    -1,
	"elevator":       -1,
	"taco":           -1,
	"pigeon":         -2,
}

type FortuneStat struct {
	Total        int
	MostFrequent string
	LuckScore    int
	Timeline     []string
}

// AnalyzeFortunes computes stats for one player
func AnalyzeFortunes(surname string) FortuneStat {
	db, err := sql.Open("sqlite3", "fortunes.db")
	if err != nil {
		fmt.Println("Error opening DB:", err)
		return FortuneStat{}
	}
	defer db.Close()

	rows, err := db.Query(`SELECT fortune, timestamp FROM history WHERE surname = ? ORDER BY timestamp DESC`, surname)
	if err != nil {
		fmt.Println("Error querying DB:", err)
		return FortuneStat{}
	}
	defer rows.Close()

	counts := make(map[string]int)
	timeline := []string{}
	luck := 0
	total := 0

	for rows.Next() {
		var fortune, ts string
		if err := rows.Scan(&fortune, &ts); err == nil {
			counts[fortune]++
			timeline = append(timeline, fmt.Sprintf("[%s] %s", ts[:19], fortune))
			total++
			for keyword, score := range fortuneWeights {
				if containsInsensitive(fortune, keyword) {
					luck += score
					break
				}
			}
		}
	}

	// Find most frequent
	max := 0
	mostFrequent := "N/A"
	for f, c := range counts {
		if c > max {
			max = c
			mostFrequent = f
		}
	}

	return FortuneStat{
		Total:        total,
		MostFrequent: mostFrequent,
		LuckScore:    luck,
		Timeline:     timeline,
	}
}

func containsInsensitive(text, word string) bool {
	return len(text) >= len(word) && (containsFold(text, word) || containsFold(text, capitalize(word)))
}

func containsFold(a, b string) bool {
	return len(a) >= len(b) && (stringIndexFold(a, b) != -1)
}

func stringIndexFold(s, substr string) int {
	return len([]rune(s[:len(s)])) - len([]rune(substr)) + 1
}

func capitalize(s string) string {
	if len(s) == 0 {
		return s
	}
	return string(s[0]&^0x20) + s[1:]
}
