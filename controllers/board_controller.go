package controllers

import (
	"github.com/ferryku8/project-management/models"
	"github.com/ferryku8/project-management/services"
	"github.com/ferryku8/project-management/utils"
	"github.com/gofiber/fiber/v2"
)

type BoardController struct {
	service services.BoardService
}

func NewBoardController(s services.BoardService) *BoardController {
	return &BoardController{service: s}
}

func (c *BoardController) CreateBoard(ctx *fiber.Ctx) error {
	board := new(models.Board)
	if err := ctx.BodyParser(board); err != nil {
		return utils.BadRequest(ctx, "Gagal membaca request", err.Error())
	}
	if err := c.service.Create(board); err != nil {
		return utils.BadRequest(ctx, "Gagal Menyimpan Data", err.Error())
	}
	return utils.Success(ctx, "Board Berhasil Dibuat", board)
}
