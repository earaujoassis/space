package notifications

import (
	"github.com/hibiken/asynq"

	"github.com/earaujoassis/space/internal/utils"
)

func (s *NotificationsTestSuite) TestNotificationsRepository__AnnounceWithDefaultSettings() {
	notifier := NewNotifier(s.cfg, s.rm)
	inspector := asynq.NewInspector(asynq.RedisClientOpt{Addr: "localhost:6380"})
	defer inspector.Close()
	user := s.createUser()

	notifier.Announce(user, "user.created", utils.H{})
	queueInfo, err := inspector.GetQueueInfo("default")
	s.Require().NoError(err)
	s.Equal(1, queueInfo.Pending)
	s.cleanupQueue()

	notifier.Announce(user, "user.created", utils.H{})
	queueInfo, err = inspector.GetQueueInfo("default")
	s.Require().NoError(err)
	s.Equal(1, queueInfo.Pending)
	s.cleanupQueue()

	notifier.Announce(user, "user.update_password", utils.H{})
	queueInfo, err = inspector.GetQueueInfo("default")
	s.Require().NoError(err)
	s.Equal(1, queueInfo.Pending)
	s.cleanupQueue()

	notifier.Announce(user, "user.update_secrets", utils.H{})
	queueInfo, err = inspector.GetQueueInfo("default")
	s.Require().NoError(err)
	s.Equal(1, queueInfo.Pending)
	s.cleanupQueue()

	notifier.Announce(user, "user.email_verification", utils.H{})
	queueInfo, err = inspector.GetQueueInfo("default")
	s.Require().NoError(err)
	s.Equal(1, queueInfo.Pending)
	s.cleanupQueue()

	notifier.Announce(user, "session.created", utils.H{})
	queueInfo, err = inspector.GetQueueInfo("default")
	s.Require().NoError(err)
	s.Equal(1, queueInfo.Pending)
	s.cleanupQueue()

	notifier.Announce(user, "session.magic", utils.H{})
	queueInfo, err = inspector.GetQueueInfo("default")
	s.Require().NoError(err)
	s.Equal(1, queueInfo.Pending)
	s.cleanupQueue()

	notifier.Announce(user, "client.token_introspection", utils.H{})
	queueInfo, err = inspector.GetQueueInfo("default")
	s.Require().NoError(err)
	s.Equal(0, queueInfo.Pending)
	s.cleanupQueue()

	notifier.Announce(user, "client.userinfo_introspection", utils.H{})
	queueInfo, err = inspector.GetQueueInfo("default")
	s.Require().NoError(err)
	s.Equal(0, queueInfo.Pending)
	s.cleanupQueue()

	notifier.Announce(user, "user.authorization_granted", utils.H{})
	queueInfo, err = inspector.GetQueueInfo("default")
	s.Require().NoError(err)
	s.Equal(0, queueInfo.Pending)
	s.cleanupQueue()
}

func (s *NotificationsTestSuite) TestNotificationsRepository__AnnounceWithDefinedSettings() {
	notifier := NewNotifier(s.cfg, s.rm)
	inspector := asynq.NewInspector(asynq.RedisClientOpt{Addr: "localhost:6380"})
	defer inspector.Close()
	user := s.createUser()
	repositories := s.rm

	setting := repositories.Settings().FindOrDefault(user, "notifications", "client-application-email-notifications", "token-introspection")
	setting.Value = "true"
	err := repositories.Settings().Create(&setting)
	s.Require().NoError(err)
	notifier.Announce(user, "client.token_introspection", utils.H{})
	queueInfo, err := inspector.GetQueueInfo("default")
	s.Require().NoError(err)
	s.Equal(1, queueInfo.Pending)
	s.cleanupQueue()

	setting = repositories.Settings().FindOrDefault(user, "notifications", "client-application-email-notifications", "userinfo-introspection")
	setting.Value = "true"
	err = repositories.Settings().Create(&setting)
	s.Require().NoError(err)
	notifier.Announce(user, "client.userinfo_introspection", utils.H{})
	queueInfo, err = inspector.GetQueueInfo("default")
	s.Require().NoError(err)
	s.Equal(1, queueInfo.Pending)
	s.cleanupQueue()

	setting = repositories.Settings().FindOrDefault(user, "notifications", "system-email-notifications", "client-authorization")
	setting.Value = "true"
	err = repositories.Settings().Create(&setting)
	s.Require().NoError(err)
	notifier.Announce(user, "user.authorization_granted", utils.H{})
	queueInfo, err = inspector.GetQueueInfo("default")
	s.Require().NoError(err)
	s.Equal(1, queueInfo.Pending)
	s.cleanupQueue()

	setting = repositories.Settings().FindOrDefault(user, "notifications", "system-email-notifications", "authentication")
	setting.Value = "false"
	err = repositories.Settings().Create(&setting)
	s.Require().NoError(err)
	notifier.Announce(user, "session.created", utils.H{})
	queueInfo, err = inspector.GetQueueInfo("default")
	s.Require().NoError(err)
	s.Equal(0, queueInfo.Pending)
	s.cleanupQueue()
}
