package models

import (
	"encoding/json"
	"time"

	"github.com/go-playground/validator/v10"

	"github.com/earaujoassis/space/internal/services/volatile"
)

// Action is a model/struct used to represent ephemeral actions/sessions in the application
type Action struct {
	UUID        string    `validate:"omitempty,uuid4" json:"uuid"`
	User        User      `validate:"required" json:"-"`
	UserID      uint      `json:"user_id"`
	Client      Client    `validate:"required" json:"-"`
	ClientID    uint      `json:"client_id"`
	Moment      int64     `json:"moment"`
	ExpiresIn   int64     `json:"expires_in"`
	IP          string    `validate:"required" json:"ip"`
	UserAgent   string    `validate:"required" json:"user_agent"`
	Token       string    `validate:"omitempty,alphanum" json:"token"`
	Scopes      string    `validate:"required,scope" json:"scopes"`
	Description string    `validate:"required,action" json:"description"`
	CreatedAt   time.Time `json:"created_at"`
}

func validAction(fl validator.FieldLevel) bool {
	description := fl.Field().String()
	if description != NotSpecialAction && description != UpdateUserAction {
		return false
	}
	return true
}

// Save saves an Action entry in a memory store (Redis)
func (action *Action) Save() error {
	action.UserID = action.User.ID
	action.ClientID = action.Client.ID
	action.UUID = generateUUID()
	action.CreatedAt = time.Now().UTC()
	action.Token = GenerateRandomString(64)
	action.Moment = time.Now().UTC().Unix()
	action.ExpiresIn = shortestExpirationLength
	if err := validateModel("validate", action); err != nil {
		return err
	}
	actionJSON, _ := json.Marshal(action)
	volatile.TransactionWrapper(func() {
		volatile.SetFieldAtKey("models.actions", action.UUID, actionJSON)
		volatile.SetFieldAtKey("models.actions.indexes", action.Token, action.UUID)
		volatile.AddToSortedSetAtKey("models.actions.rank", action.Moment, action.UUID)
	})
	return nil
}

// Delete deletes an Action entry in a memory store (Redis)
func (action *Action) Delete() {
	volatile.TransactionWrapper(func() {
		if !volatile.CheckFieldExistence("models.actions", action.UUID) {
			return
		}
		volatile.DeleteFieldAtKey("models.actions.indexes", action.Token)
		volatile.DeleteFieldAtKey("models.actions", action.UUID)
		volatile.RemoveFromSortedSetAtKey("models.actions.rank", action.UUID)
	})
}

// WithinExpirationWindow checks if an Action entry is still valid (time-based)
func (action *Action) WithinExpirationWindow() bool {
	now := time.Now().UTC().Unix()
	return now <= action.Moment + action.ExpiresIn
}

// CanUpdateUser checks if an Action description is valid for user update actions
func (action *Action) CanUpdateUser() bool {
	return action.Description == UpdateUserAction
}

// RetrieveActionByUUID obtains an Action entry from its UUID
func RetrieveActionByUUID(uuid string) Action {
	var action Action

	volatile.TransactionWrapper(func() {
		if !volatile.CheckFieldExistence("models.actions", uuid) {
			action = Action{}
			return
		}
		actionString := volatile.GetFieldAtKey("models.actions", uuid).ToString()
		if err := json.Unmarshal([]byte(actionString), &action); err != nil {
			action = Action{}
			return
		}
	})

	return action
}

// RetrieveActionByToken obtains an Action entry from its token-string
func RetrieveActionByToken(token string) Action {
	var action Action

	volatile.TransactionWrapper(func() {
		if !volatile.CheckFieldExistence("models.actions.indexes", token) {
			action = Action{}
			return
		}

		actionUUID := volatile.GetFieldAtKey("models.actions.indexes", token).ToString()
		action = RetrieveActionByUUID(actionUUID)
	})

	return action
}
