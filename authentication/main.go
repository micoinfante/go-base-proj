package main

import (
	"authentication/db"
	"flag"
	"github.com/joho/godotenv"
	"log"
)

var (
	local bool
)

func init() {
	flag.BoolVar(&local, "local", true, "run service local")
	flag.Parse()
}

func main() {
	if local {
		err := godotenv.Load()
		if err != nil {
			log.Panic(err)
		}
	}
	config := db.NewConfig()
	conn, err := db.NewConnection(config)
	if err != nil {
		log.Panicln(err)
	}

	defer conn.Close()
}
