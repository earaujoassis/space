package config

import (
    "os"
    "strings"
    "encoding/json"
    "io/ioutil"
    "log"

    "github.com/hashicorp/vault/api"
)

const (
    localConfigurationFile = ".config.local.json"
    configurationStoreFile = ".config.yml"
)

// Config struct with configuration data for the application
type Config struct {
    ApplicationKey string `json:"application_key"`
    DatastoreHost string `json:"datastore_host"`
    DatastoreNamePrefix string `json:"datastore_name_prefix"`
    DatastorePassword string `json:"datastore_password"`
    DatastoreSslMode string `json:"datastore_ssl_mode"`
    DatastoreUser string `json:"datastore_user"`
    MailFrom string `json:"mail_from"`
    MailerAccess string `json:"mailer_access"`
    MemorystoreHost string `json:"memory_store_host"`
    MemorystoreIndex int `json:"memory_store_index"`
    MemorystorePassword string `json:"memory_store_password"`
    MemorystorePort int `json:"memory_store_port"`
    SessionSecret string `json:"session_secret"`
    SessionSecure bool `json:"session_secure"`
    StorageSecret string `json:"storage_secret"`
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
//      it could be: development, testing, production
func Environment() string {
    return environment
}

// IsEnvironment checks if the current environment for the application
//      is the same as defined in `env`
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
//   1. it attempts to load it from the .config.local.json file;
//   2. it checks for the .config.yml file and loads it from Vault; or
//   3. it fails without any configuration option available
func LoadConfig() {
    var globalService Service
    var dataStream []byte
    var err error
    var client *api.Client
    var secret *api.Secret

    if _, jErr := os.Stat(localConfigurationFile); jErr == nil {
        // .config.local.json exists
        dataStream, err = ioutil.ReadFile(localConfigurationFile)
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
    } else {
        // no configuration option available
        log.Fatal("> No configuration option is available; fatal")
    }

    err = json.Unmarshal([]byte(dataStream), &globalConfig)
    if err != nil {
        panic(err)
    }
}
