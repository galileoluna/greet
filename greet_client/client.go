package main

import (
	"fmt"
	"log"

	"github.com/galileoluna/greet/greetpb"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Soy el cliente")
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil {
		log.Fatal("no se pudo conectar al puerto %v", err)
	}
	defer cc.Close()

	c := greetpb.NewGreetServiceClient(cc)
	log.Printf("Cliente creado %f", c)
}
