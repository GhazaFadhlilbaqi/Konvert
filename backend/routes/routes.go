package routes

import (
	"net/http"

	"github.com/GhazaFadhlilbaqi/Konvert/backend/handlers" //handlers functions (i.e views.py)
)

func Routes() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/converter", handlers.Converter)
}
