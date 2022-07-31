package routes

import (
	"NetFinal-Fiber_Gorm/database"
	"NetFinal-Fiber_Gorm/models"

	"github.com/gofiber/fiber/v2"
)

type GradeReport struct {
	Report_ID   int     `json:"report_id"`
	CourseRefer int     `json:"course_id"`
	Grade       float32 `json:"grade"`
}

// func calculateAVG(avarage float32, courses int, newgrade float32) float32 {
// 	avarage = ((avarage * float32(courses)) + newgrade) / (float32(courses) + 1)
// 	return avarage
// }

func EnterGrade(c *fiber.Ctx) error {
	var grade models.GradeReport

	if err := c.BodyParser(&grade); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	if grade.StudentRefer == 0 {
		return c.Status(400).JSON("Please enter a number for Student_ID")
	}

	if grade.StudentRefer < 10000000 && grade.StudentRefer > 1000000 {
		return c.Status(400).JSON("Student_ID must be a number with 8 digits")
	}

	if err := findStudent(grade.StudentRefer, &grade.Student); err != nil {
		return c.Status(400).JSON(err)
	}

	if grade.CourseRefer == 0 {
		return c.Status(400).JSON("Please enter course informations")
	}

	if err := findCourse(grade.CourseRefer, &grade.Course); err != nil {
		return c.Status(400).JSON(err)
	}

	var dup models.GradeReport

	database.Database.GOlestan.Find(&dup,
		"student_refer = ? AND course_refer = ?", grade.StudentRefer, grade.CourseRefer)

	if dup.Report_ID != 0 {
		return c.Status(400).JSON("Duplicate information")
	}

	database.Database.GOlestan.Create(&grade)

	avarage := grade.Student.Avarage

	courses := grade.Student.Courses

	new_grade := grade.Grade

	// grade.Student.Avarage = calculateAVG(grade.Student.Avarage, grade.Student.Courses, grade.Grade)
	avarage = ((avarage * float32(courses)) + new_grade) / (float32(courses) + 1)

	grade.Student.Avarage = avarage

	grade.Student.Courses++

	database.Database.GOlestan.Save(grade.Student)

	return c.Status(200).JSON(grade)
}

func GetReport(c *fiber.Ctx) error {
	student_id, err := c.ParamsInt("studentid")

	report := []GradeReport{}

	if err != nil {
		return c.Status(400).JSON("Student_ID should be an integer!")
	}

	database.Database.GOlestan.Find(&report, "student_refer = ? ", student_id)

	return c.Status(200).JSON(report)
}

func UpdateGrade(c *fiber.Ctx) error {
	var report models.GradeReport

	if err := c.BodyParser(&report); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	if report.StudentRefer == 0 {
		return c.Status(400).JSON("Please enter a number for Student_ID")
	}

	if report.StudentRefer < 10000000 && report.StudentRefer > 1000000 {
		return c.Status(400).JSON("Student_ID must be a number with 8 digits")
	}

	if err := findStudent(report.StudentRefer, &report.Student); err != nil {
		return c.Status(400).JSON(err)
	}

	if report.CourseRefer == 0 {
		return c.Status(400).JSON("Please enter course informations")
	}

	if err := findCourse(report.CourseRefer, &report.Course); err != nil {
		return c.Status(400).JSON(err)
	}

	var selectedreport models.GradeReport

	database.Database.GOlestan.Find(&selectedreport,
		"student_refer = ? AND course_refer = ?", report.StudentRefer, report.CourseRefer)

	avarage := report.Student.Avarage

	courses := report.Student.Courses

	new_grade := report.Grade

	avarage = ((avarage * float32(courses)) - selectedreport.Grade + new_grade) / (float32(courses))

	report.Student.Avarage = avarage

	selectedreport.Grade = report.Grade

	database.Database.GOlestan.Save(&selectedreport)

	database.Database.GOlestan.Save(report.Student)

	return c.Status(200).JSON(report)
}

func DeleteGrade(c *fiber.Ctx) error {
	var report models.GradeReport

	if err := c.BodyParser(&report); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	if report.StudentRefer == 0 {
		return c.Status(400).JSON("Please enter a number for Student_ID")
	}

	if report.StudentRefer < 10000000 && report.StudentRefer > 1000000 {
		return c.Status(400).JSON("Student_ID must be a number with 8 digits")
	}

	if err := findStudent(report.StudentRefer, &report.Student); err != nil {
		return c.Status(400).JSON(err)
	}

	if report.CourseRefer == 0 {
		return c.Status(400).JSON("Please enter course informations")
	}

	if err := findCourse(report.CourseRefer, &report.Course); err != nil {
		return c.Status(400).JSON(err)
	}

	var selectedreport models.GradeReport

	database.Database.GOlestan.Find(&selectedreport,
		"student_refer = ? AND course_refer = ?", report.StudentRefer, report.CourseRefer)

	if err := database.Database.GOlestan.Delete(&selectedreport).Error; err != nil {
		return c.Status(404).JSON(err.Error())
	}

	var student models.Student

	database.Database.GOlestan.Find(&student, "student_id = ?", report.StudentRefer)

	avarage := student.Avarage

	courses := student.Courses

	remove_grade := selectedreport.Grade

	avarage = ((avarage * float32(courses)) - remove_grade) / (float32(courses) - 1)

	student.Avarage = avarage

	student.Courses--

	database.Database.GOlestan.Save(student)

	return c.Status(200).SendString("Grade Successfully Deleted")

}
