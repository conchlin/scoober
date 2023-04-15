package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func main() {
	config := LoadConfig()
	initDatabase(config.Database.Host, config.Database.Port,
		config.Database.User, config.Database.Password, config.Database.Dbname)
}

func initDatabase(host string, port string, user string, password string, dbname string) {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected to Scoober")
}
