package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"example.com/p-service/models"
	"example.com/p-service/models/request"
	"github.com/stretchr/testify/suite"
)

func TestSignUpHandlerSuite(t *testing.T) {
	suite.Run(t, new(signUpHandlerSuite))
}

type signUpHandlerSuite struct {
	suite.Suite

	model request.SignUp

	request          *http.Request
	responseRecorder *httptest.ResponseRecorder

	sut signUpHandler
}

func (s *signUpHandlerSuite) SetupTest() {
	s.model = request.SignUp{
		Username: "some_name",
	}

	data, err := json.Marshal(s.model)
	s.Require().NoError(err)

	s.request = httptest.NewRequest("POST", "/", bytes.NewReader(data))
	s.responseRecorder = httptest.NewRecorder()

	s.sut = signUpHandler{}
}

// Test cases -----------------------------------------------------------------

func (s *signUpHandlerSuite) TestSignUpShouldOnlyAcceptPost() {
	expectedError, err := json.Marshal(
		models.NewErrorResponse("Forbidden method"),
	)
	s.Require().NoError(err)

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

		body, err := ioutil.ReadAll(s.responseRecorder.Body)
		s.Require().NoError(err)

		// Assert
		s.Equal(http.StatusForbidden, s.responseRecorder.Code)
		s.Equal(string(expectedError), string(body))
	}
}

func (s *signUpHandlerSuite) TestSignUpShouldAcceptPost() {
	// Arrange
	s.request.Method = "POST"

	// Act
	s.serveHTTP()

	_, err := ioutil.ReadAll(s.responseRecorder.Body)
	s.Require().NoError(err)

	// Assert
	s.Equal(http.StatusOK, s.responseRecorder.Code)
}

// Halpers --------------------------------------------------------------------

func (s *signUpHandlerSuite) serveHTTP() {
	s.sut.ServeHTTP(s.responseRecorder, s.request)
}
