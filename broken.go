// debe escuchar a los ingenieros en puerto 50050
package main

import (
	"context"
	"log"
	"math/rand"
	"net"
	"time"

	pb "code/proto"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedGeneralServer
}

func obtenerDireccionAleatoria() string {
	direcciones := []string{
		"10.35.169.12:50030",
		"10.35.169.13:50030",
		"10.35.169.14:50030",
	}

	// Inicializar la semilla aleatoria
	rand.Seed(time.Now().UnixNano())

	// Elegir aleatoriamente una dirección de la lista
	return direcciones[rand.Intn(len(direcciones))]
}

func (s *server) SendCommand(ctx context.Context, req *pb.Command) (*pb.Ip, error) {
	randomIp := obtenerDireccionAleatoria()
	response := &pb.Ip{Ip: randomIp}
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
	pb.RegisterGeneralServer(s, &server{})
	// Iniciar servidor
	log.Println("Servidor gRPC iniciado en localhost:50050")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Error al servir: %v", err)
	}

}
