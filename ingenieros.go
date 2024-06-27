package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"

	pb "code/proto"

	"google.golang.org/grpc"
)

var registro = make([]string, 0, 20)

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
	fmt.Printf("El comando se enviará al servidor en %s\n", resp.Ip)

	// Iniciar conexión con servior en resp.Ip
	conn, err := grpc.Dial(resp.Ip, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Error al conectar: %v", err)
	}
	defer conn.Close()
	// Crear un cliente gRPC
	sf := pb.NewGeneralClient(conn)
	// Enviar comando a Servidor Fulcrum
	vector, err := sf.RegisterCommand(context.Background(), request)
	if err != nil {
		log.Fatalf("Error al enviar comando: %v", err)
		return false
	}
	fmt.Printf("Servidor %s envía vector [%d, %d, %d] de archivo \"%s.txt\"\n", resp.Ip, vector.Sv1, vector.Sv2, vector.Sv3, ns)
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

	// Iniciar interfaz con usuario
	scanner := bufio.NewScanner(os.Stdin)
	continuar := true
	for continuar {
		fmt.Println("Digite acción a realizar\n")
		fmt.Println("'1' para enviar comando\n")
		fmt.Println("'2' para finalizar\n>>")
		scanner.Scan()
		input := scanner.Text()
		if input == "1" {
			var tipo, nomsec, nombas, valor string
			fmt.Println("\nComando (1:Agregar Base, 2:Renombrar Base, 3:Actualizar Valor o Borrar Base): ")
			fmt.Scan(&tipo)
			fmt.Println("\nNombre del sector: ")
			fmt.Scan(&nomsec)
			fmt.Println("\nNombre de la base: ")
			fmt.Scan(&nombas)
			fmt.Println("\nValor/Nuevo nombre/Nuevo valor/0: ")
			fmt.Scan(&valor)
			enviarComandoBroken(bkn, tipo, nomsec, nombas, valor)
			fmt.Println("\n Comando enviado. ")
		} else if input == "2" {
			continuar = false
		} else {
			fmt.Println("Digito fuera de rango\n")
		}
	}
}
