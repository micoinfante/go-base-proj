package repository

import (
	"authentication/authentication/models"
	"authentication/db"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"testing"
	"time"
)

func init() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Panicln(err)
	}
	cfg := db.NewConfig()
	conn, err := db.NewConnection(cfg)
	defer conn.Close()

	r := NewUsersRepository(conn)
	err = r.(*usersRepository).DeleteAll()
	if err != nil {
		log.Panicln(err)
	}
}

func TestUsersRepository_Save(t *testing.T) {
	cfg := db.NewConfig()
	conn, err := db.NewConnection(cfg)
	assert.NoError(t, err)
	defer conn.Close()

	id := primitive.NewObjectID()

	user := &models.User{
		Id:       id,
		Name:     "Test Name",
		Email:    fmt.Sprintf("%s@email.test", id.Hex()),
		Password: "password",
		Created:  time.Now(),
		Updated:  time.Now(),
	}

	r := NewUsersRepository(conn)
	err = r.Save(user)
	assert.NoError(t, err)

	found, err := r.GetById(user.Id.Hex())
	assert.NoError(t, err)
	assert.NotNil(t, found)
}

func TestUsersRepository_GetById(t *testing.T) {
	cfg := db.NewConfig()
	conn, err := db.NewConnection(cfg)
	assert.NoError(t, err)
	defer conn.Close()

	id := primitive.NewObjectID()

	user := &models.User{
		Id:       id,
		Name:     "Test Name",
		Email:    fmt.Sprintf("%s@email.test", id.Hex()),
		Password: "password",
		Created:  time.Now(),
		Updated:  time.Now(),
	}

	// conn db.Connection
	r := NewUsersRepository(conn)
	err = r.Save(user)
	assert.NoError(t, err)

	// Existing Check
	found, err := r.GetById(user.Id.Hex())
	assert.NoError(t, err)
	assert.NotNil(t, found)
	assert.Equal(t, user.Id, found.Id)
	assert.Equal(t, user.Email, found.Email)
	assert.Equal(t, user.Name, found.Name)
	assert.Equal(t, user.Password, found.Password)

	// Not existing check
	found, err = r.GetById(primitive.NewObjectID().Hex())
	assert.Error(t, err)
	assert.EqualError(t, mongo.ErrNoDocuments, err.Error())
	assert.Nil(t, found)
}

func TestUsersRepository_GetByEmail(t *testing.T) {
	cfg := db.NewConfig()
	conn, err := db.NewConnection(cfg)
	assert.NoError(t, err)
	defer conn.Close()

	id := primitive.NewObjectID()

	user := &models.User{
		Id:       id,
		Name:     "Test Name",
		Email:    fmt.Sprintf("%s@email.test", id.Hex()),
		Password: "password",
		Created:  time.Now(),
		Updated:  time.Now(),
	}

	// conn db.Connection
	r := NewUsersRepository(conn)
	err = r.Save(user)
	assert.NoError(t, err)

	// Existing Check
	found, err := r.GetByEmail(user.Email)
	assert.NoError(t, err)
	assert.NotNil(t, found)
	assert.Equal(t, user.Id, found.Id)
	assert.Equal(t, user.Email, found.Email)
	assert.Equal(t, user.Name, found.Name)
	assert.Equal(t, user.Password, found.Password)

	// Not existing check
	found, err = r.GetById(primitive.NewObjectID().Hex())
	assert.Error(t, err)
	assert.EqualError(t, mongo.ErrNoDocuments, err.Error())
	assert.Nil(t, found)
}

func TestUsersRepository_Update(t *testing.T) {
	cfg := db.NewConfig()
	conn, err := db.NewConnection(cfg)
	assert.NoError(t, err)
	defer conn.Close()

	id := primitive.NewObjectID()

	user := &models.User{
		Id:       id,
		Name:     "Test Name",
		Email:    fmt.Sprintf("%s@email.test", id.Hex()),
		Password: "password",
		Created:  time.Now(),
		Updated:  time.Now(),
	}

	r := NewUsersRepository(conn)
	err = r.Save(user)
	assert.NoError(t, err)

	found, err := r.GetById(user.Id.Hex())
	assert.NoError(t, err)
	assert.NotNil(t, found)

	user.Name = "Updated Test Name"
	err = r.Update(user)
	assert.NoError(t, err)

	found, err = r.GetById(user.Id.Hex())
	assert.NoError(t, err)
	assert.Equal(t, "Updated Test Name", found.Name)
}

func TestUsersRepository_Delete(t *testing.T) {
	cfg := db.NewConfig()
	conn, err := db.NewConnection(cfg)
	assert.NoError(t, err)
	defer conn.Close()

	id := primitive.NewObjectID()

	user := &models.User{
		Id:       id,
		Name:     "Test Name",
		Email:    fmt.Sprintf("%s@email.test", id.Hex()),
		Password: "password",
		Created:  time.Now(),
		Updated:  time.Now(),
	}

	r := NewUsersRepository(conn)
	err = r.Save(user)
	assert.NoError(t, err)

	found, err := r.GetById(user.Id.Hex())
	assert.NoError(t, err)
	assert.NotNil(t, found)

	err = r.Delete(user.Id.Hex())
	assert.NoError(t, err)

	found, err = r.GetById(user.Id.Hex())
	assert.Error(t, err)
	assert.EqualError(t, mongo.ErrNoDocuments, err.Error())
	assert.Nil(t, found)
}

func TestUsersRepository_GetAll(t *testing.T) {
	cfg := db.NewConfig()
	conn, err := db.NewConnection(cfg)
	assert.NoError(t, err)
	defer conn.Close()

	id := primitive.NewObjectID()

	user := &models.User{
		Id:       id,
		Name:     "Test Name",
		Email:    fmt.Sprintf("%s@email.test", id.Hex()),
		Password: "password",
		Created:  time.Now(),
		Updated:  time.Now(),
	}

	r := NewUsersRepository(conn)
	err = r.Save(user)
	assert.NoError(t, err)

	users, err := r.GetAll()
	assert.NoError(t, err)
	assert.NotEmpty(t, users)
}
