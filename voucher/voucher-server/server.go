package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"strconv"
	"time"

	"github.com/ianprogrammer/golang-ifood-dev/voucher/voucherpb"
	"google.golang.org/grpc"
)

type server struct{}

func (*server) Voucher(ctx context.Context, req *voucherpb.VoucherRequest) (*voucherpb.VoucherResponse, error) {

	fmt.Printf("GRPC unary call called with : %v", req)
	var value int32
	switch req.Voucher.CampaignId {
	case "voucher15":
		value = 15
	case "voucher30":
		value = 30
	default:
		value = 5
	}

	return &voucherpb.VoucherResponse{
		Value:      value,
		CustomerId: req.Voucher.CustomerId,
	}, nil

}

func (*server) VoucherStream(req *voucherpb.VoucherStreamRequest, stream voucherpb.VoucherService_VoucherStreamServer) error {

	println("Streaming RPC Call")
	var value int32
	switch req.CampaignId {
	case "voucher15":
		value = 15
	case "voucher30":
		value = 30
	default:
		value = 5
	}

	for i := 0; i < 10; i++ {
		res := &voucherpb.VoucherStreamResponse{
			Value:      value,
			CustomerId: "customer-" + strconv.Itoa(i),
		}
		stream.Send(res)
		time.Sleep(1000 * time.Millisecond)
	}
	return nil

}

func main() {
	println("Server is running")

	listener, err := net.Listen("tcp", "0.0.0.0:50051")

	if err != nil {
		log.Fatalf("Failed to listen, cause : %v", err)
	}

	s := grpc.NewServer()

	voucherpb.RegisterVoucherServiceServer(s, &server{})

	if err := s.Serve(listener); err != nil {
		log.Fatalf("Failed to serve, reason : %v", err)
	}
}
