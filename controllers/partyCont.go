package controllers

import (
	"be-pemilu/models"
	"be-pemilu/service"
	"be-pemilu/validation.go"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type parController struct {
	PartyServ service.PartyService
}

func PartyController(partServ service.PartyService) *parController {
	return &parController{partServ}
}

func (p *parController) PartyAll(c *fiber.Ctx) error {
	parties, err := p.PartyServ.PartyAll()

	if err != nil {
		return c.Status(500).JSON(err.Error())
	}
	return c.Status(200).JSON(parties)
}

func (p *parController) PartyOne(c *fiber.Ctx) error {
	u64, _ := strconv.ParseUint(c.Params("id"), 10, 32)
	id := uint(u64)
	party, err := p.PartyServ.PartyOne(id)

	if err != nil {
		return c.Status(404).JSON(err.Error())
	}
	return c.Status(200).JSON(party)
}

func (p *parController) CrePart(c *fiber.Ctx) error {
	req := new(validation.CreateParty)
	err := c.BodyParser(req)

	if err != nil {
		return c.Status(400).JSON(err.Error())
	}

	valid := validator.New()
	err1 := valid.Struct(req)

	if err1 != nil {
		return c.Status(400).JSON(err.Error())
	}

	data := models.Party{
		Name:     req.Name,
		Leader:   req.Leader,
		Image:    req.Image,
		Visimisi: req.Visimisi,
		PaslonID: req.PaslonID,
	}

	res, err := p.PartyServ.CrePart(data)

	if err != nil {
		return c.Status(400).JSON(err.Error())
	}
	return c.Status(201).JSON(res)
}

func (p *parController) UpdPart(c *fiber.Ctx) error {
	u64, _ := strconv.ParseUint(c.Params("id"), 10, 32)
	id := uint(u64)
	par, err := p.PartyServ.PartyOne(id)
	_ = par

	if err != nil {
		return c.Status(404).JSON(err.Error())
	}

	req := new(validation.UpdateParty)
	err1 := c.BodyParser(req)

	if err1 != nil {
		return c.Status(400).JSON(err.Error())
	}

	valid := validator.New()
	err2 := valid.Struct(req)

	if err2 != nil {
		return c.Status(400).JSON(err.Error())
	}

	if req.Name != "" {
		par.Name = req.Name
	}
	if req.Leader != "" {
		par.Leader = req.Leader
	}
	if req.Image != "" {
		par.Image = req.Image
	}
	if req.Visimisi != nil {
		par.Visimisi = req.Visimisi
	}
	if req.PaslonID != 0 {
		par.PaslonID = req.PaslonID
	}

	data := models.Party{
		Name:     par.Name,
		Leader:   par.Leader,
		Image:    par.Image,
		Visimisi: par.Visimisi,
		PaslonID: par.PaslonID,
	}

	res, err := p.PartyServ.UpdPart(id, data)

	if err != nil {
		return c.Status(500).JSON(err.Error())
	}
	return c.Status(200).JSON(res)
}

func (p *parController) DeletePart(c *fiber.Ctx) error {
	u64, _ := strconv.ParseUint(c.Params("id"), 10, 32)
	id := uint(u64)
	par, err := p.PartyServ.PartyOne(id)
	_ = par

	if err != nil {
		return c.Status(404).JSON(err.Error())
	}

	err1 := p.PartyServ.DeletePart(id)

	if err1 != nil {
		return c.Status(500).JSON(err1.Error())
	}
	return c.Status(200).JSON(fiber.Map{
		"Message": "Party deleted successfully!",
	})
}
