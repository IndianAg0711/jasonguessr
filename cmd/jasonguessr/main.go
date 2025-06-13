package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/golang-migrate/migrate/v4"
)

func main() {
	fmt.Println("Hello, World!")
	r := chi.NewRouter()
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		res, _ := json.Marshal("welcome")
		w.Write(res)
	})
	http.ListenAndServe(":3030", r)
}
