package repository

import (
	"authentication/db"
	"authentication/models"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"testing"
	"time"
)

func init() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Panicln(err)
	}
}

func TestUsersRepository_Save(t *testing.T) {
	cfg := db.NewConfig()
	conn, err := db.Connect(cfg)
	assert.NoError(t, err)
	defer db.Close(
		conn.Client,
		conn.Ctx,
		conn.Cancel,
	)

	id := primitive.NewObjectID()

	user := &models.User{
		Id:       id,
		Name:     "Test Name",
		Email:    fmt.Sprintf("%s@email.test", id.Hex()),
		Password: "password",
		Created:  time.Now(),
		Updated:  time.Now(),
	}

	r := NewUsersRepository(db.DB(conn))
}
