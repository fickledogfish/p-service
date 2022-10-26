package middlewares

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/suite"
)

func TestAllowedMethodsSuite(t *testing.T) {
	suite.Run(t, new(allowedMethodsSuite))
}

type allowedMethodsSuite struct {
	suite.Suite

	request          *http.Request
	responseRecorder *httptest.ResponseRecorder

	allowedMethods *[]string
	nextHandler    http.Handler

	sut func() http.Handler
}

func (s *allowedMethodsSuite) SetupTest() {
	s.request = httptest.NewRequest(http.MethodGet, "/", strings.NewReader(""))
	s.responseRecorder = httptest.NewRecorder()

	s.nextHandler = http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	s.allowedMethods = &[]string{http.MethodGet}

	// Lazy eval the SUT to get the newest values.
	s.sut = func() http.Handler {
		return AllowedMethods(*s.allowedMethods, s.nextHandler)
	}
}

// Test cases -----------------------------------------------------------------

func (s *allowedMethodsSuite) TestAllowedMethodsShouldRejectWith403ForbiddenMethods() {
	// Arrange
	s.request = httptest.NewRequest(http.MethodGet, "/", strings.NewReader(""))
	s.allowedMethods = &[]string{http.MethodPost}

	// Act
	s.sut().ServeHTTP(s.responseRecorder, s.request)

	// Assert
	s.Equal(http.StatusMethodNotAllowed, s.responseRecorder.Code)
}

func (s *allowedMethodsSuite) TestAllowedMethodsShouldIgnoreNextHandlerOnFailure() {
	// Arrange
	nextHandlerCalled := false

	s.request = httptest.NewRequest(http.MethodGet, "/", strings.NewReader(""))
	s.allowedMethods = &[]string{http.MethodPost}

	s.nextHandler = http.HandlerFunc(func(http.ResponseWriter, *http.Request) {
		nextHandlerCalled = true
	})

	// Act
	s.sut().ServeHTTP(s.responseRecorder, s.request)
	s.Require().Equal(http.StatusMethodNotAllowed, s.responseRecorder.Code)

	// Assert
	s.False(nextHandlerCalled)
}

func (s *allowedMethodsSuite) TestAllowedMethodsShouldCallTheNextHandlerOnSuccess() {
	// Arrange
	nextHandlerCalled := false
	expectedStatus := http.StatusTeapot

	s.nextHandler = http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		w.WriteHeader(expectedStatus)
		nextHandlerCalled = true
	})

	// Act
	s.sut().ServeHTTP(s.responseRecorder, s.request)
	s.Require().Equal(expectedStatus, s.responseRecorder.Code)

	// Assert
	s.True(nextHandlerCalled)
}
