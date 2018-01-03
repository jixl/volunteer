package scores

import (
	"log"
	"math/rand"
	"time"
)

type (
	total struct {
		Count string `json:"num"`
	}

	eRecord struct {
		Count int
		Page  int
	}
)

func isExit(page int) bool {
	if total.Page != page {
		total.Count = 0
		total.Page = page
		return false
	}
	total.Count++
	return total.Count >= 3
}

func sleep() {
	sleeps := [6]time.Duration{1, 2, 3, 2, 3, 4}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	index := r.Intn(6)

	log.Println("SLEEP:", sleeps[index])
	time.Sleep(sleeps[index] * time.Second)
}
