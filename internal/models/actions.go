package models

import (
	// "encoding/json"
	"time"
)

// Action is a model/struct used to represent ephemeral actions/sessions in the application
type Action struct {
	UUID        string `validate:"omitempty,uuid4" json:"uuid"`
	User        User   `validate:"required" json:"-"`
	UserID      uint   `json:"user_id"`
	Client      Client `validate:"required" json:"-"`
	ClientID    uint   `json:"client_id"`
	Moment      int64  `json:"moment"`
	ExpiresIn   int64  `json:"expires_in"`
	IP          string `validate:"required" json:"ip"`
	UserAgent   string `validate:"required" json:"user_agent"`
	Token       string `validate:"omitempty,alphanum" json:"token"`
	Scopes      string `validate:"required,scope" json:"scopes"`
	Description string `validate:"required,action" json:"description"`
	Payload     string `json:"payload"`
}

func (action *Action) Validate() error {
	return validateModel("validate", action)
}

// BeforeSave sets defaults values for fields in action token
func (action *Action) BeforeSave() {
	action.UserID = action.User.ID
	action.ClientID = action.Client.ID
	action.UUID = generateUUID()
	action.Token = GenerateRandomString(64)
	action.Moment = time.Now().UTC().Unix()
	action.ExpiresIn = shortestExpirationLength
}

// WithinExpirationWindow checks if an Action entry is still valid (time-based)
func (action *Action) WithinExpirationWindow() bool {
	now := time.Now().UTC().Unix()
	return now <= action.Moment+action.ExpiresIn
}

// CanUpdateUser checks if an Action description is valid for user update actions
func (action *Action) CanUpdateUser() bool {
	return action.Description == UpdateUserAction
}

// ActionGrantsReadAbility checks if an action entry has read-ability
func (action *Action) GrantsReadAbility() bool {
	return action.Scopes == ReadScope || action.Scopes == WriteScope || action.Scopes == OpenIDScope
}

// ActionGrantsWriteAbility checks if an action entry has write-ability
func (action *Action) GrantsWriteAbility() bool {
	return action.Scopes == WriteScope
}
