package models

import (
	"fmt"
	"net/url"
	"reflect"
	"slices"
	"strings"

	"github.com/go-playground/validator/v10"
)

func validAction(fl validator.FieldLevel) bool {
	description := fl.Field().String()
	if description != NotSpecialAction && description != UpdateUserAction {
		return false
	}
	return true
}

func validClientType(fl validator.FieldLevel) bool {
	typeField := fl.Field().String()
	if typeField != PublicClient && typeField != ConfidentialClient {
		return false
	}
	return true
}

func validCanonicalURIs(fl validator.FieldLevel) bool {
	// It's not a Client model
	if !fl.Top().CanConvert(reflect.TypeOf(Client{})) {
		return true
	}

	currentClient := fl.Top().Interface().(Client)
	// The Jupiter (internal client) is created with the following value
	if len(currentClient.CanonicalURI) == 1 && currentClient.CanonicalURI[0] == "localhost" {
		return true
	}

	for i := range currentClient.CanonicalURI {
		currentEntry := currentClient.CanonicalURI[i]
		u, err := url.Parse(currentEntry)
		if err != nil {
			return false
		}

		if !strings.Contains(u.Scheme, "http") || u.Path != "" || u.Host == "" {
			return false
		}
	}

	return true
}

func validRedirectURIs(fl validator.FieldLevel) bool {
	// It's not a Client model
	if !fl.Top().CanConvert(reflect.TypeOf(Client{})) {
		return true
	}

	currentClient := fl.Top().Interface().(Client)
	canonicalUri := currentClient.CanonicalURI
	redirectUri := currentClient.RedirectURI

	// The Jupiter (internal client) is created with the following values
	if len(canonicalUri) == 1 && canonicalUri[0] == "localhost" && len(redirectUri) == 1 && redirectUri[0] == "/" {
		return true
	}

	for i := range redirectUri {
		currentEntry := redirectUri[i]
		u, err := url.Parse(currentEntry)
		if err != nil {
			return false
		}
		currentCanonical := fmt.Sprintf("%s://%s", u.Scheme, u.Host)
		if !slices.Contains(canonicalUri, currentCanonical) {
			return false
		}
	}

	return true
}

func validClientScopes(fl validator.FieldLevel) bool {
	scopesField := fl.Field().String()

	// TODO A PublicClient can't have a ReadScope
	// It's not a Client model
	/*
		if !fl.Top().CanConvert(reflect.TypeOf(Client{})) {
			return true
		}
		currentClient := fl.Top().Interface().(Client)
	*/

	if !HasValidScopes(strings.Split(scopesField, " ")) {
		return false
	}
	return true
}

func validServiceType(fl validator.FieldLevel) bool {
	typeField := fl.Field().String()
	if typeField != PublicService && typeField != AttachedService {
		return false
	}
	return true
}

func validScope(fl validator.FieldLevel) bool {
	scope := fl.Field().String()
	if scope != PublicScope && scope != ReadScope && scope != WriteScope && scope != OpenIDScope {
		return false
	}
	return true
}

func validTokenType(fl validator.FieldLevel) bool {
	tokenType := fl.Field().String()
	if tokenType != AccessToken && tokenType != RefreshToken && tokenType != GrantToken && tokenType != IDToken {
		return false
	}
	return true
}

func validateModel(tagName string, model interface{}) error {
	validate := validator.New()
	validate.SetTagName(tagName)
	validate.RegisterValidation("client", validClientType)
	validate.RegisterValidation("scope", validScope)
	validate.RegisterValidation("restrict", validClientScopes)
	validate.RegisterValidation("token", validTokenType)
	validate.RegisterValidation("canonical", validCanonicalURIs)
	validate.RegisterValidation("redirect", validRedirectURIs)
	validate.RegisterValidation("action", validAction)
	validate.RegisterValidation("service", validServiceType)
	err := validate.Struct(model)
	return err
}

// IsValid checks if a `model` entry is valid, given the `tagName` (scope) for validation
func IsValid(tagName string, model interface{}) bool {
	err := validateModel(tagName, model)
	return err == nil
}
