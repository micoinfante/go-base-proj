package main

import (
	"authentication/api/handlers"
	"authentication/api/routes"
	"authentication/pb"
	"flag"
	"fmt"
	"github.com/gorilla/mux"
	"google.golang.org/grpc"
	"log"
	"net/http"
)

var (
	port     int
	authAddr string
)

func init() {
	flag.IntVar(&port, "port", 9080, "api service port")
	flag.StringVar(&authAddr, "auth_addr", "localhost:9001", "authentication service address")
	flag.Parse()
}

func main() {
	conn, err := grpc.Dial(authAddr, grpc.WithInsecure())
	if err != nil {
		log.Panicln(err)
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {

		}
	}(conn)

	authSvcClient := pb.NewAuthServiceClient(conn)
	authHandlers := handlers.NewAuthServiceClient(authSvcClient)
	authRoutes := routes.NewAuthRoute(authHandlers)

	router := mux.NewRouter().StrictSlash(true)
	routes.Install(router, authRoutes)

	log.Printf("API service running on [::]:%d", port)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), routes.WithCORS(router)))
}
