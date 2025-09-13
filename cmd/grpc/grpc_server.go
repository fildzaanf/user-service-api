package main

import (
	"log"
	"net"

	"user-service-api/infrastructure/config"
	"user-service-api/infrastructure/database"
	usergrpc "user-service-api/internal/user/adapter/handler/grpc"
	"user-service-api/pkg/middleware"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {

	middleware.InitLogger()

	godotenv.Load()
	config, err := config.LoadConfig()
	if err != nil {
		logrus.Fatalf("[ERROR] failed to load configuration: %v", err)
	}

	psql := database.ConnectPostgreSQL(false)

	server := grpc.NewServer(
		grpc.UnaryInterceptor(middleware.JWTUnaryInterceptor()),
	)

	usergrpc.RegisterUserServices(server, psql)

	reflection.Register(server)

	address := config.GRPC.GRPC_HOST + ":" + config.GRPC.GRPC_PORT
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	log.Printf("gRPC server running on %s", address)
	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
