package main

import (
	"os"
	"strconv"

	DBFactory "github.com/dennischiu/govuekuber/api/db"
	ServerFactory "github.com/dennischiu/govuekuber/api/server"
	"github.com/joho/godotenv"
)

var port int

var dbHost string
var dbPort string
var dbUser string
var dbPassword string
var dbDatabase string

func init() {

	if err := godotenv.Load("../config.ini"); err != nil {
		panic(err)
	}

	rawPort := os.Getenv("PORT")

	if len(rawPort) > 0 {
		var err error
		port, err = strconv.Atoi(rawPort)
		if err != nil {
			panic(err)
		}
	} else {
		port = 8000
	}

	dbHost = os.Getenv("DB_HOST")
	dbPort = os.Getenv("DB_PORT")
	dbUser = os.Getenv("DB_USER")
	dbPassword = os.Getenv("DB_PASSWORD")
	dbDatabase = os.Getenv("DB_DATABASE")

}

func main() {
	db, err := DBFactory.Connect(dbHost, dbPort, dbUser, dbPassword, dbDatabase)
	if err != nil {
		panic(err)
	}
	server := ServerFactory.NewServer(port, db)

	server.Start()
}
