package utils

import (
	"os"
)

func SetupConfigEnv() {
	os.Setenv("SPACE_ENV", "test")
	os.Setenv("SPACE_APPLICATION_KEY", "masterapplicationkey")
	os.Setenv("SPACE_MAIL_FROM", "example@example.com")
	os.Setenv("SPACE_MAILER_ACCESS", "AccessKeyId:SecretAccessKey:Region")
	os.Setenv("SPACE_SESSION_SECRET", "E93jykumzKrJOp6xKB4JduxaKLmeiPmf")
	os.Setenv("SPACE_SESSION_SECURE", "false")
	os.Setenv("SPACE_STORAGE_SECRET", "KRgwMcZdLPfo9bck")
	os.Setenv("SPACE_DATASTORE_HOST", ":memory:")
	os.Setenv("SPACE_DATASTORE_PORT", "")
	os.Setenv("SPACE_DATASTORE_NAME_PREFIX", "")
	os.Setenv("SPACE_DATASTORE_USER", "")
	os.Setenv("SPACE_DATASTORE_PASSWORD", "")
	os.Setenv("SPACE_DATASTORE_SSL_MODE", "")
	os.Setenv("SPACE_MEMORY_STORE_HOST", ":memory:")
	os.Setenv("SPACE_MEMORY_STORE_PORT", "")
	os.Setenv("SPACE_MEMORY_STORE_INDEX", "")
	os.Setenv("SPACE_MEMORY_STORE_PASSWORD", "")
}
