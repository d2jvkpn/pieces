package main

import (
	"context"
	"fmt"
	"io"
	"log"

	. "x/pkg/greetpb"

	"google.golang.org/grpc"
)

func main() {
	var (
		err    error
		conn   *grpc.ClientConn
		client GreetServiceClient
	)

	if conn, err = grpc.Dial("localhost:50051", grpc.WithInsecure()); err != nil {
		log.Fatal(err)
	}

	defer func() {
		conn.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	client = NewGreetServiceClient(conn)
	if err = doUnary(client); err != nil {
		return
	}

	if err = doServerStreaming(client); err != nil {
		return
	}
}

func doUnary(client GreetServiceClient) (err error) {
	fmt.Println(">>> Greet RPC client doUnary")
	var (
		res *GreetResponse
	)

	req := &GreetRequest{Greeting: &Greeting{
		FirstName: "Rover",
		LastName:  "Chan",
	}}

	if res, err = client.Greet(context.TODO(), req); err != nil {
		return err
	}
	fmt.Println("    received:", res.Result)

	return nil
}

func doServerStreaming(client GreetServiceClient) (err error) {
	fmt.Println(">>> Greet RPC client doServerStreaming")

	var (
		stream GreetService_Greet2Client
		res    *Greet2Response
	)

	req := &GreetRequest{Greeting: &Greeting{
		FirstName: "Rover",
		LastName:  "Chan",
	}}

	if stream, err = client.Greet2(context.TODO(), req); err != nil {
		return err
	}

	for {
		if res, err = stream.Recv(); err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		fmt.Println("    received:", res.Result)
	}

	return nil
}
