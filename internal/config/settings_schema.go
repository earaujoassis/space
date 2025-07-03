package config

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/earaujoassis/space/internal/logs"
	"github.com/earaujoassis/space/internal/utils"
)

const (
	settingsSchemaFile = "./configs/schemas/settings.json"
)

func LoadSettingsSchema() utils.H {
	var inputSchema utils.H
	var schema utils.H

	if _, err := os.Stat(settingsSchemaFile); err != nil {
		logs.Propagatef(logs.LevelPanic, "cannot load settings schema file: %s", err.Error())
	}

	dataStream, err := os.ReadFile(settingsSchemaFile)
	if err != nil {
		logs.Propagate(logs.LevelPanic, err.Error())
	}
	err = json.Unmarshal([]byte(dataStream), &inputSchema)
	if err != nil {
		logs.Propagate(logs.LevelPanic, err.Error())
	}

	if inputSchemaType, ok := inputSchema["type"].(string); !ok || inputSchemaType != "object" {
		logs.Propagate(logs.LevelPanic, "settings schema file unrecognized")
	}
	if _, ok := inputSchema["realms"].([]interface{}); !ok {
		logs.Propagate(logs.LevelPanic, "settings schema file unrecognized")
	}

	schema = utils.H{}
	realms, ok := inputSchema["realms"].([]interface{})
	if !ok {
		logs.Propagate(logs.LevelPanic, "settings schema file unrecognized")
	}
	for _, realmInterface := range realms {
		realm := utils.H(realmInterface.(map[string]interface{}))
		realmName := realm["name"].(string)
		categoriesInterface, ok := realm["categories"].([]interface{})
		if !ok {
			logs.Propagate(logs.LevelPanic, "settings schema file unrecognized")
		}
		for _, categoryInterface := range categoriesInterface {
			category := utils.H(categoryInterface.(map[string]interface{}))
			categoryName := category["name"].(string)
			propertiesInterface, ok := category["properties"].([]interface{})
			if !ok {
				logs.Propagate(logs.LevelPanic, "settings schema file unrecognized")
			}
			for _, propertyInterface := range propertiesInterface {
				var value []interface{}
				property := utils.H(propertyInterface.(map[string]interface{}))
				propertyName := property["name"].(string)
				propertyType := property["type"].(string)
				propertyDefault, ok := property["default"]
				key := fmt.Sprintf("%s.%s.%s", realmName, categoryName, propertyName)
				if ok {
					value = []interface{}{propertyType, propertyDefault}
				} else {
					value = []interface{}{propertyType}
				}
				_, ok = schema[key]
				if ok {
					logs.Propagate(logs.LevelPanic, "settings schema has duplicated path keys")
				}
				schema[key] = value
			}
		}
	}

	return schema
}
