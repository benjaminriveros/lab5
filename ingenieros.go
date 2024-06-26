package main

import (
	"context"
	"fmt"
	"log"

	pb "code/proto"

	"google.golang.org/grpc"
)

func enviarComandoBroken(c pb.GeneralClient, tipo string, ns string, nb string, valor string) bool {
	request := &pb.Command{
		Tipo:  tipo,
		Ns:    ns,
		Nb:    nb,
		Valor: valor,
	}

	// Llamar al servicio gRPC para enviar el comando
	resp, err := c.SendCommand(context.Background(), request)
	if err != nil {
		log.Fatalf("Error al enviar comando: %v", err)
		return false
	}
	fmt.Printf("broken indica enviar a ip: %s\n", resp.Ip)
	return true
}

func main() {
	// Conexión al servidor gRPC
	conn, err := grpc.Dial("10.35.169.11:50050", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Error al conectar: %v", err)
	}
	defer conn.Close()
	// Crear un cliente gRPC
	bkn := pb.NewGeneralClient(conn)

	// Terminal indica enviar mensaje
	enviarComandoBroken(bkn, "probando", "probando", "probando", "probando")
}
