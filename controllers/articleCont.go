package controllers

import (
	"be-pemilu/models"
	"be-pemilu/service"
	"be-pemilu/validation.go"
	"strconv"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type aController struct {
	ArticServ service.ArticleService
}

func ArticleController(articServ service.ArticleService) *aController {
	return &aController{articServ}
}

func (a *aController) ArtAll(c *fiber.Ctx) error {
	artics, err := a.ArticServ.ArtAll()

	if err != nil {
		return c.Status(500).JSON(err.Error())
	}
	return c.Status(200).JSON(artics)
}

func (a *aController) ArtOne(c *fiber.Ctx) error {
	u64, _ := strconv.ParseUint(c.Params("id"), 10, 32)
	id := uint(u64)
	artic, err := a.ArticServ.ArtOne(id)

	if err != nil {
		return c.Status(404).JSON(err.Error())
	}
	return c.Status(200).JSON(artic)
}

func (a *aController) CreArt(c *fiber.Ctx) error {
	req := new(validation.CreateArticle)
	err := c.BodyParser(req)

	if err != nil {
		return c.Status(400).JSON(err.Error())
	}

	valid := validator.New()
	err1 := valid.Struct(req)

	if err1 != nil {
		return c.Status(400).JSON(err1.Error())
	}

	data := models.Article{
		Title:       req.Title,
		Description: req.Description,
		Image:       req.Image,
		UserID:      req.UserID,
	}

	artic, err := a.ArticServ.CreArt(data)

	if err != nil {
		c.Status(400).JSON(err.Error())
	}
	return c.Status(201).JSON(artic)
}

func (a *aController) UpdArt(c *fiber.Ctx) error {
	u64, _ := strconv.ParseUint(c.Params("id"), 10, 32)
	id := uint(u64)
	art, err := a.ArticServ.ArtOne(id)

	if err != nil {
		return c.Status(404).JSON(err.Error())
	}

	req := new(validation.UpdateArticle)
	err1 := c.BodyParser(req)

	if err1 != nil {
		return c.Status(400).JSON(err.Error())
	}

	valid := validator.New()
	err2 := valid.Struct(req)

	if err2 != nil {
		return c.Status(400).JSON(err2.Error())
	}

	if req.Title != "" {
		art.Title = req.Title
	}
	if req.Description != "" {
		art.Description = req.Description
	}
	if req.Image != "" {
		art.Image = req.Image
	}
	if req.UserID != 0 {
		art.UserID = req.UserID
	}

	data := service.UpdateArticle{
		Title:       art.Title,
		Description: art.Description,
		Image:       art.Image,
		UserID:      art.UserID,
		UpdatedAt:   time.Now(),
	}

	res, err := a.ArticServ.UpdArt(id, data)

	if err != nil {
		return c.Status(400).JSON(err.Error())
	}
	return c.Status(200).JSON(res)
}

func (a *aController) DeleteArt(c *fiber.Ctx) error {
	u64, _ := strconv.ParseUint(c.Params("id"), 10, 32)
	id := uint(u64)
	res, err := a.ArticServ.ArtOne(id)
	_ = res

	if err != nil {
		return c.Status(404).JSON(err.Error())
	}

	err1 := a.ArticServ.DeleteArt(id)

	if err1 != nil {
		return c.Status(500).JSON(err1.Error())
	}
	return c.Status(200).JSON(fiber.Map{
		"Message": "Article deleted successfully!",
	})
}
