package server

import (
	"github.com/levshindenis/Auth_JWT-GO/internal/app/config"
	"github.com/levshindenis/Auth_JWT-GO/internal/app/database"
)

type Server struct {
	config config.Config
	db     database.Database
}
