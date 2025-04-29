// package main

import (
	"encoding/csv"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/wcharczuk/go-chart/v2"

	"database/sql"

	_ "github.com/mattn/go-sqlite3" // SQLite driver
)

func main() {
	rand.Seed(time.Now().UnixNano()) // Initialize randomness with the current time

	// 🎯 Introduction
	// 🎯 Ask the surname
	var surname string
	fmt.Println("Please enter your surname:")
	fmt.Scanln(&surname)

	// 🎯 Welcome the player personally
	fmt.Printf("\n👋 Welcome, %s! Let's start the Fortune Teller CLI Game!\n", surname)
	fmt.Println("            -            ")
	fmt.Println("🎯 Welcome to the Hassane Fortune Teller CLI Game 🔮")
	fmt.Println("            -            ")

	// Connect to SQLite database
	db, err := sql.Open("sqlite3", "fortunes.db")
	if err != nil {
		fmt.Println("Error opening database:", err)
		return
	}
	defer db.Close()

	// Create table if not exists
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS history (
		surname TEXT,
		fortune TEXT,
		timestamp TEXT
	)`)
	if err != nil {
		fmt.Println("Error creating table:", err)
		return
	}

	// 📝 Open (or create) the CSV file "history.csv" to save fortunes
	file, err := os.OpenFile("history.csv", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening/creating the file:", err)
		return
	}
	defer file.Close() // Close file when the function finishes

	writer := csv.NewWriter(file)
	defer writer.Flush() // Make sure everything is written before ending the program

	var choice string           // User input for continuing (y/n)
	var fortunesPlayed []string // 🧠 Save fortunes played in memory for final analysis

	for {
		randomIndex := rand.Intn(len(Fortunes)) // Randomly pick a fortune
		selectedFortune := Fortunes[randomIndex]

		fmt.Println("\nShuffling the crystal ball... 🔮✨")
		time.Sleep(1 * time.Second) // Small delay for dramatic effect

		var finalFortune string // Fortune after replacing variables (year, talent, money)

		switch selectedFortune {
		case "You will become a millionaire by %d 🚀!":
			randomYear := rand.Intn(63) + 2025
			finalFortune = fmt.Sprintf(selectedFortune, randomYear)

		case "You will discover a hidden talent: %s 🎨!":
			randomTalent := Talents[rand.Intn(len(Talents))]
			finalFortune = fmt.Sprintf(selectedFortune, randomTalent)

		case "You will find money on the ground: %d yen 💵!":
			randomMoney := rand.Int63n(100000000000) + 1
			finalFortune = fmt.Sprintf(selectedFortune, randomMoney)

		default:
			finalFortune = selectedFortune
		}

		_, err = db.Exec(`INSERT INTO history (surname, fortune, timestamp) VALUES (?, ?, ?)`,
			surname, finalFortune, time.Now().Format(time.RFC3339))
		if err != nil {
			fmt.Println("Error inserting into database:", err)
		}

		// 🎯 Display the fortune
		fmt.Println("Your future says:", finalFortune)

		// ✍️ Save fortune + timestamp into CSV
		err := writer.Write([]string{finalFortune, time.Now().Format(time.RFC3339)})
		if err != nil {
			fmt.Println("Error writing to CSV:", err)
		}
		writer.Flush()

		// 🧠 Save fortune into memory for later analysis
		fortunesPlayed = append(fortunesPlayed, finalFortune)

		// 🔄 Ask if the user wants another fortune
		fmt.Println("\n🔄 Do you want another fortune? (y/n): ")
		fmt.Scan(&choice)
		choice = strings.ToLower(choice)

		if choice != "y" {
			// If user says no -> end the game
			fmt.Println("\n🙏 Thank you for playing, Hassane's Fortune Teller 🔮!")
			fmt.Println("The time is", time.Now())
			fmt.Println("            by Hassane            ")
			fmt.Println("            -            ")
			break
		}
	}

	// 📊 After the game ends, show a graph of the session
	showGraph(fortunesPlayed)
}

// 📊 Function to show a simple text-based graph (bar graph)
// 📊 Function to show a graph AND save it as an image (graph.png)
func showGraph(fortunes []string) {
	counts := make(map[string]int)

	// Count each fortune
	for _, fortune := range fortunes {
		counts[fortune]++
	}

	total := len(fortunes)

	fmt.Println("\n📊 Fortune Analysis 📊")

	// Display simple text bar
	for fortune, count := range counts {
		percentage := float64(count) / float64(total) * 100
		bar := strings.Repeat("█", int(percentage/5))
		fmt.Printf("- %s\n  (%d times - %.2f%%) %s\n\n", fortune, count, percentage, bar)
	}

	// Prepare real graph data
	var bars []chart.Value
	for fortune, count := range counts {
		bars = append(bars, chart.Value{
			Value: float64(count),
			Label: fortune,
		})
	}

	// 🖼 Create the real bar chart
	graph := chart.BarChart{
		Title:    "Fortune Distribution",
		Height:   512,
		BarWidth: 60,
		Bars:     bars,
		YAxis: chart.YAxis{
			Range: &chart.ContinuousRange{
				Min: 0,
				Max: findMax(counts) + 1, // 👈 find maximum count and add +1 to be safe
			},
		},
	}

	// Create output file
	file, err := os.Create("graph.png")
	if err != nil {
		fmt.Println("Error creating graph file:", err)
		return
	}
	defer file.Close()

	// 🛠️ Properly render the chart into the file
	err = graph.Render(chart.PNG, file)
	if err != nil {
		fmt.Println("Error rendering graph:", err)
		return
	}

	fmt.Println("✅ Graph saved as 'graph.png'. Open it to see!")
}

func findMax(counts map[string]int) float64 {
	max := 0
	for _, count := range counts {
		if count > max {
			max = count
		}
	}
	return float64(max)
}
