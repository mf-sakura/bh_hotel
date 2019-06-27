package main

import (
	hpb "github.com/mf-sakura/bh_hotel/app/proto"
	"github.com/mf-sakura/bh_hotel/app/server"

	"google.golang.org/grpc"
	"net"
)

const (
	port = ":5001"
)

func main() {
	listen, err := net.Listen("tcp", port)
	if err != nil {
		panic(err)
	}
	s := grpc.NewServer()

	defer func() {
		err := recover()
		s.GracefulStop()
		if err != nil {
			panic(err)
		}
	}()
	hpb.RegisterHotelServiceServer(s, &server.HotelServiceServerImpl{})

	if err := s.Serve(listen); err != nil {
		panic(err)
	}
}
