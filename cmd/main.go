package main

import (
	"fmt"
	"net/http"

	"github.com/GhazaFadhlilbaqi/Konvert/backend/routes" //routes
)

func main() {
	routes.Routes()
	fmt.Println("Server starting on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
