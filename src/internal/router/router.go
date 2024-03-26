package router

import (
	"CatRegistry/src/internal/router/routes"
	"net/http"
)

func CreateNewHTTPServer() *http.ServeMux {
	serveMux := http.NewServeMux()
	routes.RegisterRoutes(serveMux)
	return serveMux
}
