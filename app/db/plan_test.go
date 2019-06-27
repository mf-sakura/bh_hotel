package db

import (
	"testing"
	"time"
)

func TestInsertPlan(t *testing.T) {
	PrepareFixture("testdata/plan")

	type args struct {
		plan *Plan
	}
	tests := []struct {
		name    string
		args    args
		want    int64
		wantErr bool
	}{
		{
			name: "正常系(クエリチェック)",
			args: args{&Plan{Date: time.Now(), HotelID: 1}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := InsertPlan(tt.args.plan)
			if (err != nil) != tt.wantErr {
				t.Errorf("InsertPlan() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

		})
	}
}

func TestGetPlanFromID(t *testing.T) {
	PrepareFixture("testdata/plan")

	type args struct {
		planID    int64
		forUpdate bool
	}
	tests := []struct {
		name    string
		args    args
		want    *Plan
		wantErr bool
	}{
		{
			name: "プランが存在",
			args: args{planID: 1},
		},
		{
			name:    "プランが存在しない",
			args:    args{planID: 100},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := GetPlanFromID(tt.args.planID, tt.args.forUpdate)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetPlanFromID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

		})
	}
}

func Test_decrementPlanAvailable(t *testing.T) {
	PrepareFixture("testdata/plan")
	type args struct {
		planID int64
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
			name:    "異常系　満員(型エラー)",
			args:    args{2},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := decrementPlanAvailable(tt.args.planID); (err != nil) != tt.wantErr {
				t.Errorf("decrementPlanAvailable() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_incrementPlanAvailable(t *testing.T) {
	PrepareFixture("testdata/plan")

	type args struct {
		planID int64
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "正常系(クエリ確認)",
			args: args{2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := incrementPlanAvailable(tt.args.planID); (err != nil) != tt.wantErr {
				t.Errorf("incrementPlanAvailable() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSelectPlansFromHotelID(t *testing.T) {
	PrepareFixture("testdata/plan")

	type args struct {
		hotelID int64
	}
	tests := []struct {
		name      string
		args      args
		wantCount int
		wantErr   bool
	}{
		{
			name:      "予約が存在",
			args:      args{hotelID: 1},
			wantCount: 2,
		},
		{
			name:      "予約が存在しない",
			args:      args{hotelID: 100},
			wantCount: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := SelectPlansFromHotelID(tt.args.hotelID)
			if (err != nil) != tt.wantErr {
				t.Errorf("SelectPlansFromHotelID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if len(got) != tt.wantCount {
				t.Errorf("Count of want and got is different got = %v, want %v", len(got), tt.wantCount)
			}
		})
	}
}
