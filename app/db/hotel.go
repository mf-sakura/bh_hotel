package db

import ()

type Hotel struct {
	ID   int64  `db:"id"`
	Name string `db:"name"`
}

func InsertHotel(hotel Hotel) (int64, error) {
	q := "INSERT INTO `hotel` (`name`) VALUES (?)"

	res, err := db.Exec(q, hotel.Name)
	if err != nil {
		return 0, err
	}
	return res.LastInsertId()
}

func GetHotel(id int64) (*Hotel, error) {

	q := "SELECT * FROM `hotel` WHERE id = ?"
	var hotel Hotel
	if err := db.Get(&hotel, q, id); err != nil {
		return nil, err
	}

	return &hotel, nil
}

func SelectHotelAll() ([]*Hotel, error) {
	q := "SELECT * FROM `hotel` "
	var hotels []*Hotel
	if err := db.Select(hotels, q); err != nil {
		return nil, err
	}
	return hotels, nil
}
