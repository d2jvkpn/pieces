package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	. "x/pkg/greetpb"

	"google.golang.org/grpc"
)

type Server struct{}

func (srv *Server) Greet(ctx context.Context, req *GreetRequest) (res *GreetResponse, err error) {
	res = new(GreetResponse)
	// res.Result = fmt.Sprintf("Hello, %s %s!", req.Greeting.FirstName, req.Greeting.LastName)
	firstName := req.GetGreeting().GetFirstName()
	lastName := req.GetGreeting().GetLastName()

	res.Result = fmt.Sprintf("Hello, %s %s!", firstName, lastName)

	return
}

func (srv *Server) Greet2(req *GreetRequest, stream GreetService_Greet2Server) (err error) {
	firstName := req.GetGreeting().GetFirstName()
	lastName := req.GetGreeting().GetLastName()

	for i := 0; i < 10; i++ {
		msg := fmt.Sprintf("Hello, %s %s, number: %d.", firstName, lastName, i)
		fmt.Printf(">> Sending: %q\n", msg)

		if err = stream.Send(&Greet2Response{Result: msg}); err != nil {
			break
		}
		time.Sleep(time.Second)
	}

	return err
}

func main() {
	var (
		addr string
		err  error
		lis  net.Listener
		srv  *grpc.Server
	)

	addr = ":50051"
	if lis, err = net.Listen("tcp", addr); err != nil {
		log.Fatal(err)
	}

	srv = grpc.NewServer()
	RegisterGreetServiceServer(srv, &Server{})

	log.Printf("Greet RPC server %q\n", addr)
	if err = srv.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
