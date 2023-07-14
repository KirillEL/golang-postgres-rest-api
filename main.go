package main

import (
	"fmt"
	"github.com/KirillEL/golang-postgres-rest-api/migrations"
	"github.com/KirillEL/golang-postgres-rest-api/models"
	"github.com/KirillEL/golang-postgres-rest-api/storage"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
	"log"
	"net/http"
)

type Repository struct {
	DB *gorm.DB
}

func (r *Repository) GetAllCars(ctx *fiber.Ctx) error {
	carModels := &[]models.Car{}
	err := r.DB.Find(&carModels).Error
	if err != nil {
		err := ctx.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "could not get cars",
		})
		if err != nil {
			return err
		}
	}
	ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "success!",
		"data":    carModels,
	})
	return nil
}

func (r *Repository) GetCarById(ctx *fiber.Ctx) error {
	carModel := models.Car{}
	id := ctx.Params("id")
	if id == "" {
		ctx.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "id not set!",
		})
	}

	err := r.DB.Where("id = ?", id).First(&carModel).Error
	if err != nil {
		ctx.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "not found",
		})
	}
	ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "success",
		"data":    carModel,
	})
	return nil
}

func (r *Repository) CreateCar(ctx *fiber.Ctx) error {
	carModel := models.Car{}
	err := ctx.BodyParser(&carModel)
	if err != nil {
		return err
	}
	err = r.DB.Create(&carModel).Error
	if err != nil {
		ctx.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "cannot create a car",
		})
	}
	ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "success",
	})
	return nil
}

func (r *Repository) DeleteCar(ctx *fiber.Ctx) error {
	carModel := models.Car{}
	id := ctx.Params("id")
	if id == "" {
		ctx.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"message": "id cannot be empty!",
		})
		return nil
	}
	err := r.DB.Delete(carModel, id).Error
	if err != nil {
		ctx.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "bad request",
		})
	}
	ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": fmt.Sprintf("success deleted car with id=%s", id),
	})

	return nil
}

func (r *Repository) SetupRoutes(app *fiber.App) {
	api := app.Group("/api")
	api.Get("/cars", r.GetAllCars)
	api.Get("/car/:id", r.GetCarById)
	api.Post("/car", r.CreateCar)
	api.Delete("/car/:id", r.DeleteCar)

}

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	config := storage.InitStorage()

	db, e := storage.NewConnection(&config)

	if e != nil {
		log.Fatal(err)
	}

	err = migrations.MigrateCars(db)

	if err != nil {
		log.Fatal(err)
	}

	r := Repository{
		DB: db,
	}

	app := fiber.New()
	r.SetupRoutes(app)
	app.Listen(":8085")

}
