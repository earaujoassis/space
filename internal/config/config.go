package config

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/caarlos0/env/v9"
	"github.com/hashicorp/vault/api"

	"github.com/earaujoassis/space/internal"
	"github.com/earaujoassis/space/internal/logs"
)

const (
	localConfigurationFile = ".config.local.json"
	configurationStoreFile = ".config.yml"
)

// Config struct with configuration data for the application
type Config struct {
	ApplicationKey      string `json:"application_key" env:"SPACE_APPLICATION_KEY,unset"`
	DatastoreHost       string `json:"datastore_host" env:"SPACE_DATASTORE_HOST"`
	DatastorePort       int    `json:"datastore_port" env:"SPACE_DATASTORE_PORT"`
	DatastoreNamePrefix string `json:"datastore_name_prefix" env:"SPACE_DATASTORE_NAME_PREFIX"`
	DatastoreUser       string `json:"datastore_user" env:"SPACE_DATASTORE_USER"`
	DatastorePassword   string `json:"datastore_password" env:"SPACE_DATASTORE_PASSWORD,unset"`
	DatastoreSslMode    string `json:"datastore_ssl_mode" env:"SPACE_DATASTORE_SSL_MODE"`
	MailFrom            string `json:"mail_from" env:"SPACE_MAIL_FROM"`
	MailerAccess        string `json:"mailer_access" env:"SPACE_MAILER_ACCESS,unset"`
	MemorystoreHost     string `json:"memory_store_host" env:"SPACE_MEMORY_STORE_HOST"`
	MemorystorePort     int    `json:"memory_store_port" env:"SPACE_MEMORY_STORE_PORT"`
	MemorystoreIndex    int    `json:"memory_store_index" env:"SPACE_MEMORY_STORE_INDEX"`
	MemorystorePassword string `json:"memory_store_password" env:"SPACE_MEMORY_STORE_PASSWORD,unset"`
	SessionSecret       string `json:"session_secret" env:"SPACE_SESSION_SECRET,unset"`
	SessionSecure       bool   `json:"session_secure" env:"SPACE_SESSION_SECURE,unset"`
	StorageSecret       string `json:"storage_secret" env:"SPACE_STORAGE_SECRET,unset"`
	SentryUrl           string `json:"sentry_url" env:"SPACE_SENTRY_URL,unset" envDefault:""`
}

var globalConfig Config
var environment string

// LoadEnvironment loads the environment from the SPACE_ENV env var;
//
//	it could be: development, testing, production
func LoadEnvironment() {
	environment = strings.ToLower(os.Getenv("SPACE_ENV"))
	if environment == "" {
		environment = "development"
	}
	if environment != "production" && environment != "testing" && environment != "development" {
		environment = ""
	}
}

// Environment returns the current environment for the application;
//
//	it could be: development, testing, production
func Environment() string {
	return environment
}

// IsEnvironment checks if the current environment for the application
//
//	is the same as defined in `env`
func IsEnvironment(env string) bool {
	return strings.ToLower(env) == Environment()
}

func Release() string {
	if commitHash := GetEnvVar("COMMIT_HASH"); commitHash != "" {
		return fmt.Sprintf("%s+%s", internal.Version, GetEnvVar("COMMIT_HASH"))
	}
	return internal.Version
}

// GetEnvVar gets a `key` from the environment variables
func GetEnvVar(key string) string {
	return os.Getenv(key)
}

// GetGlobalConfig returns the global configuration struct for the application
func GetGlobalConfig() Config {
	return globalConfig
}

// SetConfig sets the global configuration struct for the application
func SetConfig(config Config) {
	globalConfig = config
}

// LoadConfig loads the globalConfig structure from a JSON-based stream:
//  1. it attempts to load it from the .config.local.json file;
//  2. it checks for the .config.yml file and loads it from Vault; or
//  3. it attempts to load it from .env and the environment;
//  4. it attempts to load it from the environment, directly (no .env file); or
//  5. it fails without any configuration option available
func LoadConfig() {
	var globalService Service
	var dataStream []byte
	var err error
	var client *api.Client
	var secret *api.Secret
	var loadFromEnvVarsFlag bool = false
	var loadedFlag bool = false

	LoadEnvironment()

	if environment == "testing" {
		loadFromEnvVarsFlag = true
	}

	if _, jErr := os.Stat(localConfigurationFile); !loadFromEnvVarsFlag && jErr == nil {
		// .config.local.json exists
		dataStream, err = os.ReadFile(localConfigurationFile)
		if err != nil {
			logs.Propagate(logs.Panic, err.Error())
		}
		err = json.Unmarshal([]byte(dataStream), &globalConfig)
		if err != nil {
			logs.Propagate(logs.Panic, err.Error())
		}
		loadedFlag = true
		logs.Propagatef(logs.Info, "Configuration obtained from %s; all good\n", localConfigurationFile)
	} else if _, yErr := os.Stat(configurationStoreFile); !loadFromEnvVarsFlag && yErr == nil && os.IsNotExist(jErr) {
		// .config.yml exists
		globalService.LoadService(configurationStoreFile)
		client, err = api.NewClient(&api.Config{
			Address: globalService.Space.ConfigurationStore.Addr,
		})
		if err != nil {
			logs.Propagate(logs.Panic, err.Error())
		}
		client.SetToken(globalService.Space.ConfigurationStore.Token)
		secret, err = client.Logical().Read(globalService.Space.ConfigurationStore.Path)
		if err != nil {
			logs.Propagate(logs.Panic, err.Error())
		}

		dataStream, _ = json.Marshal(secret.Data)
		err = json.Unmarshal([]byte(dataStream), &globalConfig)
		if err != nil {
			logs.Propagate(logs.Panic, err.Error())
		}
		logs.Propagate(logs.Info, "Configuration obtained from Vault; all good")
		loadedFlag = true
	} else {
		loadFromEnvVarsFlag = true
	}

	if loadFromEnvVarsFlag {
		opts := env.Options{RequiredIfNoDef: true}
		if err = env.ParseWithOptions(&globalConfig, opts); err == nil {
			loadedFlag = true
			logs.Propagate(logs.Info, "Configuration obtained from environment; all good")
		} else {
			logs.Propagatef(logs.Error, "Cannot load configuration from environment: %s", err.Error())
		}
	}

	if !loadedFlag {
		// no configuration option available
		logs.Propagate(logs.Panic, "Application is not configured")
	}
}
