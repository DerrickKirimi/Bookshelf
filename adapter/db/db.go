package db

import (
	"database/sql"
	"fmt"

	"github.com/DerrickKirimi/Bookshelf/config"
	"github.com/go-sql-driver/mysql"
)

func New (conf *config.Config) (*sql.DB, error) {
	cfg := &mysql.Config{
		Net:						"tcp",
		Addr:						fmt.Sprintf("%v:%v, conf.Db.Host, conf.Db.Port"),
		DBName: 					conf.Db.DbName,
		User:						conf.Db.Username,
		Passwd: 					conf.Db.Password,
		AllowNativePasswords:	 	true,
		ParseTime:					true,
	}
	return sql.Open("mysql", cfg.FormatDSN())
}