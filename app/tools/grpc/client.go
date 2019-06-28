package main

import (
	hpb "github.com/mf-sakura/bh_hotel/app/proto"

	"context"
	"fmt"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50001", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	c := hpb.NewHotelServiceClient(conn)
	_, err = c.ReserveHotel(context.Background(), &hpb.ReserveHotellMessage{
		PlanId:     2,
		UserId:     1,
		SequenceId: "aaa",
	})
	if err != nil {
		fmt.Printf("error:%v", err)
	}
}
