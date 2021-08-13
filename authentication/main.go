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
	client, ctx, cancel, err :=  db.Connect(config.Dsn())
	if err != nil {
		log.Panicln(err)
	}

	defer db.Close(client, ctx, cancel)
}
