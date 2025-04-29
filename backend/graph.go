package backend

import (
	"database/sql"
	"fmt"
	"os"
	"strings"

	_ "github.com/mattn/go-sqlite3"
	"github.com/wcharczuk/go-chart/v2"
)

// GenerateGraph creates a graph only for the current player
func GenerateGraph(surname string) {
	db, err := sql.Open("sqlite3", "fortunes.db")
	if err != nil {
		fmt.Println("Error opening database:", err)
		return
	}
	defer db.Close()

	rows, err := db.Query(`SELECT fortune FROM history WHERE surname = ?`, surname)
	if err != nil {
		fmt.Println("Error querying database:", err)
		return
	}
	defer rows.Close()

	counts := make(map[string]int)
	for rows.Next() {
		var fortune string
		if err := rows.Scan(&fortune); err == nil {
			counts[fortune]++
		}
	}

	if len(counts) == 0 {
		fmt.Println("No fortunes for this player to graph.")
		return
	}

	var bars []chart.Value
	for fortune, count := range counts {
		bars = append(bars, chart.Value{
			Value: float64(count),
			Label: fortune,
		})
	}

	graph := chart.BarChart{
		Title:    surname + "'s Fortune Distribution",
		Height:   512,
		BarWidth: 60,
		Bars:     bars,
	}

	// ✅ Create "graphs/" directory if it doesn't exist
	if _, err := os.Stat("graphs"); os.IsNotExist(err) {
		err := os.Mkdir("graphs", 0755)
		if err != nil {
			fmt.Println("Error creating graphs folder:", err)
			return
		}
	}

	// ✅ Create file name with lowercase surname
	filename := fmt.Sprintf("graphs/%s.png", strings.ToLower(surname))

	f, err := os.Create(filename)
	if err != nil {
		fmt.Println("Error creating graph file:", err)
		return
	}
	defer f.Close()

	err = graph.Render(chart.PNG, f)
	if err != nil {
		fmt.Println("Error rendering graph:", err)
		return
	}

	fmt.Println("✅ Graph saved as", filename)
}
