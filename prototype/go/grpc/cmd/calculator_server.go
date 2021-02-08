package main

import (
	"fmt"
	"log"
	"net"

	. "x/pkg/calculatorpb"

	"google.golang.org/grpc"
)

type Server struct{}

func (srv *Server) PND(req *PNDRequest, stream CalculatorService_PNDServer) (err error) {
	number := req.GetNumber()
	divisor := int64(2)
	fmt.Println(">>> PND number:", number)

	for number > 1 {
		if number%divisor == 0 {
			stream.Send(&PNDResponse{PrimeFactor: divisor})
			number = number / divisor
		} else {
			divisor++
		}
	}

	return nil
}

func main() {
	var (
		addr string
		err  error
		lis  net.Listener
		srv  *grpc.Server
	)

	addr = ":50052"
	if lis, err = net.Listen("tcp", addr); err != nil {
		log.Fatal(err)
	}

	srv = grpc.NewServer()
	RegisterCalculatorServiceServer(srv, &Server{})

	log.Printf("Calculator RPC server %q\n", addr)
	if err = srv.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
