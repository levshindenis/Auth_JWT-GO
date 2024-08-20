package server

import (
	"database/sql"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/levshindenis/Auth_JWT-GO/internal/app/config"
	"github.com/levshindenis/Auth_JWT-GO/internal/app/database"
)

func (s *Server) Init(conf config.Config) error {
	s.config = conf

	newDB, err := sql.Open("pgx", conf.GetDBAddress())
	if err != nil {
		return err
	}

	s.db = database.Database{DB: newDB}
	if err = s.db.MakeTables(); err != nil {
		return err
	}

	return nil
}
