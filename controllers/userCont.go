package controllers

import (
	"be-pemilu/models"
	"be-pemilu/service"
	"be-pemilu/validation.go"

	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type uController struct {
	UserServ service.UserService
}

func UserController(userServ service.UserService) *uController {
	return &uController{userServ}
}

func (u *uController) Vote(c *fiber.Ctx) error {
	u64, _ := strconv.ParseUint(c.Params("id"), 10, 32)
	id := uint(u64)
	req := new(validation.Vote)
	err := c.BodyParser(req)

	if err != nil {
		return c.Status(400).JSON(err.Error())
	}

	valid := validator.New()
	err1 := valid.Struct(req)

	if err1 != nil {
		return c.Status(400).JSON(err1.Error())
	}

	data := models.User{
		PaslonID: req.PaslonID,
	}

	res, err := u.UserServ.Vote(id, data)
	_ = res

	if err != nil {
		return c.Status(400).JSON(err.Error())
	}
	return c.Status(200).JSON(fiber.Map{
		"paslon": res.PaslonID,
	})
}

func (co *uController) Result(ct *fiber.Ctx) error {
	users, err := co.UserServ.Result()

	if err != nil {
		return ct.Status(500).JSON(err.Error())
	}

	return ct.Status(200).JSON(users)
}

func (co *uController) FindOne(ct *fiber.Ctx) error {
	u64, _ := strconv.ParseUint(ct.Params("id"), 10, 32)
	id := uint(u64)
	users, err := co.UserServ.FindOne(id)

	if err != nil {
		return ct.Status(404).JSON(err.Error())
	}

	return ct.Status(200).JSON(users)
}

func (co *uController) Register(ct *fiber.Ctx) error {
	req := new(validation.RegisterUser)
	err := ct.BodyParser(req)

	if err != nil {
		return ct.Status(400).JSON(err.Error())
	}

	valid := validator.New()
	err1 := valid.Struct(req)

	if err1 != nil {
		return ct.Status(400).JSON(err1.Error())
	}

	data := models.User{
		Username: req.Username,
		Password: req.Password,
		Fullname: req.Fullname,
		Address:  req.Address,
		Gender:   req.Gender,
	}

	res, err := co.UserServ.Register(data)

	if err != nil {
		return ct.Status(500).JSON(err.Error())
	}

	return ct.Status(201).JSON(res)
}

func (u *uController) Update(c *fiber.Ctx) error {
	u64, _ := strconv.ParseUint(c.Params("id"), 10, 32)
	id := uint(u64)
	user, err := u.UserServ.FindOne(id)

	if err != nil {
		return c.Status(404).JSON(err.Error())
	}

	req := new(validation.UpdateUser)
	err1 := c.BodyParser(req)

	if err1 != nil {
		return c.Status(400).JSON(err.Error())
	}

	if req.Username != "" {
		user.Username = req.Username
	}
	if req.Password != "" {
		user.Password = req.Password
	}
	if req.Fullname != "" {
		user.Fullname = req.Fullname
	}
	if req.Address != "" {
		user.Address = req.Address
	}
	if req.Gender != "" {
		user.Gender = req.Gender
	}

	data := models.User{
		Username: user.Username,
		Password: user.Password,
		Fullname: user.Fullname,
		Address:  user.Address,
		Gender:   user.Gender,
	}

	upd, err := u.UserServ.UserUpdate(id, data)

	if err != nil {
		return c.Status(400).JSON(err.Error())
	}
	return c.Status(200).JSON(upd)
}

func (u *uController) DeleteUser(c *fiber.Ctx) error {
	u64, _ := strconv.ParseUint(c.Params("id"), 10, 32)
	id := uint(u64)
	res, err := u.UserServ.FindOne(id)
	_ = res

	if err != nil {
		return c.Status(404).JSON(err.Error())
	}

	err1 := u.UserServ.DeleteUser(id)

	if err1 != nil {
		return c.Status(500).JSON(err1.Error())
	}
	return c.Status(200).JSON(fiber.Map{
		"Message": "User deleted successfully!",
	})
}
