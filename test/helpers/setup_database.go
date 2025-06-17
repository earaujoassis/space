package helpers

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/lib/pq"
	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
)

func (resources *TestResources) startDatabaseResource() {
	if resources.Pool == nil {
		log.Fatalf("Resources pool has not been started")
		return
	}
	databaseResource, err := resources.Pool.RunWithOptions(&dockertest.RunOptions{
		Repository: "postgres",
		Tag:        "14.18-alpine3.22",
		Hostname:   "postgres",
		Env: []string{
			"listen_addresses='*'",
			"POSTGRES_DB=space_integration",
			"POSTGRES_HOST_AUTH_METHOD=trust",
			"POSTGRES_PASSWORD=password",
			"POSTGRES_USER=user",
		},
	}, func(config *docker.HostConfig) {
		config.AutoRemove = true
		config.RestartPolicy = docker.RestartPolicy{Name: "no"}
	})
	if err != nil {
		log.Fatalf("Could not start database resource: %s", err)
	}
	resources.DatabaseResource = databaseResource
	err = databaseResource.Expire(600)
	if err != nil {
		log.Fatalf("Could not setup expire time: %s", err)
	}
	hostAndPort := databaseResource.GetHostPort("5432/tcp")
	dbUrl := fmt.Sprintf("postgres://user:password@%s/space_integration?sslmode=disable", hostAndPort)
	os.Setenv("SPACE_DATASTORE_HOST", "localhost")
	os.Setenv("SPACE_DATASTORE_PORT", databaseResource.GetPort("5432/tcp"))
	os.Setenv("SPACE_DATASTORE_NAME_PREFIX", "space")
	os.Setenv("SPACE_DATASTORE_USER", "user")
	os.Setenv("SPACE_DATASTORE_PASSWORD", "password")
	os.Setenv("SPACE_DATASTORE_SSL_MODE", "disable")
	resources.Pool.MaxWait = 120 * time.Second
	if err = resources.Pool.Retry(func() error {
		db, err := sql.Open("postgres", dbUrl)
		if err != nil {
			return err
		}
		return db.Ping()
	}); err != nil {
		log.Fatalf("Could not connect to database resource: %s", err)
	}
	log.Printf("Database URL: %s\n", dbUrl)
}
