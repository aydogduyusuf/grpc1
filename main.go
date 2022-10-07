package main

import (
	"fmt"
	"grpc1/business"
	"grpc1/util"
	"log"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config", err)
	}

	avalancheFuji := "https://api.avax-test.network/ext/bc/C/rpc"
	client := business.InitNetwork(avalancheFuji)
	fmt.Printf("client: %v\n", client)

	store := db.NewStore(conn)
	rungRPCServer(config,store)

}

func rungRPCServer(config util.Config, store db.Store) {
	server, err := gapi.NewServer(config, store)
	if err != nil {
		log.Fatal("can not create server", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterDBchainServer(grpcServer, server)
	reflection.Register(grpcServer)

	listener, err := net.Listen("tcp", config.GRPCServerAddress)
	if err != nil {
		log.Fatal("cannot create listener", err)
	}
	log.Printf("start gRPC server at %s", listener.Addr().String())
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal("cannot start gRPC server", err)
	}
}