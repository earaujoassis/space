package tasks

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/earaujoassis/space/internal/config"
	"github.com/earaujoassis/space/internal/feature"
	"github.com/earaujoassis/space/internal/gateways/redis"
)

// ToggleFeature is used to enable or disable a feature-gate
func ToggleFeature(cfg *config.Config) {
	ms, _ := redis.NewMemoryService(cfg)
	fg := feature.NewFeatureGate(ms)
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Feature key: ")
	featureKey, _ := reader.ReadString('\n')
	featureKey = strings.Trim(featureKey, "\n")
	if !feature.IsFeatureAvailable(featureKey) {
		fmt.Printf("Key `%s` is not available as a feature\n", featureKey)
		os.Exit(1)
	}
	if fg.IsActive(featureKey) {
		fg.Disable(featureKey)
		fmt.Printf("Key `%s` is disabled\n", featureKey)
	} else {
		fg.Enable(featureKey)
		fmt.Printf("Key `%s` is enabled\n", featureKey)
	}
}
