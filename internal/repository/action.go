package repository

import (
	"encoding/json"

	"github.com/earaujoassis/space/internal/gateways/memory"
	"github.com/earaujoassis/space/internal/models"
)

type ActionRepository struct {
	*BaseMemoryRepository[models.Action]
}

func NewActionRepository(ms *memory.MemoryService) *ActionRepository {
	return &ActionRepository{
		BaseMemoryRepository: NewBaseMemoryRepository[models.Action](ms),
	}
}

func (r *ActionRepository) Create(action *models.Action) error {
	action.BeforeSave()
	if err := action.Validate(); err != nil {
		return err
	}
	actionJSON, _ := json.Marshal(action)
	r.ms.Transaction(func(c *memory.Commands) {
		c.SetFieldAtKey("models.actions", action.UUID, actionJSON)
		c.SetFieldAtKey("models.actions.indexes", action.Token, action.UUID)
		c.AddToSortedSetAtKey("models.actions.rank", action.Moment, action.UUID)
	})

	return nil
}

func (r *ActionRepository) Delete(action models.Action) {
	r.ms.Transaction(func(c *memory.Commands) {
		c.DeleteFieldAtKey("models.actions.indexes", action.Token)
		c.DeleteFieldAtKey("models.actions", action.UUID)
		c.RemoveFromSortedSetAtKey("models.actions.rank", action.UUID)
	})
}

func (r *ActionRepository) FindByUUID(uuid string) models.Action {
	var action models.Action

	r.ms.Transaction(func(c *memory.Commands) {
		if !c.CheckFieldExistence("models.actions", uuid) {
			action = models.Action{}
			return
		}
		actionString := c.GetFieldAtKey("models.actions", uuid).ToString()
		if err := json.Unmarshal([]byte(actionString), &action); err != nil {
			action = models.Action{}
			return
		}
	})

	return action
}

func (r *ActionRepository) FindByToken(token string) models.Action {
	var action models.Action

	r.ms.Transaction(func(c *memory.Commands) {
		if !c.CheckFieldExistence("models.actions.indexes", token) {
			action = models.Action{}
			return
		}

		actionUUID := c.GetFieldAtKey("models.actions.indexes", token).ToString()
		action = r.FindByUUID(actionUUID)
	})

	return action
}

func (r *ActionRepository) Authentication(token string) models.Action {
	action := r.FindByToken(token)
	if action.UUID != "" && !action.WithinExpirationWindow() {
		r.Delete(action)
		return models.Action{}
	}
	return action
}
