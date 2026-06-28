package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello from test server..."))
	})

	r.Post("/refresh", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Refreshing..."))
	})

	r.Post("/login", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Login request received...")
	})

	r.Post("/register", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Register request received...")
	})

	r.Group(func(r chi.Router) {
		// r.Use(requiresAuthMiddleware)
		r.Get("/me", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("This route should be protected, require auth."))
		})

		r.Post("/logout", func(w http.ResponseWriter, r *http.Request) {
			fmt.Println("Logging out")
		})
	})

	fmt.Println("Server started on port :8000")
	if err := http.ListenAndServe(":8000", r); err != nil {
		log.Fatal(err)
	}

}
