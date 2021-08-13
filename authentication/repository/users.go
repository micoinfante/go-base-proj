package repository

import (
	"authentication/db"
	"authentication/models"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

const UsersCollection = "users"

type UsersRepository interface {
	Save(user *models.User) error
	GetById(id string) (user *models.User, err error)
	GetByEmail(email string) (user *models.User, err error)
	GetAll(email string) (users []*models.User, err error)
	Update(user *models.User) error
	Delete(user *models.User) error
}

type usersRepository struct {
	collection *mongo.Collection
}

func NewUsersRepository(conn db.Connection) UsersRepository {
	return &usersRepository{collection: conn.DB().Collection(UsersCollection)}
}

func (r *usersRepository) Save(user *models.User) error {
	result, err := r.collection.InsertOne(context.Background(), user)
	if err != nil {
		return err
	}
	fmt.Printf("Write 1 data: %s", result)
	return nil
}

func (r *usersRepository) GetById(id string) (user *models.User, err error) {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Println("Invalid id")
	}
	if err := r.collection.FindOne(context.Background(), bson.M{"_id": objectId}).Decode(&user); err != nil {
		return nil, err
	}

	return user, nil
}

func (r *usersRepository) GetByEmail(email string) (user *models.User, err error) {
	if err := r.collection.FindOne(context.Background(), bson.M{"email": email}).Decode(&user); err != nil {
		return nil, err
	}
	return user, nil
}

func (r *usersRepository) GetAll(email string) (users []*models.User, err error) {
	result, err := r.collection.Find(context.Background(), bson.M{})

	if err != nil {
		return nil, err
	}

	if err := result.Decode(&users); err != nil {
		return nil, err
	}

	return users, nil
}

func (r *usersRepository) Update(user *models.User) error {
	_, err := r.collection.UpdateOne(context.Background(), bson.M{"_id": user.Id}, &user)
	return err
}

func (r *usersRepository) Delete(user *models.User) error {
	_, err := r.collection.DeleteOne(context.Background(), bson.M{"_id": user.Id})
	return err
}
