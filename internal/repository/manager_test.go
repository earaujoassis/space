package repository

func (s *RepositoryTestSuite) TestRepositoryManager() {
	manager := NewRepositoryManager(s.db, s.ms)
	s.Require().NotNil(manager.Actions())
	s.Require().NotNil(manager.Clients())
	s.Require().NotNil(manager.Languages())
	s.Require().NotNil(manager.Nonces())
	s.Require().NotNil(manager.Services())
	s.Require().NotNil(manager.Sessions())
	s.Require().NotNil(manager.Users())
	s.Require().NotNil(manager.Emails())
	s.Require().NotNil(manager.Settings())
}
