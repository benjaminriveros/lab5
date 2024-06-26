// debe escuchar a los ingenieros en puerto 50050
package main

import (
	"context"
	"log"
	"net"

	pb "code/proto"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedGeneralServer
}

func (s *server) SendCommand(ctx context.Context, req *pb.Command) (*pb.Ip, error) {
	response := &pb.Ip{Ip: ".13"}
	return response, nil
}

func main() {

	// Crear un servidor gRPC
	lis, err := net.Listen("tcp", ":50050")
	if err != nil {
		log.Fatalf("Error al crear el servidor: %v", err)
	}
	s := grpc.NewServer()
	// Registrar el servicio Greeter en el servidor
	pb.RegisterGeneralServer()
	// Iniciar servidor
	log.Println("Servidor gRPC iniciado en localhost:50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Error al servir: %v", err)
	}

}
