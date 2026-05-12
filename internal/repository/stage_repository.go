package repository

import (
	"database/sql"

	"github.com/maropook/gopher-slayer/internal/model"
)

type StageRepository struct {
	db *sql.DB
}

func NewStageRepository(db *sql.DB) *StageRepository {
	return &StageRepository{db: db}
}

// GetAllStages returns all stages ordered by order_num.
func (r *StageRepository) GetAllStages() ([]*model.Stage, error) {
	rows, err := r.db.Query(`
		SELECT id, name, description, required_experience, order_num
		FROM stages
		ORDER BY order_num ASC
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var stages []*model.Stage
	for rows.Next() {
		s := &model.Stage{}
		if err := rows.Scan(&s.ID, &s.Name, &s.Description, &s.RequiredExperience, &s.OrderNum); err != nil {
			return nil, err
		}
		stages = append(stages, s)
	}
	return stages, rows.Err()
}

// GetStageByID returns a single stage by ID.
func (r *StageRepository) GetStageByID(id int) (*model.Stage, error) {
	s := &model.Stage{}
	row := r.db.QueryRow(`
		SELECT id, name, description, required_experience, order_num
		FROM stages
		WHERE id = ?
	`, id)
	err := row.Scan(&s.ID, &s.Name, &s.Description, &s.RequiredExperience, &s.OrderNum)
	if err != nil {
		return nil, err
	}
	return s, nil
}

// GetEnemiesByStageID returns all enemies for a given stage.
func (r *StageRepository) GetEnemiesByStageID(stageID int) ([]*model.Enemy, error) {
	rows, err := r.db.Query(`
		SELECT id, stage_id, name, hp, max_hp, attack, experience_reward
		FROM enemies
		WHERE stage_id = ?
		ORDER BY id ASC
	`, stageID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var enemies []*model.Enemy
	for rows.Next() {
		e := &model.Enemy{}
		if err := rows.Scan(&e.ID, &e.StageID, &e.Name, &e.HP, &e.MaxHP, &e.Attack, &e.ExperienceReward); err != nil {
			return nil, err
		}
		enemies = append(enemies, e)
	}
	return enemies, rows.Err()
}

// GetTotalExperienceRewardForStage returns the sum of experience_reward for all enemies in a stage.
func (r *StageRepository) GetTotalExperienceRewardForStage(stageID int) (int, error) {
	var total int
	row := r.db.QueryRow(`
		SELECT COALESCE(SUM(experience_reward), 0)
		FROM enemies
		WHERE stage_id = ?
	`, stageID)
	if err := row.Scan(&total); err != nil {
		return 0, err
	}
	return total, nil
}
