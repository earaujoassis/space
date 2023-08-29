package config

import (
	"encoding/json"
	"log"
	"os"
	"strings"

	"github.com/caarlos0/env/v6"
	"github.com/hashicorp/vault/api"
	"github.com/joho/godotenv"
)

const (
	localConfigurationFile = ".config.local.json"
	configurationStoreFile = ".config.yml"
)

// Config struct with configuration data for the application
type Config struct {
	ApplicationKey      string `json:"application_key" env:"SPACE_APPLICATION_KEY"`
	DatastoreHost       string `json:"datastore_host" env:"SPACE_DATASTORE_HOST"`
	DatastorePort       int    `json:"datastore_port" env:"SPACE_DATASTORE_PORT"`
	DatastoreNamePrefix string `json:"datastore_name_prefix" env:"SPACE_DATASTORE_NAME_PREFIX"`
	DatastorePassword   string `json:"datastore_password" env:"SPACE_DATASTORE_PASSWORD"`
	DatastoreSslMode    string `json:"datastore_ssl_mode" env:"SPACE_DATASTORE_SSL_MODE"`
	DatastoreUser       string `json:"datastore_user" env:"SPACE_DATASTORE_USER"`
	MailFrom            string `json:"mail_from" env:"SPACE_MAIL_FROM"`
	MailerAccess        string `json:"mailer_access" env:"SPACE_MAILER_ACCESS"`
	MemorystoreHost     string `json:"memory_store_host" env:"SPACE_MEMORY_STORE_HOST"`
	MemorystoreIndex    int    `json:"memory_store_index" env:"SPACE_MEMORY_STORE_INDEX"`
	MemorystorePassword string `json:"memory_store_password" env:"SPACE_MEMORY_STORE_PASSWORD"`
	MemorystorePort     int    `json:"memory_store_port" env:"SPACE_MEMORY_STORE_PORT"`
	SessionSecret       string `json:"session_secret" env:"SPACE_SESSION_SECRET"`
	SessionSecure       bool   `json:"session_secure" env:"SPACE_SESSION_SECURE"`
	StorageSecret       string `json:"storage_secret" env:"SPACE_STORAGE_SECRET"`
}

var globalConfig Config
var environment string

func init() {
	environment = strings.ToLower(os.Getenv("SPACE_ENV"))
	if environment == "" {
		environment = "development"
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

	if _, jErr := os.Stat(localConfigurationFile); jErr == nil {
		// .config.local.json exists
		dataStream, err = os.ReadFile(localConfigurationFile)
		if err != nil {
			panic(err)
		}
		err = json.Unmarshal([]byte(dataStream), &globalConfig)
		if err != nil {
			panic(err)
		}
	} else if _, yErr := os.Stat(configurationStoreFile); yErr == nil && os.IsNotExist(jErr) {
		// .config.yml exists
		globalService.LoadService(configurationStoreFile)
		client, err = api.NewClient(&api.Config{
			Address: globalService.Space.ConfigurationStore.Addr,
		})
		if err != nil {
			panic(err)
		}
		client.SetToken(globalService.Space.ConfigurationStore.Token)
		secret, err = client.Logical().Read(globalService.Space.ConfigurationStore.Path)
		if err != nil {
			panic(err)
		}

		dataStream, _ = json.Marshal(secret.Data)
		err = json.Unmarshal([]byte(dataStream), &globalConfig)
		if err != nil {
			panic(err)
		}
	} else if err := godotenv.Load(); err == nil {
		err = env.Parse(&globalConfig)
		if err != nil {
			panic(err)
		}
	} else if err = env.Parse(&globalConfig); err == nil {
		log.Println("> Configuration obtained from environment; all good")
	} else {
		// no configuration option available
		log.Fatal("> No configuration option is available; fatal")
	}
}
