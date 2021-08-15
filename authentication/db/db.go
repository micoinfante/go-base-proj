package db

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"time"
)

type Connection interface {
	Close(c *mongo.Client, ctx context.Context, cancel context.CancelFunc)
	Connect(cfg Config) (*mongo.Client, context.Context, context.CancelFunc, error)
	DBClient() *mongo.Database
}

type conn struct {
	database *mongo.Database
	Client *mongo.Client
	Ctx    context.Context
	Cancel context.CancelFunc
}

func Close(c *mongo.Client, ctx context.Context, cancel context.CancelFunc) {
	//cancel context
	defer cancel()

	defer func() {
		if err := c.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()
}

func Connect(cfg Config) (*conn, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	client, _ := mongo.Connect(ctx, options.Client().ApplyURI(cfg.Dsn()))

	// test connection

	if err := client.Ping(context.Background(), readpref.Primary()); err != nil {

	} else {
		fmt.Printf("Connected successfully to: %s", cfg.Dsn())
	}

	return &conn{Client: client, database: client.Database(cfg.DbName()), Ctx: ctx, Cancel: cancel}, nil
}

func DB(c *conn) *mongo.Database {
	return c.database
}
