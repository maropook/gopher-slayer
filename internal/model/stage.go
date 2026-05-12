package model

type Stage struct {
	ID                 int    `json:"id"`
	Name               string `json:"name"`
	Description        string `json:"description"`
	RequiredExperience int    `json:"required_experience"`
	OrderNum           int    `json:"order_num"`
	// IsUnlocked is computed in the service layer (not stored in DB).
	// true when hero.Experience >= stage.RequiredExperience
	IsUnlocked bool `json:"is_unlocked"`
}

type ClearStageResponse struct {
	Message          string `json:"message"`
	ExperienceGained int    `json:"experience_gained"`
	NewExperience    int    `json:"new_experience"`
}
