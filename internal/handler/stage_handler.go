package handler

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/maropook/gopher-slayer/internal/service"
)

type StageHandler struct {
	stageService *service.StageService
}

func NewStageHandler(stageService *service.StageService) *StageHandler {
	return &StageHandler{stageService: stageService}
}

// GetStages returns all stages with unlock status.
// GET /api/stages
func (h *StageHandler) GetStages(c echo.Context) error {
	stages, err := h.stageService.GetAllStages()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, stages)
}

// GetEnemies returns all enemies for a given stage.
// GET /api/stages/:id/enemies
func (h *StageHandler) GetEnemies(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid stage id"})
	}
	enemies, err := h.stageService.GetEnemiesByStageID(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, enemies)
}

// ClearStage marks a stage as cleared and grants experience to the hero.
// POST /api/stages/:id/clear
func (h *StageHandler) ClearStage(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid stage id"})
	}
	result, err := h.stageService.ClearStage(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, result)
}
