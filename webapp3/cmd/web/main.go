package main

import (
	"encoding/gob"
	"log"
	"net/http"
	"time"
	"webapp3/models"
	"webapp3/pkg/config"
	"webapp3/pkg/dbdriver"
	"webapp3/pkg/handlers"
	"webapp3/pkg/render"

	"github.com/alexedwards/scs/v2"
)

var sessionManager *scs.SessionManager
var app config.AppConfig

func main() {

	db, err := run()
	if err != nil {
		log.Fatal(err)
	}
	defer db.SQL.Close()

	srv := &http.Server{
		Addr:    ":8080",
		Handler: routes(&app),
	}
	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

}

func run()(*dbdriver.DB, error) {
	gob.Register(models.Article{})

	gob.Register(models.User{})
	gob.Register(models.Post{})

	sessionManager = scs.New()
	sessionManager.Lifetime = 24 * time.Hour
	sessionManager.Cookie.Persist = true
	sessionManager.Cookie.Secure = false
	sessionManager.Cookie.SameSite = http.SameSiteLaxMode
	app.Session = sessionManager

	db, err := dbdriver.ConnectSQL("host=localhost port=5432 dbname=blog_db user=postgres password=root") 

	if err != nil {
		log.Fatal("Can't connect to database")
	}
	repo := handlers.NewRepo(&app, db)
	handlers.NewHandlers(repo)
	render.NewAppConfig(&app)
	return db, nil

}