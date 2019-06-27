package db

import (
	"database/sql"
	"reflect"
	"testing"
)

func TestSelectReservationsFromUserID(t *testing.T) {

	PrepareFixture("testdata/reservation")

	type args struct {
		userID int64
	}
	tests := []struct {
		name    string
		args    args
		want    []*Reservation
		wantErr bool
	}{
		{
			name: "正常系",
			args: args{1},
			want: []*Reservation{
				&Reservation{
					ID:         1,
					UserID:     1,
					SequenceID: "aaa",
					PlanID:     3,
				},
			},
		},
		{
			name: "存在しないユーザー",
			args: args{100},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := SelectReservationsFromUserID(tt.args.userID)
			if (err != nil) != tt.wantErr {
				t.Errorf("SelectReservationsFromUserID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SelectReservationsFromUserID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReserveHotel(t *testing.T) {
	PrepareFixture("testdata/reservation")

	type args struct {
		reservation *Reservation
	}
	tests := []struct {
		name    string
		args    args
		want    int64
		wantErr bool
	}{
		{
			name: "正常系",
			args: args{
				&Reservation{
					UserID:     1,
					PlanID:     1,
					SequenceID: "ccc",
				},
			},
		},
		{
			name: "異常系 満室",
			args: args{
				&Reservation{
					UserID:     1,
					PlanID:     2,
					SequenceID: "ddd",
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := ReserveHotel(tt.args.reservation)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReserveHotel() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func Test_insertReservation(t *testing.T) {
	PrepareFixture("testdata/reservation")

	type args struct {
		reservation *Reservation
	}
	tests := []struct {
		name    string
		args    args
		want    sql.Result
		wantErr bool
	}{
		{
			name: "正常系(クエリチェック)",
			args: args{&Reservation{
				PlanID: 1,
			}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := insertReservation(tt.args.reservation)
			if (err != nil) != tt.wantErr {
				t.Errorf("insertReservation() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func Test_deleteReservation(t *testing.T) {
	PrepareFixture("testdata/reservation")

	type args struct {
		reservationID int64
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "idが存在する",
			args: args{1},
		},
		{
			name: "idが存在しない",
			args: args{100},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := deleteReservation(tt.args.reservationID); (err != nil) != tt.wantErr {
				t.Errorf("deleteReservation() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCancelHotel(t *testing.T) {
	PrepareFixture("testdata/reservation")

	type args struct {
		reservationID int64
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "正常系",
			args: args{1},
		},
		{
			name:    "異常系 存在しないID",
			args:    args{100},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := CancelHotel(tt.args.reservationID); (err != nil) != tt.wantErr {
				t.Errorf("CancelHotel() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGetReservationFromID(t *testing.T) {
	PrepareFixture("testdata/reservation")

	type args struct {
		reservationID int64
		forUpdate     bool
	}
	tests := []struct {
		name    string
		args    args
		want    *Reservation
		wantErr bool
	}{
		{
			name: "予約が存在",
			args: args{reservationID: 1},
		},
		{
			name:    "予約が存在しない",
			args:    args{reservationID: 100},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := GetReservationFromID(tt.args.reservationID, tt.args.forUpdate)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetReservationFromID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
