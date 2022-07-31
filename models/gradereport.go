package models

type GradeReport struct {
	// ID           uint    `json:"id"`
	Report_ID    int     `json:"report_id" gorm:"primaryKey"`
	StudentRefer int     `json:"student_id"`
	Student      Student `gorm:"foreignKey:StudentRefer"`
	CourseRefer  int     `json:"course_id"`
	Course       Course  `gorm:"foreignKey:CourseRefer"`
	Grade        float32 `json:"grade"`
}
