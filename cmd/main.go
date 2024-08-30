package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"mywebsite.tv/name/cmd/database"
	"mywebsite.tv/name/cmd/handlers"
	"mywebsite.tv/name/cmd/repository"
	"mywebsite.tv/name/cmd/service"
)

func main() {
	app := fiber.New(fiber.Config{
		Views: html.New("./views", ".html"),
	})

	db := database.Connect()
	repo := repository.NewRepo(db)
	srv := service.NewService(repo)
	handler := handlers.NewHandler(srv)

	api := app.Group("/api/v1")
	api.Get("/posts", handler.GetPosts)
	api.Get("/posts/:id", handler.GetPostID)
	api.Post("/posts", handler.CreatePosts)
	api.Put("/posts/:id", handler.UpdatePost)
	api.Delete("/posts/:id", handler.DeletePost)
	app.Static("/", "./views")
	app.Static("/css", "css")

	app.Listen(":8080")
}
