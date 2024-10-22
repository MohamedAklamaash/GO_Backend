package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/MohamedAklamaash/GO_Simple_CRUD_Backend_WIth_SQL/cmd/api"
	"github.com/MohamedAklamaash/GO_Simple_CRUD_Backend_WIth_SQL/config"
	"github.com/MohamedAklamaash/GO_Simple_CRUD_Backend_WIth_SQL/db"
	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {

	cfg := mysql.Config{
		User:                 config.Envs.DbUserName,
		Passwd:               config.Envs.DBPassword,
		Addr:                 config.Envs.DBAddress,
		DBName:               config.Envs.DBName,
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	}

	db, err := db.NewMySQLStorage(cfg)
	if err != nil {
		log.Fatal(err)
	}
	initStorage(db)

	server := api.NewAPIServer(":8080", db)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}

func initStorage(db *sql.DB) {
	err := db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("DB is running")
}
