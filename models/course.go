package models

type Course struct {
	// ID        uint   `json:"id"`
	Course_ID uint   `json:"course_id" gorm:"primaryKey"`
	Name      string `json:"course_name"`
}
