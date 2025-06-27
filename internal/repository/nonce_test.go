package repository

import (
	"github.com/earaujoassis/space/internal/models"
)

func (s *RepositoryTestSuite) TestNonceRepository__Create() {
	repository := NewNonceRepository(s.Memory)

	nonce := models.Nonce{
		ClientKey: models.GenerateRandomString(32),
		Code:      "",
		Nonce:     models.GenerateRandomString(128),
	}
	ok := repository.Create(nonce)
	s.Require().True(ok)

	nonce = models.Nonce{
		ClientKey: models.GenerateRandomString(32),
		Code:      "",
		Nonce:     models.GenerateRandomString(7),
	}
	ok = repository.Create(nonce)
	s.Require().False(ok)

	nonce = models.Nonce{
		ClientKey: models.GenerateRandomString(32),
		Code:      "",
		Nonce:     models.GenerateRandomString(129),
	}
	ok = repository.Create(nonce)
	s.Require().False(ok)
}

func (s *RepositoryTestSuite) TestNonceRepository__RetrieveByCode() {
	repository := NewNonceRepository(s.Memory)

	code := models.GenerateRandomString(32)
	nonceSrt := models.GenerateRandomString(128)
	nonce := models.Nonce{
		ClientKey: models.GenerateRandomString(32),
		Code:      code,
		Nonce:     nonceSrt,
	}
	ok := repository.Create(nonce)
	s.Require().True(ok)

	retrieved := repository.RetrieveByCode(code)
	s.Require().NotEmpty(retrieved.Code)
	s.Require().NotEmpty(retrieved.Nonce)
	s.Equal(nonceSrt, retrieved.Nonce)
	s.Equal(code, retrieved.Code)
}
