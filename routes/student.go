package routes

import (
	"NetFinal-Fiber_Gorm/database"
	"NetFinal-Fiber_Gorm/models"

	// "errors"
	"time"

	"github.com/gofiber/fiber/v2"
)

type Student struct {
	Student_ID uint    `json:"student_id"`
	Avarage    float32 `json:"avarage"`
	FirstName  string  `json:"first_name"`
	LastName   string  `json:"last_name"`
	Courses    int     `json:"courses"`
	// Last_updated time.Time
	UpdatedAt time.Time
}

func findStudent(student_id int, student *models.Student) any {
	database.Database.GOlestan.Find(&student, "Student_ID = ?", student_id)

	if student.Student_ID == 0 {
		return "This student doesn't exist"
	}
	return nil
}

// func checkDup(student_id int, student *models.Student) any {
// 	if err := findStudent(student_id, student); err != nil {
// 		return "A student with this Student ID already exists\n Please enter another Student_ID"
// 	}
// 	return nil
// }

func CreateStudent(c *fiber.Ctx) error {
	var student models.Student

	if err := c.BodyParser(&student); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	if student.Student_ID == 0 {
		return c.Status(400).JSON("Please enter a number for Student_ID")
	}

	if student.Student_ID < 10000000 && student.Student_ID > 1000000 {
		return c.Status(400).JSON("Student_ID must be a number with 8 digits")
	}

	var dup Student
	database.Database.GOlestan.First(&dup, "student_id = ?", student.Student_ID)

	if dup.Student_ID != 0 {
		return c.Status(400).JSON("A student with this Student ID already exists")
	}

	// student.Last_updated = time.Now()

	database.Database.GOlestan.Create(&student)

	return c.Status(200).JSON(student)
}

func GetStudent(c *fiber.Ctx) error {
	student_id, err := c.ParamsInt("studentid")

	var student models.Student

	if err != nil {
		return c.Status(400).JSON("Student_ID should be an integer!")
	}

	if err := findStudent(student_id, &student); err != nil {
		return c.Status(400).JSON(err)
	}

	return c.Status(200).JSON(student)
}

func GetAllStudents(c *fiber.Ctx) error {
	students := []models.Student{}

	database.Database.GOlestan.Find(&students)

	return c.Status(200).JSON(students)
}

func UpdateStudent(c *fiber.Ctx) error {
	student_id, err := c.ParamsInt("studentid")

	var student models.Student

	if err != nil {
		return c.Status(400).JSON("Student_ID should be an integer!")
	}

	if err := findStudent(student_id, &student); err != nil {
		return c.Status(400).JSON(err)
	}

	type UpdateStudent struct {
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
	}

	var updateStudentData UpdateStudent

	if err := c.BodyParser(&updateStudentData); err != nil {
		return c.Status(500).JSON(err.Error())
	}

	student.FirstName = updateStudentData.FirstName
	student.LastName = updateStudentData.LastName

	database.Database.GOlestan.Save(&student)

	return c.Status(200).JSON(student)
}

func DeleteStudent(c *fiber.Ctx) error {
	student_id, err := c.ParamsInt("studentid")

	var student models.Student

	if err != nil {
		return c.Status(400).JSON("Student_ID should be an integer!")
	}

	if err := findStudent(student_id, &student); err != nil {
		return c.Status(400).JSON(err)
	}

	if err := database.Database.GOlestan.Delete(&student).Error; err != nil {
		return c.Status(404).JSON(err.Error())
	}

	return c.Status(200).SendString("Student Successfully Deleted")
}
