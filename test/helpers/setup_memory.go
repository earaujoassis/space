package helpers

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gomodule/redigo/redis"
	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
)

func (resources *TestResources) startMemoryResource() {
	if resources.Pool == nil {
		log.Fatalf("Resources pool has not been started")
		return
	}
	memoryResource, err := resources.Pool.RunWithOptions(&dockertest.RunOptions{
		Repository:   "redis",
		Tag:          "7.4-alpine3.21",
		Hostname:     "redis",
	}, func(config *docker.HostConfig) {
		config.AutoRemove = true
		config.RestartPolicy = docker.RestartPolicy{Name: "no"}
	})
	if err != nil {
		log.Fatalf("Could not start memory resource: %s", err)
	}
	resources.MemoryResource = memoryResource
	err = memoryResource.Expire(600)
	if err != nil {
		log.Fatalf("Could not setup expire time: %s", err)
	}
	hostAndPort := memoryResource.GetHostPort("6379/tcp")
	memoryUrl := fmt.Sprintf("redis://%s/0", hostAndPort)
	os.Setenv("SPACE_MEMORY_STORE_HOST", "localhost")
	os.Setenv("SPACE_MEMORY_STORE_PORT", memoryResource.GetPort("6379/tcp"))
	os.Setenv("SPACE_MEMORY_STORE_INDEX", "0")
	os.Setenv("SPACE_MEMORY_STORE_PASSWORD", "")
	resources.Pool.MaxWait = 120 * time.Second
	if err = resources.Pool.Retry(func() error {
		memoryStore, err := redis.DialURL(memoryUrl)
		if err != nil {
			return err
		}
		defer memoryStore.Close()
		_, err = memoryStore.Do("PING")
		return err
	}); err != nil {
		log.Fatalf("Could not connect to memory resource: %s", err)
	}
	log.Printf("Memory URL: %s\n", memoryUrl)
}
