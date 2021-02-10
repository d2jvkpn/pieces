package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

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

	if err = doClientStreaming(client); err != nil {
		return
	}
}

func doServerStreaming(client CalculatorServiceClient) (err error) {
	fmt.Println(">>> Calculator RPC client doServerStreaming")
	var (
		stream CalculatorService_PNDClient
		res    *Number
	)

	req := &Number{Value: 1239039284}
	if stream, err = client.PND(context.TODO(), req); err != nil {
		return err
	}

	for {
		if res, err = stream.Recv(); err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		fmt.Println("    received:", res.Value)
	}

	return nil
}

func doClientStreaming(client CalculatorServiceClient) (err error) {
	fmt.Println(">>> Calculator RPC client doClientStreaming")
	var (
		stream CalculatorService_MultiplyClient
		res    *Number
	)

	if stream, err = client.Multiply(context.TODO()); err != nil {
		return err
	}

	for i := 0; i < 15; i++ {
		fmt.Println("    sending:", i+1)
		if err = stream.Send(&Number{Value: int64(i + 1)}); err != nil {
			return err
		}
		time.Sleep(time.Second)
	}

	if res, err = stream.CloseAndRecv(); err != nil {
		return err
	}

	fmt.Printf("%v\n", res.Value)
	return nil
}
