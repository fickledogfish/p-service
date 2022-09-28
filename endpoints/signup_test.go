package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/suite"
)

func Test(t *testing.T) {
	suite.Run(t, new(signUpHandlerSuite))
}

type signUpHandlerSuite struct {
	suite.Suite

	request          *http.Request
	responseRecorder *httptest.ResponseRecorder

	sut signUpHandler
}

func (s *signUpHandlerSuite) SetupTest() {
	s.request = httptest.NewRequest("POST", "/", nil)

	s.responseRecorder = httptest.NewRecorder()

	s.sut = signUpHandler{}
}

// Test cases -----------------------------------------------------------------

func (s *signUpHandlerSuite) TestSignUpShouldOnlyAcceptPost() {
	for _, method := range []string{
		"GET",
		"PUT",
		"DELETE",
		"UPDATE",
	} {
		// Arrange
		s.request.Method = method

		// Act
		s.serveHTTP()

		// Assert
		s.Assert().Equal(http.StatusForbidden, s.responseRecorder.Code)
	}
}

// Halpers --------------------------------------------------------------------

func (sut *signUpHandlerSuite) serveHTTP() {
	sut.sut.ServeHTTP(sut.responseRecorder, sut.request)
}
