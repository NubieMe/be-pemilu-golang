package routes

import (
	"be-pemilu/controllers"
	"be-pemilu/database"
	"be-pemilu/service"

	"github.com/gofiber/fiber/v2"
)

func RouteApp(f *fiber.App) {
	s := service.Repo2DB(database.DB)
	u := controllers.UserController(s)
	a := controllers.ArticleController(s)
	par := controllers.PartyController(s)
	pas := controllers.PaslonController(s)

	app := fiber.New()
	f.Mount("/api", app)

	app.Patch("/vote/:id", u.Vote)
	app.Get("/result", u.Result)
	app.Get("/user/:id", u.FindOne)
	app.Post("/register", u.Register)
	app.Patch("/user/:id", u.Update)
	app.Delete("/user/:id", u.DeleteUser)

	app.Get("/article", a.ArtAll)
	app.Get("/article/:id", a.ArtOne)
	app.Post("/article", a.CreArt)
	app.Patch("/article/:id", a.UpdArt)
	app.Delete("/article/:id", a.DeleteArt)

	app.Get("/party", par.PartyAll)
	app.Get("/party/:id", par.PartyOne)
	app.Post("/party", par.CrePart)
	app.Patch("/party/:id", par.UpdPart)
	app.Delete("/party/:id", par.DeletePart)

	app.Get("/paslon", pas.PaslonAll)
	app.Get("/paslon/:id", pas.PaslonOne)
	app.Post("/paslon", pas.CrePas)
	app.Patch("/paslon/:id", pas.UpdPas)
	app.Delete("/paslon/:id", pas.DeletePas)
}
