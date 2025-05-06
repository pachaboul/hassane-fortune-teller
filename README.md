
# 🔮 Hassane's Fortune Teller App

Welcome to **Hassane's Fortune Teller** — a magical fortune-telling CLI and GUI application built with **Go (Golang)**, using **AI integration** via **Mistral model** powered by **Ollama**!

✨ Discover random fortunes, AI-generated magical sentences, track your fortune history, view your luck statistics, and more!

---

## 🚀 Features

- 🎯 Fortune Teller CLI and GUI (built with [Fyne](https://fyne.io/))
- 🔮 Random Fortunes (Pre-written fortunes)
- 🤖 **AI-Generated Fortunes** using **Mistral model** (via Ollama local server)
- 📊 Personal Graph of Fortune Distribution
- 🧪 Statistics: Total fortunes, most frequent fortune, luck score
- 🗓 Timeline of all fortunes (with timestamps)
- 🛡️ Offline First: **Works even without internet!**
- 💾 Fortune history saved locally
- 📤 Future: Export to CSV (coming soon!)

---

## 📸 Screenshots

| Home | After Stop |
|:--|:--|
| ![Home Screen](./screenshots/home.png) | ![Fortune Report](./screenshots/report.png) |

*(You can add real screenshots after)*

---

## 🛠️ Installation

### 1. Clone the project

```bash
git clone https://github.com/pachaboul/hassane-fortune-teller.git
cd hassane-fortune-teller
```

### 2. Install Go modules

```bash
go mod tidy
```

### 3. Install and Run Ollama (for AI)

- Download Ollama from [https://ollama.com/download](https://ollama.com/download)
- Install and start Ollama
- Run Mistral model locally:

```bash
ollama run mistral
```

Ollama will open a local API server at `http://localhost:11434`.

### 4. Run the App

```bash
go run .
```

👉 Enjoy predicting your magical future!

---

## ⚙️ Project Structure

```bash
.
├── backend/
│   ├── ai.go             # Connects to Ollama AI
│   ├── db.go             # SQLite3 database logic
│   ├── fortunes.go       # Pre-written fortune data
│   ├── talents.go        # Talent list for fortunes
│   └── statistics.go     # Fortune statistics (luck score, etc.)
├── graphs/               # Saved fortune graphs
├── history.csv           # Saved fortune history
├── main.go               # Main application (CLI + GUI)
├── fortunes.db           # SQLite3 database (auto created)
├── go.mod
├── go.sum
└── README.md             # 📄 (this file)
```

---

## 📚 Technologies Used

- [GoLang](https://golang.org/)
- [Fyne](https://fyne.io/) (for GUI)
- [SQLite3](https://www.sqlite.org/)
- [Ollama](https://ollama.com/) (to run Mistral model locally)
- [Mistral AI Model](https://mistral.ai/) (language generation)

---

## 👌 Author

- **Aboul Hassane Cissé**  
  [@pachaboul](https://github.com/pachaboul)

---

## 📣 Future Improvements

- [ ] Export fortunes to CSV
- [ ] Add Leaderboard for most lucky players
- [ ] Fine-tune Mistral model for custom fortune-teller mode
- [ ] Allow theme selection (Fantasy, Adventure, Romance)

---

## 📜 License

This project is licensed under the MIT License.
