package main

import (
	"log"
	"net"

	pb "github.com/whatsadebugger/grpc-examples/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type server struct{}

func (s *server) AddPhone(ctx context.Context, in *pb.PhoneReq) (*pb.PhoneResp, error) {
	return &pb.PhoneResp{
		Imei:       in.Imei,
		Modelname:  in.Modelname,
		Unlockcode: in.Unlockcode,
	}, nil
}

func main() {
	l, err := net.Listen("tcp", ":8080")
	no(err)
	s := grpc.NewServer()
	pb.RegisterPhoneTrackerServer(s, &server{})
	no(s.Serve(l))
}

func no(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
