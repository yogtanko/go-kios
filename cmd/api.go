package main

import (
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/yogtanko/go-kios/internal/auth"
	myMiddleware "github.com/yogtanko/go-kios/internal/middleware"
	"github.com/yogtanko/go-kios/internal/products"
)

func (app *application) mount() http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Use(middleware.Timeout(60 * time.Second))

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("all good"))
	})

	// middleware
	myMiddleware.Init(app.config.jwtSecret)
	// Services
	productService := products.NewService()

	// Handlers
	authHandler := auth.NewHandler(app.config.jwtSecret)
	productHandler := products.NewHandler(productService)

	r.Post("/login", authHandler.Login)

	r.Group(func(r chi.Router) {
		r.Use(myMiddleware.Authenticator)
		r.Get("/products", productHandler.ListProducts)
	})

	return r
}

func (app *application) run(h http.Handler) error {
	srv := &http.Server{
		Addr:         app.config.addr,
		Handler:      h,
		WriteTimeout: time.Second * 30,
		ReadTimeout:  time.Second * 10,
		IdleTimeout:  time.Minute,
	}
	log.Printf("server has started at addr %s", app.config.addr)
	return srv.ListenAndServe()
}

type application struct {
	config config
}

type config struct {
	addr      string
	db        dbConfig
	jwtSecret []byte
}

type dbConfig struct {
	dsn string
}
