package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/maropook/gopher-slayer/internal/service"
)

type BattleHandler struct {
	battleService *service.BattleService
}

func NewBattleHandler(battleService *service.BattleService) *BattleHandler {
	return &BattleHandler{battleService: battleService}
}

// Attack handles the hero attacking an enemy.
// The server calculates the damage; the client tracks enemy HP.
// POST /api/battle/attack
func (h *BattleHandler) Attack(c echo.Context) error {
	var req service.AttackRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request body"})
	}
	result := h.battleService.HeroAttack(req)
	return c.JSON(http.StatusOK, result)
}

// EnemyAttack handles an enemy attacking the hero.
// The server calculates the damage; the client tracks hero HP.
// POST /api/battle/enemy-attack
func (h *BattleHandler) EnemyAttack(c echo.Context) error {
	var req service.EnemyAttackRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request body"})
	}
	result := h.battleService.EnemyAttack(req)
	return c.JSON(http.StatusOK, result)
}
