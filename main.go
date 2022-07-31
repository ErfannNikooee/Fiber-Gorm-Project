package main

import (
	"NetFinal-Fiber_Gorm/database"
	"NetFinal-Fiber_Gorm/routes"
	"log"

	"github.com/gofiber/fiber/v2"
)

func setupRoutes(app *fiber.App) {
	//Student Routes
	app.Post("/students", routes.CreateStudent)

	app.Get("/students", routes.GetAllStudents)

	app.Get("/students/:studentid", routes.GetStudent)

	app.Put("/students/:studentid", routes.UpdateStudent)

	app.Delete("/students/:studentid", routes.DeleteStudent)

	//Course Routes

	app.Post("/courses", routes.CreateCourse)

	app.Get("/courses", routes.GetAllCourse)

	app.Get("/courses/:course_id", routes.GetCourse)

	app.Put("/courses/:course_id", routes.UpateCourse)

	app.Delete("/courses/:course_id", routes.DeleteCourse)

	//Term Routes

	app.Post("/gradereport", routes.EnterGrade)

	app.Get("/gradereport/:studentid", routes.GetReport)

	app.Put("/gradereport", routes.UpdateGrade)

	app.Delete("/gradereport", routes.DeleteGrade)

}

func main() {
	database.ConnectDb()

	app := fiber.New()

	setupRoutes(app)

	log.Fatal(app.Listen(":3000"))
}
