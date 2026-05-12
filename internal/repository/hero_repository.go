package repository

import (
	"database/sql"

	"github.com/maropook/gopher-slayer/internal/model"
)

type HeroRepository struct {
	db *sql.DB
}

func NewHeroRepository(db *sql.DB) *HeroRepository {
	return &HeroRepository{db: db}
}

// GetHero returns the single hero (id=1).
// This game has only one hero, so we always query by id=1.
func (r *HeroRepository) GetHero() (*model.Hero, error) {
	hero := &model.Hero{}
	row := r.db.QueryRow(`
		SELECT id, name, hp, max_hp, attack, level, experience
		FROM heroes
		WHERE id = 1
	`)
	err := row.Scan(&hero.ID, &hero.Name, &hero.HP, &hero.MaxHP, &hero.Attack, &hero.Level, &hero.Experience)
	if err != nil {
		return nil, err
	}
	return hero, nil
}

// UpdateName updates the hero's name.
// This is an example of a DB update — students can reference this for Lv2.
func (r *HeroRepository) UpdateName(name string) error {
	_, err := r.db.Exec(`UPDATE heroes SET name = ? WHERE id = 1`, name)
	return err
}

// UpdateExperience updates the hero's experience points.
// This is called in the Lv2 workshop task (ClearStage).
func (r *HeroRepository) UpdateExperience(experience int) error {
	_, err := r.db.Exec(`UPDATE heroes SET experience = ? WHERE id = 1`, experience)
	return err
}

// UpdateHP updates the hero's current HP.
// This is the target of the Lv3 workshop task (students add the route for this).
func (r *HeroRepository) UpdateHP(hp int) error {
	_, err := r.db.Exec(`UPDATE heroes SET hp = ? WHERE id = 1`, hp)
	return err
}
