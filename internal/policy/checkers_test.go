package policy

import (
	"github.com/brianvoe/gofakeit/v7"
)

func (s *PolicyTestSuite) TestSignInAttempts() {
	var result string

	ip := gofakeit.IPv4Address()
	result = s.rls.SignInAttemptStatus(ip)
	s.Equal(Clear, result)
	s.rls.RegisterSignInAttempt(ip)
	result = s.rls.SignInAttemptStatus(ip)
	s.Equal(Clear, result)
	result = s.rls.SignInAttemptStatus(ip)
	s.Equal(Clear, result)
	s.rls.RegisterSignInAttempt(ip)
	result = s.rls.SignInAttemptStatus(ip)
	s.Equal(Clear, result)
	s.rls.RegisterSignInAttempt(ip)
	result = s.rls.SignInAttemptStatus(ip)
	s.Equal(Preblocked, result)
	s.rls.RegisterSignInAttempt(ip)
	result = s.rls.SignInAttemptStatus(ip)
	s.Equal(Preblocked, result)
	s.rls.RegisterSignInAttempt(ip)
	result = s.rls.SignInAttemptStatus(ip)
	s.Equal(Blocked, result)

	s.rls.RegisterSuccessfulSignIn(ip)
	result = s.rls.SignInAttemptStatus(ip)
	s.Equal(Clear, result)
}

func (s *PolicyTestSuite) TestSignUpAttempts() {
	var result string

	ip := gofakeit.IPv4Address()
	result = s.rls.SignUpAttemptStatus(ip)
	s.Equal(Clear, result)
	s.rls.RegisterSignUpAttempt(ip)
	result = s.rls.SignUpAttemptStatus(ip)
	s.Equal(Clear, result)
	result = s.rls.SignUpAttemptStatus(ip)
	s.Equal(Clear, result)
	s.rls.RegisterSignUpAttempt(ip)
	result = s.rls.SignUpAttemptStatus(ip)
	s.Equal(Clear, result)
	s.rls.RegisterSignUpAttempt(ip)
	result = s.rls.SignUpAttemptStatus(ip)
	s.Equal(Preblocked, result)
	s.rls.RegisterSignUpAttempt(ip)
	result = s.rls.SignUpAttemptStatus(ip)
	s.Equal(Preblocked, result)
	s.rls.RegisterSignUpAttempt(ip)
	result = s.rls.SignUpAttemptStatus(ip)
	s.Equal(Blocked, result)

	s.rls.RegisterSuccessfulSignUp(ip)
	result = s.rls.SignUpAttemptStatus(ip)
	s.Equal(Clear, result)
}
