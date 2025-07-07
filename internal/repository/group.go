package repository

import (
	"github.com/earaujoassis/space/internal/gateways/database"
	"github.com/earaujoassis/space/internal/models"
)

type GroupRepository struct {
	*BaseRepository[models.Group]
}

func NewGroupRepository(db *database.DatabaseService) *GroupRepository {
	return &GroupRepository{
		BaseRepository: NewBaseRepository[models.Group](db),
	}
}

func (r *GroupRepository) FindOrCreate(user models.User, client models.Client) models.Group {
	var group models.Group
	r.db.GetDB().
		Preload("Client").
		Preload("User").
		Preload("User.Client").
		Preload("User.Language").
		Where("user_id = ? AND client_id = ?", user.ID, client.ID).
		First(&group)
	return group
}
