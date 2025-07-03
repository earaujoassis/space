package repository

import (
	"fmt"

	"github.com/earaujoassis/space/internal/config"
	"github.com/earaujoassis/space/internal/gateways/database"
	"github.com/earaujoassis/space/internal/models"
	"github.com/earaujoassis/space/internal/utils"
)

type SettingRepository struct {
	*BaseRepository[models.Setting]
	Schema utils.H
}

func NewSettingRepository(db *database.DatabaseService) *SettingRepository {
	return &SettingRepository{
		BaseRepository: NewBaseRepository[models.Setting](db),
		Schema:         config.LoadSettingsSchema(),
	}
}

// ReduceForUser creates a map with the default settings,
//
// overridden by the user settings
func (r *SettingRepository) ReduceForUser(user models.User) utils.H {
	settings := r.GetAllForUser(user)
	schema := r.Schema
	reduced := make(utils.H)
	userSettings := make(utils.H)
	for _, setting := range settings {
		key, value := setting.Reduce()
		userSettings[key] = value
	}

	for key, info := range schema {
		propertyInfo := info.([]interface{})
		userValue, ok := userSettings[key]
		if ok {
			reduced[key] = userValue
		} else if len(propertyInfo) > 1 {
			reduced[key] = propertyInfo[1]
		}
	}

	return reduced
}

// GetAllForUser lists all stored settings for a given user
func (r *SettingRepository) GetAllForUser(user models.User) []models.Setting {
	settings := make([]models.Setting, 0)

	if err := r.db.GetDB().Where("user_id = ?", user.ID).
		Order("id ASC").
		Find(&settings).Error; err != nil {
		return make([]models.Setting, 0)
	}

	return settings
}

func (r *SettingRepository) getDefaultSetting(realm, category, property string) models.Setting {
	schema := r.Schema
	key := fmt.Sprintf("%s.%s.%s", realm, category, property)
	info, ok := schema[key].([]interface{})
	if !ok {
		return models.Setting{}
	}

	value := ""
	if len(info) > 1 {
		value = fmt.Sprintf("%v", info[1])
	}
	return models.Setting{
		Realm:    realm,
		Category: category,
		Property: property,
		Type:     info[0].(string),
		Value:    value,
	}
}

// Create creates a new setting entry
func (r *SettingRepository) Create(setting *models.Setting) error {
	defaultError := fmt.Errorf("invalid user setting")
	schema := r.Schema
	key := fmt.Sprintf("%s.%s.%s", setting.Realm, setting.Category, setting.Property)
	info, ok := schema[key].([]interface{})
	if !ok {
		return defaultError
	}
	setting.Type = info[0].(string)
	if _, err := setting.DeserializeValue(); err != nil {
		return defaultError
	}
	return r.BaseRepository.Create(setting)
}

// FindOrDefault attempts to obtain a setting for a user, or returning the default setting
func (r *SettingRepository) FindOrDefault(user models.User, realm, category, property string) models.Setting {
	var setting models.Setting
	result := r.db.GetDB().
		Preload("User").
		Preload("User.Client").
		Preload("User.Language").
		Where("user_id = ? AND realm = ? AND category = ? AND property = ?",
			user.ID, realm, category, property).
		First(&setting)
	if result.Error == nil {
		return setting
	}

	setting = r.getDefaultSetting(realm, category, property)
	setting.User = user
	return setting
}

// FindOrFallback attempts to obtain a setting for a user, or returning the given fallback
func (r *SettingRepository) FindOrFallback(user models.User, realm, category, property, fallback string) models.Setting {
	var setting models.Setting
	result := r.db.GetDB().
		Preload("User").
		Preload("User.Client").
		Preload("User.Language").
		Where("user_id = ? AND realm = ? AND category = ? AND property = ?",
			user.ID, realm, category, property).
		First(&setting)
	if result.Error == nil {
		return setting
	}

	schema := r.Schema
	key := fmt.Sprintf("%s.%s.%s", realm, category, property)
	info, ok := schema[key].([]interface{})
	if !ok {
		return models.Setting{}
	}

	return models.Setting{
		User:     user,
		Realm:    realm,
		Category: category,
		Property: property,
		Type:     info[0].(string),
		Value:    fallback,
	}
}
