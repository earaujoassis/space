package integration

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/earaujoassis/space/test/helpers"
)

type OAuthProviderSuite struct {
	helpers.OAuthTestSuite
}

func TestOAuthProviderSuite(t *testing.T) {
	suite.Run(t, new(OAuthProviderSuite))
}
