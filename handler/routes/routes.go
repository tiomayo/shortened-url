package routes

import (
	"fmt"
	"net/http"
	"shortened-url/helper"
	"shortened-url/model"
	"shortened-url/repository/routes"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
)

type handlerInterface interface {
	Shorten(ctx *fiber.Ctx) error
	Get(ctx *fiber.Ctx) error
	All(ctx *fiber.Ctx) error
}
type Handler struct {
	repo      routes.RepoInterface
	validator *validator.Validate
}

func New(repo routes.RepoInterface, validator *validator.Validate) handlerInterface {
	return Handler{
		repo:      repo,
		validator: validator,
	}
}

func (h Handler) Get(ctx *fiber.Ctx) error {
	original := h.repo.Get(ctx.Params("url"))
	if original.Url == "" {
		return fiber.ErrNotFound
	}

	original.AddCounter()
	h.repo.Set(original)

	return ctx.Redirect(original.Url, http.StatusMovedPermanently)
}

func (h Handler) Shorten(ctx *fiber.Ctx) error {
	newRoute := model.RouteRequest{}
	if err := ctx.BodyParser(&newRoute); err != nil {
		return err
	}

	if errs := h.validator.Struct(newRoute); errs != nil {
		return errs
	}
	newRoute.UserID = utils.CopyString(ctx.Get("User-X"))
	newRoute.Shortened = helper.GenerateSlug(h.repo)
	h.repo.Set(newRoute)

	ctx.Append("Max-Unique-Slug", fmt.Sprintf("%v", helper.CalculateUniqueSlugs()))
	return ctx.JSON(map[string]string{"Shortened": fmt.Sprintf("%v/%v", ctx.BaseURL(), newRoute.Shortened)})
}

func (h Handler) All(ctx *fiber.Ctx) error {
	listRoute := h.repo.List(ctx, ctx.Get("User-X", ""))
	return ctx.JSON(listRoute)
}
