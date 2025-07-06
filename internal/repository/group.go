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
