package main

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"example.com/p-service/models/request"
	"example.com/p-service/models/response"
	"github.com/stretchr/testify/suite"
)

func TestSignUpHandlerSuite(t *testing.T) {
	suite.Run(t, new(signUpHandlerSuite))
}

type signUpHandlerSuite struct {
	suite.Suite

	requestModel request.SignUp

	request          *http.Request
	responseRecorder *httptest.ResponseRecorder

	sut http.Handler
}

func (s *signUpHandlerSuite) SetupTest() {
	email := "master_of_the_universe@example.com"
	s.requestModel = request.SignUp{
		Username: "My Name",
		Password: "123",
		Email:    &email,
	}

	data, err := json.Marshal(s.requestModel)
	s.Require().NoError(err)

	s.request = httptest.NewRequest("POST", "/", bytes.NewReader(data))
	s.responseRecorder = httptest.NewRecorder()

	s.sut = signUp{}.handler()
}

// Test cases -----------------------------------------------------------------

func (s *signUpHandlerSuite) TestSignUpShouldOnlyAcceptPost() {
	expectedError, err := json.Marshal(
		response.NewError("Method not allowed"),
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
		s.T().Logf("Testing for method %s", method)
		s.Equal(http.StatusMethodNotAllowed, s.responseRecorder.Code)
		s.Equal(string(expectedError), string(body))
	}
}

func (s *signUpHandlerSuite) TestSignUpShouldAcceptPost() {
	// Act
	s.serveHTTP()

	_, err := ioutil.ReadAll(s.responseRecorder.Body)
	s.Require().NoError(err)

	// Assert
	s.Equal(http.StatusOK, s.responseRecorder.Code)
}

func (s *signUpHandlerSuite) TestServeHTTPShouldReturnTheCorrectModel() {
	// Act
	s.serveHTTP()

	bodyData, err := ioutil.ReadAll(s.responseRecorder.Body)
	s.Require().NoError(err)

	res := response.SignUp{}
	err = json.Unmarshal(bodyData, &res)
	s.Require().NoError(err)

	// Assert
	s.Equal(http.StatusOK, s.responseRecorder.Code)
	s.Equal(s.requestModel.Username, res.Username)
}

func (s *signUpHandlerSuite) TestServeHTTPShouldReturnInternalErrorOnInvalidJSON() {
	// Arrange
	s.request.Body = io.NopCloser(
		strings.NewReader(strings.Repeat(".", BODY_MAX_BYTES+1)),
	)

	// Act
	s.serveHTTP()

	// Assert
	s.Equal(http.StatusInternalServerError, s.responseRecorder.Code)
}

// Helpers --------------------------------------------------------------------

func (s *signUpHandlerSuite) serveHTTP() {
	s.sut.ServeHTTP(s.responseRecorder, s.request)
}
