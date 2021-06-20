package controllers

import (
	"database/sql"
	"fmt"
	"gotodo/database"
	"gotodo/models"

	"github.com/gofiber/fiber/v2"
)

type CreateRequest struct {
	Todo   string `json:"todo"`
	Detail string `json:"detail"`
}

func Create(c *fiber.Ctx) error {
	var req CreateRequest
	if err := c.BodyParser(&req); err != nil {
		return err
	}

	var todo models.Todo
	todo.Todo = req.Todo
	if req.Detail != "" {
		todo.Detail = sql.NullString{String: req.Detail, Valid: true}
	}
	todo.Status = "p"

	result := database.DB.Create(&todo)
	if result.Error != nil {
		return result.Error
	}

	return c.JSON(fiber.Map{
		"id": fmt.Sprintf("%d", todo.ID),
	})
}
