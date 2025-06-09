package services

import (
	"github.com/earaujoassis/space/internal/models"
)

// CreateAction creates an ephemeral Action entry
func CreateAction(user models.User, client models.Client, ip, userAgent, scopes string, description string) models.Action {
	var action models.Action = models.Action{
		User:        user,
		Client:      client,
		IP:          ip,
		UserAgent:   userAgent,
		Scopes:      scopes,
		Description: description,
	}
	if err := action.Save(); err != nil {
		return models.Action{}
	}
	return action
}

// ActionAuthentication authenticates an ephemeral Action entry (time-based)
func ActionAuthentication(token string) models.Action {
	var action models.Action = models.RetrieveActionByToken(token)
	if action.UUID != "" && !action.WithinExpirationWindow() {
		action.Delete()
		return models.Action{}
	}
	return action
}

// ActionGrantsReadAbility checks if an action entry has read-ability
func ActionGrantsReadAbility(action models.Action) bool {
	return action.Scopes == models.ReadScope || action.Scopes == models.WriteScope || action.Scopes == models.OpenIDScope
}

// ActionGrantsWriteAbility checks if an action entry has write-ability
func ActionGrantsWriteAbility(action models.Action) bool {
	return action.Scopes == models.WriteScope
}
