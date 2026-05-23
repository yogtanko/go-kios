package main

import (
	"log/slog"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/yogtanko/go-kios/internal/auth"
	myMiddleware "github.com/yogtanko/go-kios/internal/middleware"
	"github.com/yogtanko/go-kios/internal/postgress"
	"github.com/yogtanko/go-kios/internal/products"
	"github.com/yogtanko/go-kios/pkg/config"
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
	var err error
	app.db, err = postgress.NewDatabase(&app.config.Db.DBUrl)
	if err != nil {
		slog.Error("Gagal create koneksi ke DB", "error", err.Error())
		return nil
	}
	// middleware
	myMiddleware.Init(app.config.JwtSecret)
	// Services
	productService := products.NewService(app.db)

	// Handlers
	authHandler := auth.NewHandler(app.config.JwtSecret)
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
		Addr:         app.config.Addr,
		Handler:      h,
		WriteTimeout: time.Second * 30,
		ReadTimeout:  time.Second * 10,
		IdleTimeout:  time.Minute,
	}
	slog.Info("server has started at addr " + app.config.Addr)
	return srv.ListenAndServe()
}

type application struct {
	config *config.Config
	db     *postgress.Database
}
