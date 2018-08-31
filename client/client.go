package main

import (
	"log"
	"time"

	pb "github.com/whatsadebugger/grpc-examples/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	address = "localhost:8080"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	no(err)
	defer conn.Close()

	c := pb.NewPhoneTrackerClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	phone := &pb.PhoneReq{
		Imei:       "666585287256616",
		Modelname:  "Samsung Galaxy Note 4",
		Unlockcode: "44444444",
	}
	resp, err := c.AddPhone(ctx, phone)
	log.Printf("Imei: %+v ModelName: %+v UnlockCode: %+v\n", resp.Imei, resp.Modelname, resp.Unlockcode)
}

func no(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
