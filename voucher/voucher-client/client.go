package main

import (
	"context"
	"fmt"
	"io"
	"log"

	"github.com/ianprogrammer/golang-ifood-dev/voucher/voucherpb"
	"google.golang.org/grpc"
)

func main() {
	println("GRPC Client")

	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Couldn't connect with host, cause: %v", err)
	}
	defer conn.Close()

	c := voucherpb.NewVoucherServiceClient(conn)

	//unaryCall(c)
	streamingCall(c)

}

func streamingCall(c voucherpb.VoucherServiceClient) {
	fmt.Println("Streaming Server Call")
	req := &voucherpb.VoucherStreamRequest{
		CampaignId: "voucher15",
	}

	stream, err := c.VoucherStream(context.Background(), req)

	if err != nil {
		log.Fatalf("Error while calling streaming voucher GRPC %v", err)
	}

	for {
		msg, err := stream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Error while reading the stream %v", err)
		}
		log.Println("Response", msg)
	}

}

func unaryCall(c voucherpb.VoucherServiceClient) {
	fmt.Print("Unary Call")
	req := &voucherpb.VoucherRequest{
		Voucher: &voucherpb.Voucher{
			CustomerId: "customer_id",
			OrderId:    "order_id",
			CampaignId: "voucher15",
		},
	}

	res, err := c.Voucher(context.Background(), req)

	if err != nil {
		log.Fatalf("Error while calling voucher GRPC %v", err)
	}
	fmt.Printf("Response %v", res)

}
