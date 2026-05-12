package service

import (
	"fmt"
	"math/rand"
)

type BattleService struct{}

func NewBattleService() *BattleService {
	return &BattleService{}
}

type AttackRequest struct {
	HeroAttack int `json:"hero_attack"`
}

type AttackResponse struct {
	Damage  int    `json:"damage"`
	Message string `json:"message"`
}

type EnemyAttackRequest struct {
	EnemyAttack int    `json:"enemy_attack"`
	EnemyName   string `json:"enemy_name"`
}

// calculateDamage computes the damage dealt by an attacker.
// The damage has a slight random variance of ±20%.
//
// [Lv1 workshop task - bug plant location]
// The bug version changes this to: return 0
// Students must fix it to return the correct damage value.
func calculateDamage(attack int) int {
	if attack <= 0 {
		return 0
	}
	variance := int(float64(attack) * 0.2)
	if variance == 0 {
		return attack
	}
	// damage is in range [attack - variance, attack + variance]
	return attack - variance + rand.Intn(variance*2+1)
}

// HeroAttack calculates the damage dealt by the hero.
func (s *BattleService) HeroAttack(req AttackRequest) AttackResponse {
	damage := calculateDamage(req.HeroAttack)
	return AttackResponse{
		Damage:  damage,
		Message: fmt.Sprintf("You dealt %d damage!", damage),
	}
}

// EnemyAttack calculates the damage dealt by an enemy.
//
// [Lv4 workshop task - bug plant location]
// The bug version adds time.Sleep(3*time.Second) and negates the damage
// for a specific enemy name (e.g. "Hell Hound"), causing the hero to heal instead.
func (s *BattleService) EnemyAttack(req EnemyAttackRequest) AttackResponse {
	damage := calculateDamage(req.EnemyAttack)
	return AttackResponse{
		Damage:  damage,
		Message: fmt.Sprintf("%s dealt %d damage!", req.EnemyName, damage),
	}
}
