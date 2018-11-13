package scores

import (
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"
)

const firstPage = 1

type (
	total struct {
		Count string `json:"num"`
	}

	category struct {
		year  string
		page  int
		name  string
		execs int
	}
)

func (c category) isExit() bool {
	c.execs++
	return c.execs >= 3
}

func sleep() {
	sleeps := [6]time.Duration{1, 2, 3, 1, 2, 1}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	index := r.Intn(6)

	log.Println("SLEEP:", sleeps[index])
	time.Sleep(sleeps[index] * time.Second)
}

type spider interface {
	getURI() string
	parse(resp *http.Response) (int, error)
	isExit() bool
}

func spiderScores(s spider) (int, error) {
	sleep()
	uri := s.getURI()
	resp, err := http.Get(uri)
	log.Println(uri)
	if err != nil {
		log.Println("ERROR GET: ", uri, err)
		if s.isExit() {
			os.Exit(-1)
		}
		return spiderScores(s)
	}

	return s.parse(resp)
}
