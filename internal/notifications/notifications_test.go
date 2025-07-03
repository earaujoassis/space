package notifications

import (
	"slices"
	"testing"

	"github.com/alicebob/miniredis/v2"
	"github.com/brianvoe/gofakeit/v7"
	"github.com/hibiken/asynq"
	"github.com/stretchr/testify/suite"

	"github.com/earaujoassis/space/internal/config"
	"github.com/earaujoassis/space/internal/gateways/database"
	"github.com/earaujoassis/space/internal/gateways/memory"
	"github.com/earaujoassis/space/internal/logs"
	"github.com/earaujoassis/space/internal/models"
	"github.com/earaujoassis/space/internal/repository"
	"github.com/earaujoassis/space/test/utils"
)

type NotificationsTestSuite struct {
	suite.Suite
	cfg *config.Config
	rm  *repository.RepositoryManager
	db  *database.DatabaseService
	ms  *memory.MemoryService
	mr  *miniredis.Miniredis
}

func (s *NotificationsTestSuite) SetupSuite() {
	if err := utils.EnsureProjectRoot(); err != nil {
		s.T().Fatalf("Failed to change to project root: %v", err)
	}

	provider := miniredis.NewMiniRedis()
	err := provider.StartAddr("localhost:6380")
	if err != nil {
		logs.Propagate(logs.LevelPanic, err.Error())
	}
	s.mr = provider

	utils.SetupConfigEnv()
	cfg, _ := config.Load()
	cfg.MemorystoreHost = "localhost"
	cfg.MemorystorePort = 6380
	db, err := database.NewDatabaseService(cfg)
	s.Require().NoError(err)
	if err != nil {
		s.T().Fatalf("Could not create new database service: %v", err)
	}
	s.db = db
	err = utils.RunUnitTestMigrator(db.GetDB())
	s.Require().NoError(err)
	ms, err := memory.NewMemoryService(cfg)
	s.Require().NoError(err)
	if err != nil {
		s.T().Fatalf("Could not create new memory service: %v", err)
	}
	s.ms = ms
	s.rm = repository.NewRepositoryManager(db, ms)
	s.cfg = cfg
}

func (s *NotificationsTestSuite) SetupTest() {
	s.cleanupQueue()
	s.cleanupDatabase()
	s.cleanupRedis()
}

func (s *NotificationsTestSuite) cleanupQueue() {
	inspector := asynq.NewInspector(asynq.RedisClientOpt{Addr: "localhost:6380"})
	defer inspector.Close()
	queues, err := inspector.Queues()
	if !slices.Contains(queues, "default") {
		return
	}
	s.Require().NoError(err)
	_, err = inspector.DeleteAllPendingTasks("default")
	if err != nil && err != asynq.ErrQueueNotFound {
		s.T().Fatalf("Failed to clear pending tasks: %v", err)
	}
}

func (s *NotificationsTestSuite) cleanupDatabase() {
	db, err := s.db.GetDB().DB()
	s.Require().NoError(err)
	db.Exec("DELETE FROM sessions")
	db.Exec("DELETE FROM users")
	db.Exec("DELETE FROM services")
	db.Exec("DELETE FROM clients")
	db.Exec("DELETE FROM languages")
}

func (s *NotificationsTestSuite) cleanupRedis() {
	s.ms.Do("FLUSHDB")
}

func (s *NotificationsTestSuite) TearDownSuite() {
	s.ms.Close()
	s.db.Close()
}

func (s *NotificationsTestSuite) createUser() models.User {
	repositories := s.rm

	username := gofakeit.Username()
	email := gofakeit.Email()
	user := models.User{
		FirstName:     gofakeit.FirstName(),
		LastName:      gofakeit.LastName(),
		Username:      username,
		Email:         email,
		Passphrase:    gofakeit.Password(true, true, true, true, false, 10),
		CodeSecret:    gofakeit.Password(true, true, true, true, false, 64),
		RecoverSecret: gofakeit.Password(true, true, true, true, false, 64),
	}
	user.Client = models.Client{
		Name:         gofakeit.Company(),
		Secret:       models.GenerateRandomString(64),
		CanonicalURI: []string{"localhost"},
		RedirectURI:  []string{"/"},
		Scopes:       models.PublicScope,
		Type:         models.PublicClient,
	}
	user.Language = models.Language{
		Name:    "English",
		IsoCode: "en-US",
	}
	err := repositories.Users().Create(&user)
	s.Require().NoError(err)

	return user
}

func TestRepositorySuite(t *testing.T) {
	suite.Run(t, new(NotificationsTestSuite))
}
