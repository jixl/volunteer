package scores

import (
	"log"
	"math/rand"
	// "net/http"
	"time"
)

type (
	total struct {
		Count string `json:"num"`
	}

	eRecord struct {
		Count int
		Page  string
	}
)

func (e eRecord) isExit(page string) bool {
	if e.Page != page {
		e.Count = 0
		e.Page = page
		return false
	}
	e.Count++
	return e.Count >= 3
}

func sleep() {
	sleeps := [6]time.Duration{1, 2, 3, 2, 3, 1}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	index := r.Intn(6)

	log.Println("SLEEP:", sleeps[index])
	time.Sleep(sleeps[index] * time.Second)
}

// type spider interface {
// 	scores(uri string) (int, error)
// 	parse(resp *http.Response) (int, error)
// }
