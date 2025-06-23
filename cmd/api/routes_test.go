package main

import (
	"net/http"
	"testing"

	"github.com/go-chi/chi/v5"
)

func Test_routes_exist(t *testing.T) {
	// why ignore DB, Models
	// because we dont access in main.go
	testApp := Config{}

	// returns http handler (sets up everything before)
	testRoutes := testApp.routes()

	// type assertion to chi.Router
	chiRoutes := testRoutes.(chi.Router)

	routes := []string{"/authenticate"}

	for _, route := range routes {
		routeExists(t, chiRoutes, route)
	}
}

func routeExists(t *testing.T, routes chi.Router, route string) {
	found := false

	// for each route, call callback
	// compare route, which we pass in with what chi gets when walking through our route
	_ = chi.Walk(routes, func(method string, foundRoute string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		if route == foundRoute {
			found = true
		}
		return nil
	})

	if !found {
		t.Errorf("did not find %s in registered routes", route)
	}
}
