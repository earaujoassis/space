package feature

import (
	"golang.org/x/exp/slices"
)

const (
	// UserCreate make it possible to create users (through sign-up)
	UserCreate string = "user.create"
	// UserAdminify make it possible to set users as admin (using the application key)
	UserAdminify string = "user.adminify"
)

func IsFeatureAvailable(feature string) bool {
	availableFeatures := []string{UserCreate, UserAdminify}
	return slices.Contains(availableFeatures, feature)
}
