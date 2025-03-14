package connection

import (
	"database/sql"
	"elearning/app/config"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func DatabaseConnection(config config.Database) *sql.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.User,
		config.Password,
		config.Host,
		config.Port,
		config.Database,
	)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("DB_ERROR:" + err.Error())
	}
	err = db.Ping()
	if err != nil {
		log.Fatal("Gagal Koneksi:" + err.Error())
	}
	//db.SetConnMaxLifetime(time.Minute * 3)
	//db.SetMaxIdleConns(10)
	//db.SetMaxIdleConns(10)
	return db

}
