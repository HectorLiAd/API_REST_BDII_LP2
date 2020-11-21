package main

import (
	"net/http"
	"os"

	// "github.com/API_REST_BDII_LP2/database"

	"github.com/HectorLiAd/API_REST_BDII_LP2/handler"
	"github.com/go-chi/chi"
)

func main() {
	// db := database.InitDB()
	// defer db.Close()
	r := chi.NewRouter()
	r.Use(handler.GetCors().Handler)
	r.Get("/persona", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome personita 1"))
	})
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	http.ListenAndServe(":"+port, r)
}
