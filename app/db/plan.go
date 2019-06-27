package db

import (
	"time"
)

type Plan struct {
	ID          int64     `db:"id"`
	HotelID     int64     `db:"hotel_id"`
	Description string    `db:"description"`
	Date        time.Time `db:"date"`
	Total       uint64    `db:"total"`
	Available   uint64    `db:"available"`
	Cost        uint64    `db:"cost"`
}

func InsertPlan(plan *Plan) (int64, error) {
	q := "INSERT INTO `plan` ( `hotel_id`, `description`, `date`, `total`, `available`, `cost`) VALUES " +
		" (?, ?, ?, ?, ?, ?)"
	res, err := db.Exec(q, plan.HotelID, plan.Description, plan.Date, plan.Total, plan.Available, plan.Cost)
	if err != nil {
		return 0, err
	}
	return res.LastInsertId()
}

func GetPlanFromID(planID int64, forUpdate bool) (*Plan, error) {
	q := "SELECT * FROM `plan` WHERE `id` = ? "
	if forUpdate {
		q += "FOR UPDATE"
	}
	var plan Plan
	if err := db.Get(&plan, q, planID); err != nil {
		return nil, err
	}
	return &plan, nil
}

func decrementPlanAvailable(planID int64) error {
	q := "UPDATE `plan` SET `available` = `available` - 1 WHERE `id` = ?"
	if _, err := db.Exec(q, planID); err != nil {
		return err
	}
	return nil
}

func incrementPlanAvailable(planID int64) error {
	q := "UPDATE `plan` SET `available` = `available` + 1 WHERE `id` = ?"
	if _, err := db.Exec(q, planID); err != nil {
		return err
	}
	return nil
}

func SelectPlansFromHotelID(hotelID int64) ([]*Plan, error) {
	q := "SELECT * FROM `plan` WHERE `hotel_id` = ?"
	var plans []*Plan
	if err := db.Select(&plans, q, hotelID); err != nil {
		return nil, err
	}
	return plans, nil
}
