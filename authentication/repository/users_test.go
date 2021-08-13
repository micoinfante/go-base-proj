package repository

import (
	"authentication/db"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func init() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Panicln(err)
	}
}

func TestUsersRepository_Save(t *testing.T) {
	cfg := db.NewConfig()
	client, ctx, cancel, err := db.Connect(cfg.Dsn())
	assert.NoError(t, err)
	defer db.Close(client, ctx, cancel)

	// TODO add test
}
