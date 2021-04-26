package db

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"selenium-check-awingu/log"
)

type Sql struct {
	Db            *sqlx.DB
	ConnectString string
}

func (s *Sql) Connect() {

	dataSource := s.ConnectString
	s.Db = sqlx.MustConnect("postgres", dataSource)

	if err := s.Db.Ping(); err != nil {
		log.Error(err.Error())
		return
	}

	fmt.Println("Đã kết nối đến Database")
}

func (s *Sql) Close() {
	s.Db.Close()
}
