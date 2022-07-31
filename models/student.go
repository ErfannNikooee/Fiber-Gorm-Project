package models

import "time"

type Student struct {
	// ID           uint    `json:"id" gorm:"primaryKey"`
	Student_ID uint    `json:"student_id" gorm:"primaryKey"`
	FirstName  string  `json:"first_name"`
	LastName   string  `json:"last_name"`
	Avarage    float32 `json:"avarage"`
	Courses    int     `json:"courses"`
	// Last_updated time.Time
	UpdatedAt time.Time
}
