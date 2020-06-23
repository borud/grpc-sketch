package main

import (
	"log"
	"math/rand"
	"net"

	pb "github.com/borud/tls-sketch/pkg/sketchpb"
	"google.golang.org/grpc"
)

type sketchServer struct{}

func (s *sketchServer) GetRandomNumbers(request *pb.Request, stream pb.Sketch_GetRandomNumbersServer) error {
	for i := uint64(0); i < request.NumRandomNumbers; i++ {
		err := stream.Send(&pb.Response{RandomNumber: rand.Uint64()})
		if err != nil {
			log.Printf("Error streaming : %v", err)
			return err
		}
	}

	return nil
}

func main() {
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatalf("Error listening to port: %v", err)
	}

	log.Printf("Listening to port %s", listener.Addr().String())

	grpcServer := grpc.NewServer()
	pb.RegisterSketchServer(grpcServer, &sketchServer{})
	grpcServer.Serve(listener)
}
