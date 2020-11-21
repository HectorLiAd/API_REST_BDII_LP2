package main

import (
	"net/http"
	"os"

	// "github.com/API_REST_BDII_LP2/database"

	"github.com/HectorLiAd/API_REST_BDII_LP2/database"
	"github.com/HectorLiAd/API_REST_BDII_LP2/handler"
	"github.com/HectorLiAd/API_REST_BDII_LP2/persona"
	"github.com/go-chi/chi"
)

func main() {
	db := database.InitDB()
	defer db.Close()

	var personaRepository = persona.NewRepository(db)
	var personaServicio = persona.NerService(personaRepository)

	r := chi.NewRouter()
	r.Use(handler.GetCors().Handler)

	r.Mount("/persona", persona.MakeHTTPHandler(personaServicio))

	// dandole un puerto segun lo que requiera heroku
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	http.ListenAndServe(":"+port, r)
}
