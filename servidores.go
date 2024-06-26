package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "code/proto"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
)

const ipCoordinador = "10.35.169.13"

type server struct {
	pb.UnimplementedGeneralServer
}

func (s *server) RegisterCommand(ctx context.Context, req *pb.Command) (*empty.Empty, error) {
	// Lógica de registro del comando aquí
	fmt.Printf("registrar en este servidor mensaje: %s\n", req.Ns)
	// Si no hay errores específicos, retornar nil como error
	return &empty.Empty{}, nil
}

func main() {

	// Obtener la dirección IP local
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		log.Fatal("Error al obtener la dirección IP:", err)
	}
	var ipLocal string
	for _, addr := range addrs {
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				ipLocal = ipnet.IP.String()
				break
			}
		}
	}
	if ipLocal == "" {
		log.Fatal("No se pudo determinar la dirección IP local")
	}
	fmt.Println("IP local del servidor:", ipLocal)

	// Crear un servidor gRPC
	lis, err := net.Listen("tcp", ":50030")
	if err != nil {
		log.Fatalf("Error al crear el servidor: %v", err)
	}
	s := grpc.NewServer()
	// Registrar el servicio General en el servidor
	pb.RegisterGeneralServer(s, &server{})
	// Iniciar el servidor
	log.Println("Servidor gRPC iniciado en localhost:50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Error al servir: %v", err)
	}

	/*
		// Ejecutar si es que es servidor coordinador
		if ipLocal == ipCoordinador {
			fmt.Println("Este servidor es el coordinador. Iniciando propagación de cambios...")
			// iniciar go rutina con: "go servidor.iniciarPropagacion()"
		} else {
			fmt.Println("Este servidor no es el coordinador.")
		}*/
}
