package main

import (
	"log"
	"net/http"

	"github.com/levshindenis/Auth_JWT-GO/internal/app/config"
	"github.com/levshindenis/Auth_JWT-GO/internal/app/handlers"
	"github.com/levshindenis/Auth_JWT-GO/internal/app/router"
)

func main() {
	var (
		conf config.Config
		mh   handlers.MyHandler
	)

	conf.ParseFlags()

	if conf.GetDBAddress() == "" {
		log.Fatalf("Set DB_ADDRESS")
	}

	if err := mh.Init(conf); err != nil {
		panic(err)
	}

	if err := http.ListenAndServe(conf.GetServerAddress(), router.Router(mh)); err != nil {
		panic(err)
	}
}
