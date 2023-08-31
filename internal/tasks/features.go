package tasks

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/earaujoassis/space/internal/feature"
)

// ToggleFeature is used to enable or disable a feature-gate
func ToggleFeature() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Feature key: ")
	featureKey, _ := reader.ReadString('\n')
	featureKey = strings.Trim(featureKey, "\n")
	if !feature.IsFeatureAvailable(featureKey) {
		fmt.Printf("Key `%s` is not available as a feature\n", featureKey)
		os.Exit(1)
	}
	if feature.IsActive(featureKey) {
		feature.Disable(featureKey)
		fmt.Printf("Key `%s` is disabled\n", featureKey)
	} else {
		feature.Enable(featureKey)
		fmt.Printf("Key `%s` is enabled\n", featureKey)
	}
}
