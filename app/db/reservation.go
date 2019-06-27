package db

import (
	"database/sql"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

type Reservation struct {
	ID         int64  `db:"id"`
	PlanID     int64  `db:"plan_id"`
	UserID     int64  `db:"user_id"`
	SequenceID string `db:"sequence_id"`
}

func SelectReservationsFromUserID(userID int64) ([]*Reservation, error) {
	q := "SELECT * FROM `reservation` WHERE `user_id` = ?"
	var reservations []*Reservation
	if err := db.Select(&reservations, q, userID); err != nil {
		return nil, err
	}
	return reservations, nil
}

func ReserveHotel(reservation *Reservation) (int64, error) {
	var id int64
	err := db.RunInTx(func(tx *sqlx.Tx) error {
		plan, err := GetPlanFromID(reservation.PlanID, true)
		if err != nil {
			return err
		}
		if plan.Available == 0 {
			return errors.New("All plans are booked.")
		}
		if err := decrementPlanAvailable(reservation.PlanID); err != nil {
			return err
		}
		res, err := insertReservation(reservation)
		if err != nil {
			return err
		}
		id, err = res.LastInsertId()
		if err != nil {
			return err
		}
		return nil
	})
	return id, err
}

func CancelHotel(reservationID int64) error {
	err := db.RunInTx(func(tx *sqlx.Tx) error {
		reservation, err := GetReservationFromID(reservationID, true)
		if err != nil {
			return err
		}
		plan, err := GetPlanFromID(reservation.PlanID, true)
		if err != nil {
			return err
		}
		if plan.Available >= plan.Total {
			return errors.New("All plans are vacancy.")
		}
		if err := incrementPlanAvailable(plan.ID); err != nil {
			return err
		}
		err = deleteReservation(reservationID)
		if err != nil {
			return err
		}

		return nil
	})

	return err
}

func insertReservation(reservation *Reservation) (sql.Result, error) {
	q := "INSERT INTO `reservation` ( `plan_id`, `user_id`, `sequence_id`) VALUES " +
		" (?, ?, ?)"
	res, err := db.Exec(q, reservation.PlanID, reservation.UserID, reservation.SequenceID)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func deleteReservation(reservationID int64) error {
	q := "DELETE FROM `reservation` WHERE `id` = ?"
	_, err := db.Exec(q, reservationID)
	if err != nil {
		return err
	}
	return nil
}

func GetReservationFromID(reservationID int64, forUpdate bool) (*Reservation, error) {
	q := "SELECT * FROM `reservation` WHERE `id` = ? "
	if forUpdate {
		q += "FOR UPDATE"
	}
	var reservation Reservation
	if err := db.Get(&reservation, q, reservationID); err != nil {
		return nil, err
	}
	return &reservation, nil
}
