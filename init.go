package getrational

import (
	"log"
	"math/rand"
	"net/http"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())

	templ, err := loadTemplates()
	if err != nil {
		log.Fatalf("Can not load templates: %s", err)
	}

	questions, err := readDatabase()
	if err != nil {
		log.Fatalf("Can not read database: %s", err)
	}

	http.Handle("/api/questions/random", questionHandler(questions))

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static/"))))
	http.Handle("/play", playHandler(templ, questions))
	http.Handle("/", indexHandler(templ))
}