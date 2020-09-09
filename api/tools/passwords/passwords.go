package main

import (
	"log"
	"os"

	"github.com/dennischiu/govuekuber/api/models/passwords"
)

func main() {
	password := os.Args[1]
	hash, err := passwords.Encrypt(password)
	if err != nil {
		panic(err)
	}
	log.Println(password + " ----> " + hash)
}
