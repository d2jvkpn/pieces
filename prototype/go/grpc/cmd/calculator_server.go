package main

import (
	"fmt"
	"io"
	"log"
	"net"

	. "x/pkg/calculatorpb"

	"google.golang.org/grpc"
)

type Server struct{}

func (srv *Server) PND(req *Number, stream CalculatorService_PNDServer) (err error) {
	fmt.Println(">>> PND processing")
	number := req.GetValue()
	divisor := int64(2)
	fmt.Println("    received:", number)

	for number > 1 {
		if number%divisor == 0 {
			stream.Send(&Number{Value: divisor})
			number = number / divisor
		} else {
			divisor++
		}
	}

	return nil
}

func (srv *Server) Multiply(stream CalculatorService_MultiplyServer) (err error) {
	fmt.Println(">>> Multiply processing")
	var req, res *Number

	res = &Number{Value: 1}
	for {
		if req, err = stream.Recv(); err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		fmt.Println("    received:", req.Value)
		res.Value *= req.Value
	}

	return stream.SendAndClose(res)
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

	log.Printf("### Calculator RPC server %q\n", addr)
	if err = srv.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
