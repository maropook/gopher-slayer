package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/maropook/gopher-slayer/internal/model"
	"github.com/maropook/gopher-slayer/internal/service"
)

type HeroHandler struct {
	heroService *service.HeroService
}

func NewHeroHandler(heroService *service.HeroService) *HeroHandler {
	return &HeroHandler{heroService: heroService}
}

// GetHero returns the current hero status.
// GET /api/hero
func (h *HeroHandler) GetHero(c echo.Context) error {
	hero, err := h.heroService.GetHero()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, hero)
}

// UpdateName updates the hero's name.
// PUT /api/hero/name
// This is an example of a PUT endpoint — students can reference this for Lv2.
func (h *HeroHandler) UpdateName(c echo.Context) error {
	var req model.UpdateNameRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request body"})
	}
	if req.Name == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "name is required"})
	}
	if err := h.heroService.UpdateName(req.Name); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, map[string]string{"message": "Name updated successfully"})
}

// UpdateExperience updates the hero's experience points.
// PUT /api/hero/experience
func (h *HeroHandler) UpdateExperience(c echo.Context) error {
	var req model.UpdateExperienceRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request body"})
	}
	if err := h.heroService.UpdateExperience(req.Experience); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, map[string]string{"message": "Experience updated successfully"})
}

// UpdateHP updates the hero's current HP.
// PUT /api/hero/hp
//
// [Lv3 workshop task - bug plant location]
// This handler is fully implemented, but in the bug version the route
// registration in main.go is commented out, causing 404.
// Students must add the route: api.PUT("/hero/hp", heroHandler.UpdateHP)
func (h *HeroHandler) UpdateHP(c echo.Context) error {
	var req model.UpdateHPRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request body"})
	}
	if req.HP <= 0 {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "hp must be greater than 0"})
	}
	if err := h.heroService.UpdateHP(req.HP); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, map[string]string{"message": "HP updated successfully"})
}
