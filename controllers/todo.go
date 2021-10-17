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
type ChangeStatusRequest struct {
	ID     int64  `json:"id"`
	Status string `json:"status"`
}

type ChangeStatusResponse struct {
	ID int64 `json:"id"`
}

func Changestatus(c *fiber.Ctx) error {
	var req ChangeStatusRequest
	if err := c.BodyParser(&req); err != nil {
		return err
	}
	result := database.DB.Model(models.Todo{}).Where("id=?", req.ID).Update("status", req.Status)
	if result.Error != nil {
		return result.Error
	}
	return c.JSON(fiber.Map{
		"id": fmt.Sprintf("%d", req.ID),
	})
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

type ListsResponse struct {
	Lists []TodoListView `json:"lists"`
}

type TodoListView struct {
	ID     int64  `json:"id"`
	Todo   string `json:"todo"`
	Detail string `json:"detail"`
	Status string `json:"status"`
}

func Lists(c *fiber.Ctx) error {
	var todos []models.Todo
	var res ListsResponse

	result := database.DB.Where("status!=?", "d").Find(&todos)
	if result.Error != nil {
		return result.Error
	}

	for _, v := range todos {
		tmp := TodoListView{
			ID:     int64(v.ID),
			Todo:   v.Todo,
			Detail: v.Detail.String,
			Status: v.Status,
		}

		res.Lists = append(res.Lists, tmp)
	}

	return c.JSON(res)
}
