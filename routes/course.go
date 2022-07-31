package routes

import (
	"NetFinal-Fiber_Gorm/database"
	"NetFinal-Fiber_Gorm/models"

	"github.com/gofiber/fiber/v2"
)

type Course struct {
	Course_ID uint   `json:"course_id"`
	Name      string `json:"course_name"`
}

func findCourse(course_id int, course *models.Course) any {
	database.Database.GOlestan.Find(&course, "Course_ID = ?", course_id)

	if course.Course_ID == 0 {
		return "This Course doesn't exist"
	}
	return nil
}

func CreateCourse(c *fiber.Ctx) error {

	var course models.Course

	if err := c.BodyParser(&course); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	if course.Course_ID == 0 || course.Name == "" {
		return c.Status(400).JSON("Please enter course informations")
	}

	var dup Course
	database.Database.GOlestan.First(&dup, "course_id = ?", course.Course_ID)
	if dup.Course_ID != 0 || dup.Name == course.Name {
		return c.Status(400).JSON("A course with this information already exists")
	}

	database.Database.GOlestan.Create(&course)

	return c.Status(200).JSON(course)
}

func GetCourse(c *fiber.Ctx) error {
	course_id, err := c.ParamsInt("course_id")

	var course models.Course

	if err != nil {
		return c.Status(400).JSON("Course_ID should be an integer!")
	}

	if err := findCourse(course_id, &course); err != nil {
		return c.Status(400).JSON(err)
	}

	return c.Status(200).JSON(course)

}

func GetAllCourse(c *fiber.Ctx) error {

	courses := []models.Course{}

	database.Database.GOlestan.Find(&courses)

	return c.Status(200).JSON(courses)
}

func UpateCourse(c *fiber.Ctx) error {
	course_id, err := c.ParamsInt("course_id")

	var course models.Course

	if err != nil {
		return c.Status(400).JSON("Course_ID should be an integer!")
	}

	if err := findCourse(course_id, &course); err != nil {
		return c.Status(400).JSON(err)
	}

	type UpdateCourse struct {
		Name string `json:"course_name"`
	}

	var updateCourseData UpdateCourse

	if err := c.BodyParser(&updateCourseData); err != nil {
		return c.Status(500).JSON(err.Error())
	}

	course.Name = updateCourseData.Name

	database.Database.GOlestan.Save(&course)

	return c.Status(200).JSON(course)

}

func DeleteCourse(c *fiber.Ctx) error {
	course_id, err := c.ParamsInt("course_id")

	var course models.Course

	if err != nil {
		return c.Status(400).JSON("Course_ID should be an integer!")
	}

	if err := findCourse(course_id, &course); err != nil {
		return c.Status(400).JSON(err)
	}

	if err := database.Database.GOlestan.Delete(&course).Error; err != nil {
		return c.Status(404).JSON(err.Error())
	}

	return c.Status(200).SendString("Course Successfully Deleted")

}
