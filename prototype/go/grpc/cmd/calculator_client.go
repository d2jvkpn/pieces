package main

import (
	"context"
	"fmt"
	"io"
	"log"

	. "x/pkg/calculatorpb"

	"google.golang.org/grpc"
)

func main() {
	var (
		err    error
		conn   *grpc.ClientConn
		client CalculatorServiceClient
	)

	if conn, err = grpc.Dial("localhost:50052", grpc.WithInsecure()); err != nil {
		log.Fatal(err)
		return
	}

	defer func() {
		conn.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	client = NewCalculatorServiceClient(conn)

	if err = doServerStreaming(client); err != nil {
		return
	}
}

func doServerStreaming(client CalculatorServiceClient) (err error) {
	fmt.Println(">>> Calculator RPC client doServerStreaming")
	var (
		resStream CalculatorService_PNDClient
		result    *PNDResponse
	)

	req := &PNDRequest{Number: 1239039284}
	if resStream, err = client.PND(context.TODO(), req); err != nil {
		return err
	}

	for {
		if result, err = resStream.Recv(); err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		fmt.Println(result)
	}

	return nil
}
