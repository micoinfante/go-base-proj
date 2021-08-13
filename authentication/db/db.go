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
	Connect(uri string) (*mongo.Client, context.Context, context.CancelFunc, error)
	DB() *mongo.Database
}

type conn struct {
	session  *mongo.Session
	database *mongo.Database
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

func Connect(uri string) (*mongo.Client, context.Context, context.CancelFunc, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))

	// test connection

	if err := client.Ping(context.Background(), readpref.Primary()); err != nil {
		return client, ctx, cancel, err
	} else {
		fmt.Printf("Connected successfully to: %s", uri)
	}
	return client, ctx, cancel, err
}

// parameter - return type
func DB(conn *conn) *mongo.Database {
	return conn.database
}
