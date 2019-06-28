package main

import (
	"github.com/mf-sakura/bh_hotel/app/config"
	"github.com/mf-sakura/bh_hotel/app/db"
	hpb "github.com/mf-sakura/bh_hotel/app/proto"
	"github.com/mf-sakura/bh_hotel/app/server"

	"google.golang.org/grpc"
	"net"
)

const (
	port = ":5001"
)

func main() {
	conf, err := config.LoadConifg()
	if err != nil {
		panic(err)
	}
	dsn, err := db.CreateDataSourceName(conf.Port, conf.Host, "bh_hotel", conf.User, conf.Password)
	if err != nil {
		panic(err)
	}
	if err := db.NewDB(dsn); err != nil {
		panic(err)
	}
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
