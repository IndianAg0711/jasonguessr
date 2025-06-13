package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/IndianAg0711/jasonguessr/internal/db"
	"github.com/go-chi/chi/v5"
)

func main() {
	fmt.Println("Hello, World!")

	// run any pending db migrations
	db.Migrate()

	r := chi.NewRouter()
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		res, _ := json.Marshal("welcome")
		w.Write(res)
	})
	http.ListenAndServe(":3030", r)
}
