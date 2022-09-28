package request

import (
	"encoding"
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/suite"
)

func TestSignUpSuite(t *testing.T) {
	suite.Run(t, new(signUpSuite))
}

type signUpSuite struct {
	suite.Suite
}

func (s *signUpSuite) SetupTest() {
}

// Test cases -----------------------------------------------------------------

func (s signUpSuite) TestEnsureSignUpImplementsBinaryUnmarshaler() {
	s.Implements((*encoding.BinaryUnmarshaler)(nil), new(SignUp))
}

func (s signUpSuite) TestUnmarshalBinaryShouldAcceptAValidJSONString() {
	// Arrange
	email := "person.random@example.com"

	expectedSignUp := SignUp{
		Username: "Random Person",
		Password: "123",
		Email:    &email,
	}

	data, err := json.Marshal(expectedSignUp)
	s.Require().NoError(err)

	sut := SignUp{}

	// Act
	err = sut.UnmarshalBinary(data)

	// Assert
	s.NoError(err)
	s.Equal(expectedSignUp, sut)
}
