package main

import (
	"context"
	"io"
	"log"

	pb "github.com/borud/tls-sketch/pkg/sketchpb"
	"google.golang.org/grpc"
)

const serverAddr = "localhost:1234"

func main() {
	conn, err := grpc.Dial(serverAddr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Error connecting to server %s: %v", serverAddr, err)
	}
	defer conn.Close()

	client := pb.NewSketchClient(conn)

	stream, err := client.GetRandomNumbers(context.Background(), &pb.Request{NumRandomNumbers: 10})
	if err != nil {
		log.Fatalf("Error calling GetRandomNumbers(): %v", err)
	}

	for {
		response, err := stream.Recv()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("%v.GetRandomNumbers(_) = _, %v", client, err)
		}
		log.Printf("Random number returned : %d", response.RandomNumber)
	}

}
