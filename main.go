package main

import (
	"fmt"
	"my-first-go/backend"
	"strings"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("🔮 Hassane's Fortune Teller")

	surnameEntry := widget.NewEntry()
	surnameEntry.SetPlaceHolder("Enter your surname here...")

	fortuneLabel := widget.NewLabel("Your fortune will appear here...")

	var playerSurname string

	// Button to start the game
	startButton := widget.NewButton("Start Fortune Teller", func() {
		playerSurname = surnameEntry.Text
		if playerSurname == "" {
			fortuneLabel.SetText("❗ Please enter your surname first!")
			return
		}
		surnameEntry.Disable()
		fortuneLabel.SetText("✅ Hello " + playerSurname + "! Click below to get your fortune 🔮")
	})

	// Button to generate an AI fortune
	aiFortuneButton := widget.NewButton("🤖 Generate AI Fortune", func() {
		if playerSurname == "" {
			fortuneLabel.SetText("❗ Please enter your surname first!")
			return
		}

		fortuneLabel.SetText("🧠 Thinking...") // Show thinking

		go func() {
			fortune := backend.GenerateAIFortune()
			backend.SaveFortune(playerSurname, fortune)

			for i := 1; i <= len(fortune); i++ {
				currentText := "🔮 " + fortune[:i]
				time.Sleep(40 * time.Millisecond)
				fortuneLabel.SetText(currentText)
			}
		}()
	})

	// Button to generate a normal fortune
	getFortuneButton := widget.NewButton("Get New Fortune 🔮", func() {
		if playerSurname == "" {
			fortuneLabel.SetText("❗ Please enter your surname and click Start first!")
			return
		}

		fortune := backend.GenerateFortune(playerSurname)
		backend.SaveFortune(playerSurname, fortune)

		fortuneLabel.SetText("🔮 ")
		go func() {
			for i := 1; i <= len(fortune); i++ {
				currentText := "🔮 " + fortune[:i]
				time.Sleep(40 * time.Millisecond)
				fortuneLabel.SetText(currentText)
			}
		}()
	})

	// Button to stop, show graph, stats, and exit
	stopButton := widget.NewButton("❌ Stop and Show Graph", func() {
		backend.GenerateGraph(playerSurname)

		stats := backend.AnalyzeFortunes(playerSurname)

		// Fortune statistics
		statBox := container.NewVBox(
			widget.NewLabel(fmt.Sprintf("🧮 Total Fortunes: %d", stats.Total)),
			widget.NewLabel(fmt.Sprintf("⭐ Most Frequent: %s", stats.MostFrequent)),
			widget.NewLabel(fmt.Sprintf("🍀 Luck Score: %d", stats.LuckScore)),
		)

		// Timeline of fortunes
		timelineText := strings.Join(stats.Timeline, "\n")
		timelineLabel := widget.NewMultiLineEntry()
		timelineLabel.SetText(timelineText)
		timelineLabel.Disable()
		timelineLabel.Wrapping = fyne.TextWrapWord
		timelineLabel.SetMinRowsVisible(6)

		// Graph Image
		graphFile := fmt.Sprintf("graphs/%s.png", strings.ToLower(playerSurname))
		graphImage := canvas.NewImageFromFile(graphFile)
		graphImage.FillMode = canvas.ImageFillOriginal

		// Exit Button
		exitButton := widget.NewButton("Exit Now 🚪", func() {
			a.Quit()
		})

		// Set all content once
		w.SetContent(container.NewVBox(
			widget.NewLabel("📊 Your Fortune Report"),
			statBox,
			graphImage,
			widget.NewLabel("🗓️ Timeline:"),
			timelineLabel,
			exitButton,
		))
	})

	// Layout
	content := container.NewVBox(
		widget.NewLabel("🎯 Welcome to Hassane's Fortune Teller! 🎯"),
		surnameEntry,
		startButton,
		aiFortuneButton,
		getFortuneButton,
		fortuneLabel,
		stopButton,
	)

	w.Resize(fyne.NewSize(500, 600))
	w.SetContent(content)
	w.ShowAndRun()
}
