package main

import (
	"log"

	"github.com/MohamedAklamaash/GO_Simple_CRUD_Backend_WIth_SQL/cmd/api"
	"github.com/MohamedAklamaash/GO_Simple_CRUD_Backend_WIth_SQL/config"
	"github.com/MohamedAklamaash/GO_Simple_CRUD_Backend_WIth_SQL/db"
	"github.com/go-sql-driver/mysql"
)

func main() {
	db, err := db.NewMySQLStorage(mysql.Config{
		User:                 config.GetConfigs().DbUserName,
		Passwd:               config.GetConfigs().DBPassword,
		Addr:                 config.GetConfigs().DBAddress,
		DBName:               config.GetConfigs().DBName,
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	})
	if err != nil {
		log.Fatal(err)
	}
	server := api.NewAPIServer(":8080", db)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}

}
