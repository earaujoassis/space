package integration

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/earaujoassis/space/test/helpers"
)

type OIDCProviderSuite struct {
	helpers.OIDCTestSuite
}

func TestOIDCProviderSuite(t *testing.T) {
	suite.Run(t, new(OIDCProviderSuite))
}
