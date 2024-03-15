package controllers

import (
	"be-pemilu/models"
	"be-pemilu/service"
	"be-pemilu/validation.go"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type pasController struct {
	PasServ service.PaslonService
}

func PaslonController(pasServ service.PaslonService) *pasController {
	return &pasController{pasServ}
}

func (p *pasController) PaslonAll(c *fiber.Ctx) error {
	paslons, err := p.PasServ.PaslonAll()

	if err != nil {
		return c.Status(500).JSON(err.Error())
	}
	return c.Status(200).JSON(paslons)
}

func (p *pasController) PaslonOne(c *fiber.Ctx) error {
	u64, _ := strconv.ParseUint(c.Params("id"), 10, 32)
	id := uint(u64)
	paslon, err := p.PasServ.PaslonOne(id)

	if err != nil {
		return c.Status(404).JSON(err.Error())
	}
	return c.Status(200).JSON(paslon)
}

func (p *pasController) CrePas(c *fiber.Ctx) error {
	req := new(validation.CreatePaslon)
	err := c.BodyParser(req)

	if err != nil {
		return c.Status(400).JSON(err.Error())
	}

	valid := validator.New()
	err1 := valid.Struct(req)

	if err1 != nil {
		return c.Status(400).JSON(err1.Error())
	}

	data := models.Paslon{
		Name:     req.Name,
		Image:    req.Image,
		Visimisi: req.Visimisi,
	}

	res, err := p.PasServ.CrePas(data)

	if err != nil {
		return c.Status(400).JSON(err.Error())
	}
	return c.Status(201).JSON(res)
}

func (p *pasController) UpdPas(c *fiber.Ctx) error {
	u64, _ := strconv.ParseUint(c.Params("id"), 10, 32)
	id := uint(u64)
	pas, err := p.PasServ.PaslonOne(id)

	if err != nil {
		return c.Status(404).JSON(err.Error())
	}

	req := new(validation.UpdatePaslon)
	err1 := c.BodyParser(req)

	if err1 != nil {
		return c.Status(400).JSON(err1.Error())
	}

	valid := validator.New()
	err2 := valid.Struct(req)

	if err2 != nil {
		return c.Status(400).JSON(err2.Error())
	}

	if req.Name != "" {
		pas.Name = req.Name
	}
	if req.Image != "" {
		pas.Image = req.Image
	}
	if req.Visimisi != nil {
		pas.Visimisi = req.Visimisi
	}

	data := service.PaslonAPI{
		Name:     pas.Name,
		Image:    pas.Image,
		Visimisi: pas.Visimisi,
	}

	res, err := p.PasServ.UpdPas(id, data)

	if err != nil {
		return c.Status(500).JSON(err.Error())
	}
	return c.Status(200).JSON(res)
}

func (p *pasController) DeletePas(c *fiber.Ctx) error {
	u64, _ := strconv.ParseUint(c.Params("id"), 10, 32)
	id := uint(u64)
	pas, err := p.PasServ.PaslonOne(id)
	_ = pas

	if err != nil {
		return c.Status(404).JSON(err.Error())
	}

	err1 := p.PasServ.DeletePas(id)

	if err1 != nil {
		return c.Status(500).JSON(err1.Error())
	}
	return c.Status(200).JSON(fiber.Map{
		"Message": "Paslon deleted successfully!",
	})
}
