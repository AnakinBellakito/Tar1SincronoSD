package main

import (
	"containerized-go-central/proto"
	"context"
	"fmt"
	"math/rand"
	"net"
	"os"
	"strconv"
	"strings"
	"time"

	grpc "google.golang.org/grpc"
)

var sedeo = time.Now().UTC().UnixNano()
var init_keys = calculateInitKeys()
var currentKeys, remKeys int

type server struct {
	proto.UnimplementedGreeterServer
}

func generateId() int {
	var sample int = 100
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(sample)
}

func (s *server) Notify(ctx context.Context, req *proto.EmptyRequest) (*proto.HelloReply, error) {
	fmt.Println("Starting to send message NTF")

	return &proto.HelloReply{
		Response: int32(init_keys),
	}, nil
}

func (s *server) QuantKeys(ctx context.Context, req *proto.EmptyRequest) (*proto.HelloReply, error) {

	fmt.Println("Starting to send message QK")

	return &proto.HelloReply{
		Response: int32(100),
	}, nil
}
func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	//fmt.Print(string(txt))
	var cantInteraciones int

	listener, err := net.Listen("tcp", ":8090")
	check(err)
	serv := grpc.NewServer()
	proto.RegisterGreeterServer(serv, &server{})
	serv.Serve(listener)

	fmt.Print("Type a number for Iteraciones: \n")
	fmt.Scan(&cantInteraciones)
	fmt.Println(string(cantInteraciones))

	for i := 1; i <= cantInteraciones; i++ {
		fmt.Println("Generacion: %d/%d", i, cantInteraciones)

	}

	//esscribir #Llaves/Servidor/Timestamp
	//esscribir #solicitantes/Exito/fracaso <- guardar y enviar a los regionales

	//definir que pasa en los distintos flujos
}

func calculateInitKeys() int {

	rand.Seed(sedeo)
	var high int
	txt, err := os.ReadFile("./parametro_init.txt")
	check(err)
	firstLine := string(txt)
	primeraLine := strings.Split(firstLine, "-")

	for pos, v := range primeraLine {
		fmt.Printf("Pos: %d -> Value:%s\n", pos, v)
		if pos == 1 {
			high, _ = strconv.Atoi(v)
		}
	}
	val := rand.Intn(high + 1)

	return val
}
