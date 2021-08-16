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
	Close()
	DB() *mongo.Database
}

type conn struct {
	database *mongo.Database
	Client   *mongo.Client
	Ctx      context.Context
	Cancel   context.CancelFunc
}

func NewConnection(cfg Config) (Connection, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	client, _ := mongo.Connect(ctx, options.Client().ApplyURI(cfg.Dsn()))

	// test connection

	if err := client.Ping(context.Background(), readpref.Primary()); err != nil {

	} else {
		fmt.Printf("Connected successfully to: %s", cfg.Dsn())
	}

	return &conn{Client: client, database: client.Database(cfg.DbName()), Ctx: ctx, Cancel: cancel}, nil
}

func (conn *conn) DB() *mongo.Database {
	return conn.database
}

func (conn *conn) Close() {
	//cancel context
	defer conn.Cancel()

	defer func() {
		if err := conn.Client.Disconnect(conn.Ctx); err != nil {
			panic(err)
		}
	}()
}
