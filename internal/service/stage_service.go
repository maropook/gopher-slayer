package service

import (
	"fmt"

	"github.com/maropook/gopher-slayer/internal/model"
	"github.com/maropook/gopher-slayer/internal/repository"
)

type StageService struct {
	stageRepo *repository.StageRepository
	heroRepo  *repository.HeroRepository
}

func NewStageService(stageRepo *repository.StageRepository, heroRepo *repository.HeroRepository) *StageService {
	return &StageService{stageRepo: stageRepo, heroRepo: heroRepo}
}

// GetAllStages returns all stages with unlock status computed from hero experience.
func (s *StageService) GetAllStages() ([]*model.Stage, error) {
	hero, err := s.heroRepo.GetHero()
	if err != nil {
		return nil, err
	}

	stages, err := s.stageRepo.GetAllStages()
	if err != nil {
		return nil, err
	}

	// is_unlocked is computed here, not stored in DB
	for _, stage := range stages {
		stage.IsUnlocked = hero.Experience >= stage.RequiredExperience
	}
	return stages, nil
}

// GetEnemiesByStageID returns all enemies for a given stage.
func (s *StageService) GetEnemiesByStageID(stageID int) ([]*model.Enemy, error) {
	return s.stageRepo.GetEnemiesByStageID(stageID)
}

// ClearStage grants experience to the hero after clearing a stage.
//
// [Lv2 workshop task - bug plant location]
// The bug version removes the heroRepo.UpdateExperience call below.
// Students must add the DB update call back to fix the game.
func (s *StageService) ClearStage(stageID int) (*model.ClearStageResponse, error) {
	// 1. Get the stage
	stage, err := s.stageRepo.GetStageByID(stageID)
	if err != nil {
		return nil, fmt.Errorf("stage not found: %w", err)
	}

	// 2. Get current hero
	hero, err := s.heroRepo.GetHero()
	if err != nil {
		return nil, fmt.Errorf("failed to get hero: %w", err)
	}

	// 3. Calculate total experience reward from all enemies in this stage
	expGained, err := s.stageRepo.GetTotalExperienceRewardForStage(stageID)
	if err != nil {
		return nil, fmt.Errorf("failed to calculate experience: %w", err)
	}

	newExp := hero.Experience + expGained

	// 4. Save the new experience to DB
	// [Lv2 bug plant: delete this block to break the game]
	if err := s.heroRepo.UpdateExperience(newExp); err != nil {
		return nil, fmt.Errorf("failed to update experience: %w", err)
	}

	return &model.ClearStageResponse{
		Message:          fmt.Sprintf("Stage '%s' cleared!", stage.Name),
		ExperienceGained: expGained,
		NewExperience:    newExp,
	}, nil
}
