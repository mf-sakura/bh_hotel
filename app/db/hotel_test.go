package db

import (
	"reflect"
	"testing"
)

func TestInsertHotel(t *testing.T) {
	PrepareFixture("testdata/hotel")
	type args struct {
		hotel Hotel
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "正常系(クエリチェック)",
			args: args{
				Hotel{Name: "hoge"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if _, err := InsertHotel(tt.args.hotel); (err != nil) != tt.wantErr {
				t.Errorf("InsertHotel() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGetHotel(t *testing.T) {
	PrepareFixture("testdata/hotel")

	type args struct {
		id int64
	}

	tests := []struct {
		name    string
		args    args
		want    *Hotel
		wantErr bool
	}{
		{
			name: "正常系",
			args: args{1},
			want: &Hotel{
				ID:   1,
				Name: "foo",
			},
		},
		{
			name:    "存在しないID",
			args:    args{100},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetHotel(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetHotel() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetHotel() = %v, want %v", got, tt.want)
			}
		})
	}
}
