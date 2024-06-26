package main

import (
	"fmt"
	"log"
	"net"

	pb "code/proto"
)

const ipCoordinador = "10.35.169.13"

type server struct {
	pb.UnimplementedGeneralServer
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
	/*
		lis, err := net.Listen("tcp", ":50051")
		if err != nil {
			log.Fatalf("Error al crear el servidor: %v", err)
		}
		s := grpc.NewServer()*/
	// Registrar el servicio Greeter en el servidor
	//pb.RegisterGreeterServer(s, &server{})

	// Ejecutar si es que es servidor coordinador
	if ipLocal == ipCoordinador {
		fmt.Println("Este servidor es el coordinador. Iniciando propagación de cambios...")
		// iniciar go rutina con: "go servidor.iniciarPropagacion()"
	} else {
		fmt.Println("Este servidor no es el coordinador.")
	}
	// Iniciar el servidor
	/*
		log.Println("Servidor gRPC iniciado en localhost:50051")
		if err := s.Serve(lis); err != nil {
			log.Fatalf("Error al servir: %v", err)
		}*/
}
