package main

import (
	"authentication/db"
	"authentication/pb"
	"authentication/repository"
	"authentication/services"
	"flag"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"log"
	"net"
)

var (
	local bool
	port  int
)

func init() {
	flag.IntVar(&port, "port", 9001, "authentication service port")
	flag.BoolVar(&local, "local", true, "run authentication service local")
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

	usersRepository := repository.NewUsersRepository(conn)
	authService := services.NewAuthService(usersRepository)

	listener, err := net.Listen("tcp", ":9001")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterAuthServiceServer(grpcServer, authService)

	log.Printf("\nauthentication service running on [::]:%d\n\n", port)

	grpcServer.Serve(listener)
}
