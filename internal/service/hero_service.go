package service

import (
	"github.com/maropook/gopher-slayer/internal/model"
	"github.com/maropook/gopher-slayer/internal/repository"
)

type HeroService struct {
	heroRepo *repository.HeroRepository
}

func NewHeroService(heroRepo *repository.HeroRepository) *HeroService {
	return &HeroService{heroRepo: heroRepo}
}

func (s *HeroService) GetHero() (*model.Hero, error) {
	return s.heroRepo.GetHero()
}

func (s *HeroService) UpdateName(name string) error {
	return s.heroRepo.UpdateName(name)
}

func (s *HeroService) UpdateExperience(experience int) error {
	return s.heroRepo.UpdateExperience(experience)
}

func (s *HeroService) UpdateHP(hp int) error {
	return s.heroRepo.UpdateHP(hp)
}
