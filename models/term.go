package models

type Term struct {
	// ID           uint    `json:"id"`
	Term_ID      int     `json:"term_id" gorm:"primaryKey"`
	StudentRefer int     `json:"student_id"`
	Student      Student `gorm:"foreignKey:StudentRefer"`
	CourseRefer  int     `json:"course_id"`
	Course       Course  `gorm:"foreignKey:CourseRefer"`
}
