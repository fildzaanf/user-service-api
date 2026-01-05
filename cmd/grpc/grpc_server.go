package main

import (
	"fmt"
	"log"
	"net"

	"user-service-api/infrastructure/config"
	"user-service-api/infrastructure/database"
	userGRPC "user-service-api/internal/user/adapter/handler/grpc"
	"user-service-api/pkg/middleware"

	"github.com/common-nighthawk/go-figure"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	middleware.InitLogger()

	godotenv.Load()
	cfg, err := config.LoadConfig()
	if err != nil {
		logrus.Fatalf("[ERROR] failed to load configuration: %v", err)
	}

	psql := database.ConnectPostgreSQL(false)

	server := grpc.NewServer(
		grpc.UnaryInterceptor(middleware.JWTUnaryInterceptor()),
	)

	userGRPC.RegisterUserServices(server, psql)
	reflection.Register(server) // enable reflection for grpcurl and other tools

	address := cfg.USERSERVICE.USER_GRPC_HOST + ":" + cfg.USERSERVICE.USER_GRPC_PORT
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	fig := figure.NewFigure("USER SERVICE API", "small", true)
	fig.Print()

	fmt.Printf("\n📡 Listening on %s\n", address)

	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
