package main

import (
	"net/http"
	"webapp3/pkg/config"
	"webapp3/pkg/handlers"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)

func routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()
	mux.Use(middleware.Recoverer)
	mux.Use(LogRequestInfo)

	mux.Use(NoSurf)
	mux.Use(SetupSession)

	mux.Get("/", handlers.Repo.HomeHandler)
	mux.Get("/about", handlers.Repo.AboutHandler)
	mux.Get("/login", handlers.Repo.LoginHandler)
	mux.Post("/login", handlers.Repo.PostLoginHandler)
	mux.Get("/logout", handlers.Repo.LogoutHandler)

	mux.Get("/page", handlers.Repo.PageHandler)
	mux.Get("/makepost", handlers.Repo.MakePostHandler)

	mux.Post("/makepost", handlers.Repo.PostMakePostHandler)

	mux.Get("/article-received", handlers.Repo.ArticleReceived)


	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))




	return mux
}
