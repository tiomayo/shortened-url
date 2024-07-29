package routes

import (
	"fmt"
	"shortened-url/config"
	"shortened-url/model"

	"github.com/gofiber/fiber/v2"
)

type routesRepository struct {
	cfg *config.Cfg
}

type RepoInterface interface {
	Get(slug string) model.RouteRequest
	Set(req model.RouteRequest)
	List(c *fiber.Ctx, userID string) []model.RouteRequest
}

func New(cfg *config.Cfg) RepoInterface {
	return routesRepository{
		cfg: cfg,
	}
}

func (r routesRepository) Get(slug string) model.RouteRequest {
	if v, found := r.cfg.Routes[slug]; found {
		return v
	}
	return model.RouteRequest{}
}

func (r routesRepository) Set(req model.RouteRequest) {
	r.cfg.Routes[req.Shortened] = req
}

func (r routesRepository) List(c *fiber.Ctx, userID string) []model.RouteRequest {
	result := []model.RouteRequest{}
	for _, v := range r.cfg.Routes {
		if userID == "admin" || (userID != "" && userID == v.UserID) {
			result = append(result, model.RouteRequest{
				Url:       v.Url,
				Shortened: fmt.Sprintf("%v/%v", c.BaseURL(), v.Shortened),
				UserID:    v.UserID,
				Counter:   v.Counter,
			})
		}
	}
	return result
}
