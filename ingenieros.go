package main

import (
	"context"
	"fmt"
	"log"

	pb "code/proto"

	"google.golang.org/grpc"
)

func main() {
	// Conexión al servidor gRPC
	conn, err := grpc.Dial("10.35.169.12:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Error al conectar: %v", err)
	}
	defer conn.Close()

	// Crear un cliente gRPC
	c := pb.NewGreeterClient(conn)

	// Llamar al servicio SayHello
	response, err := c.SayHello(context.Background(), &pb.HelloRequest{Name: "Mundo"})
	if err != nil {
		log.Fatalf("Error al llamar al servicio SayHello: %v", err)
	}

	// Imprimir la respuesta del servidor2
	fmt.Println("Respuesta del servidor:", response.Message)
}
