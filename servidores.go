package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"

	pb "code/proto"

	"google.golang.org/grpc"
)

const ipCoordinador = "10.35.169.13"

var idServidor int
var fileDict = make(map[string][]int32)

type server struct {
	pb.UnimplementedGeneralServer
}

// Función para agregar datos al diccionario
func dicVectData(filename string, data []int32) {
	fileDict[filename] = data
}

// Función para obtener datos del diccionario
func getVecData(filename string) ([]int32, bool) {
	data, exists := fileDict[filename]
	return data, exists
}

// Metodo grpc
func (s *server) RegisterCommand(ctx context.Context, req *pb.Command) (*pb.Vector, error) {
	// Crear o abrir el archivo
	filename := "txts/" + req.Ns + ".txt"
	file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		fmt.Println("Error al abrir/crear el archivo:", err)
		return nil, err // Deberías devolver un error en caso de fallo
	}
	defer file.Close()
	line := req.Tipo + " " + req.Ns + " " + req.Nb
	if req.Tipo != "borrar base" {
		line = line + " " + req.Valor
	}
	// Escribir en el archivo
	_, err = file.WriteString(line + "\n") // Asegúrate de agregar un salto de línea al final
	if err != nil {
		fmt.Println("Error al escribir en el archivo:", err)
		return nil, err // Devuelve un error si ocurre algún problema al escribir
	}
	// Actualizar o inicializar el vector en memoria
	vector, exists := getVecData(filename)
	if !exists {
		// Inicializar vector si no existe
		vector = []int32{0, 0, 0}
	}
	vector[idServidor]++ // Incrementar el contador para el servidor actual
	// Guardar o actualizar el vector en memoria
	fmt.Printf("nuevo vector de  %s en SvId:%d, [%d, %d, %d]\n", filename, idServidor, vector[0], vector[1], vector[2])
	dicVectData(filename, vector)
	// Preparar la respuesta a retornar
	retorno := &pb.Vector{
		Sv1: vector[0],
		Sv2: vector[1],
		Sv3: vector[2],
	}
	// Retornar la respuesta y nil como error (si no hay errores específicos)

	return retorno, nil
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
	ultimoDigito := int(ipLocal[len(ipLocal)-1] - '0')

	// Restarle 2 al último dígito
	idServidor = ultimoDigito - 2
	fmt.Printf("IP local del servidor:%s, con id %d\n", ipLocal, idServidor)

	// Crear un servidor gRPC
	lis, err := net.Listen("tcp", ":50030")
	if err != nil {
		log.Fatalf("Error al crear el servidor: %v", err)
	}
	s := grpc.NewServer()
	// Registrar el servicio General en el servidor
	pb.RegisterGeneralServer(s, &server{})
	// Iniciar el servidor
	log.Println("Servidor gRPC iniciado en localhost:50030")
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
