package server

import (
	"context"
	"github.com/opentracing/opentracing-go"

	"github.com/mf-sakura/bh_hotel/app/db"
	hpb "github.com/mf-sakura/bh_hotel/app/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

type HotelServiceServerImpl struct {
}

func (h *HotelServiceServerImpl) GetHotels(ctx context.Context, req *hpb.GetHotelsMessage) (*hpb.GetHotelsResponse, error) {
	return nil, nil
}

func (h *HotelServiceServerImpl) GetPlans(ctx context.Context, req *hpb.GetPlansMessage) (*hpb.GetPlansResponse, error) {
	return nil, nil
}
func (h *HotelServiceServerImpl) ReserveHotel(ctx context.Context, req *hpb.ReserveHotellMessage) (*hpb.ReserveHotelResponse, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "reserve_hotel")
	defer span.Finish()
	id, err := db.ReserveHotel(&db.Reservation{
		UserID:     req.UserId,
		PlanID:     req.PlanId,
		SequenceID: req.SequenceId,
	})
	if err != nil {
		return nil, grpc.Errorf(codes.Internal, "db.ReserveHotel failed:%v", err)
	}
	return &hpb.ReserveHotelResponse{
		ReservationId: id,
	}, nil
}
func (h *HotelServiceServerImpl) CancelHotel(ctx context.Context, req *hpb.CancelHotelMessage) (*hpb.CancelHotelResponse, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "canncel_hotel")
	defer span.Finish()
	if err := db.CancelHotel(req.ReservationId); err != nil {
		return nil, grpc.Errorf(codes.Internal, "db.CancelHotel failed:%v", err)
	}
	return &hpb.CancelHotelResponse{
		Restult: true,
	}, nil
}
func (h *HotelServiceServerImpl) GetReservations(ctx context.Context, req *hpb.GetReservationsMessage) (*hpb.GetReservationsResponse, error) {
	return nil, nil
}
