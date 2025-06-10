package helpers

import (
	"context"
	"log"

	"github.com/ory/dockertest/v3"
)

type TestResources struct {
	Pool             *dockertest.Pool
	DatabaseResource *dockertest.Resource
	MemoryResource   *dockertest.Resource
	Cancel           context.CancelFunc
}

func NewTestResources() *TestResources {
	var cancel context.CancelFunc
	var newTestResources TestResources

	_, cancel = context.WithCancel(context.Background())
	newTestResources.Cancel = cancel

	return &newTestResources
}

func (resources *TestResources) StartResources() {
	var pool *dockertest.Pool
	var err error

	pool, err = dockertest.NewPool("")
	if err != nil {
		log.Fatalf("Could not construct docker pool: %s", err)
	}
	err = pool.Client.Ping()
	if err != nil {
		log.Fatalf("Could not connect to Docker: %s", err)
	}
	resources.Pool = pool
	resources.startDatabaseResource()
	resources.startMemoryResource()
}

func (resources *TestResources) PurgeResources() {
	if err := resources.Pool.Purge(resources.DatabaseResource); err != nil {
		log.Fatalf("Could not purge database resource: %s", err)
	}
	if err := resources.Pool.Purge(resources.MemoryResource); err != nil {
		log.Fatalf("Could not purge memory resource: %s", err)
	}

	resources.Cancel()
}
