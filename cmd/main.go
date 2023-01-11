package main

import (
	"fmt"
	"log"
	"net"

	"github.com/Fachrulmustofa20/go-grpc-product-svc/pkg/config"
	"github.com/Fachrulmustofa20/go-grpc-product-svc/pkg/db"
	"github.com/Fachrulmustofa20/go-grpc-product-svc/pkg/pb"
	"github.com/Fachrulmustofa20/go-grpc-product-svc/pkg/services"
	"google.golang.org/grpc"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalln("Failed at config:", err)
	}

	h := db.Init(cfg.DBUrl)
	listen, err := net.Listen("tcp", cfg.Port)
	if err != nil {
		log.Fatalln("Failed to listing:", err)
	}

	fmt.Println("Product Svc on", cfg.Port)
	s := services.Server{
		H: h,
	}

	grpcServer := grpc.NewServer()
	pb.RegisterProductServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalln("Failed to serve:", err)
	}
}
