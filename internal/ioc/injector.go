package ioc

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/earaujoassis/space/internal/config"
	"github.com/earaujoassis/space/internal/feature"
	"github.com/earaujoassis/space/internal/gateways/memory"
	"github.com/earaujoassis/space/internal/notifications"
	"github.com/earaujoassis/space/internal/policy"
	"github.com/earaujoassis/space/internal/repository"
)

const AppContextKey = "app_context"

func InjectAppContext(appCtx *AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set(AppContextKey, appCtx)
		c.Next()
	}
}

func GetAppContext(c *gin.Context) *AppContext {
	return c.MustGet(AppContextKey).(*AppContext)
}

func GetConfig(c *gin.Context) *config.Config {
	return GetAppContext(c).Config
}

func GetDB(c *gin.Context) *gorm.DB {
	return GetAppContext(c).DB.GetDB()
}

func GetMemoryService(c *gin.Context) *memory.MemoryService {
	return GetAppContext(c).Memory
}

func GetRepositories(c *gin.Context) *repository.RepositoryManager {
	return GetAppContext(c).Repositories
}

func GetFeatureGate(c *gin.Context) *feature.FeatureGate {
	return GetAppContext(c).FeatureGate
}

func GetRateLimitService(c *gin.Context) *policy.RateLimitService {
	return GetAppContext(c).RateLimit
}

func GetNotifier(c *gin.Context) *notifications.Notifier {
	return GetAppContext(c).Notifier
}
