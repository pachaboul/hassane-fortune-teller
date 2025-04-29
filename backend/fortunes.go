package backend

import (
	"fmt"
	"math/rand"
	"time"
)

var Fortunes = []string{
	"You will become a millionaire by %d 🚀!",
	"A pigeon is plotting against you 🐦!",
	"Someone will compliment your hair soon 💇‍♂️!",
	"Beware of tacos 🌮 today!",
	"You will discover a hidden talent: %s 🎨!",
	"An unexpected journey awaits you ✈️!",
	"Your code will compile perfectly on the first try 🧑‍💻!",
	"You will find money on the ground: %d yen 💵!",
	"Beware of sneaky cats 🐱!",
	"Great news will come in your email 📧!",
	"You will become invisible for 5 minutes 🫥!",
	"Your favorite food will betray you 🍕!",
	"A bird will deliver an important message 🦅!",
	"You will wake up with super strength 💪!",
	"Be cautious of mysterious elevators 🛗!",
	"Your singing voice will charm someone 🎤!",
	"You will receive a mysterious gift 🎁!",
	"Unexpected rain will bless your day 🌧️!",
	"You will find a secret passage in your city 🛤️!",
	"You will meet a talking cat 🐈‍⬛!",
}

func GenerateFortune(surname string) string {
	rand.Seed(time.Now().UnixNano())
	selected := Fortunes[rand.Intn(len(Fortunes))] // ✅ fixed here!

	switch selected {
	case "You will become a millionaire by %d 🚀!":
		randomYear := rand.Intn(20) + 2025
		return fmt.Sprintf(selected, randomYear)

	case "You will discover a hidden talent: %s 🎨!":
		randomTalent := Talents[rand.Intn(len(Talents))]
		return fmt.Sprintf(selected, randomTalent)

	case "You will find money on the ground: %d yen 💵!":
		randomMoney := rand.Int63n(1000000000) + 1
		return fmt.Sprintf(selected, randomMoney)

	default:
		return selected
	}
}
