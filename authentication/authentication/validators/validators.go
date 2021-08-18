package validators

import (
	"authentication/pb"
	"errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"strings"
)

var (
	ErrEmptyName              = errors.New("name can't be empty")
	ErrEmptyEmail             = errors.New("email can't be empty")
	ErrEmptyPassword          = errors.New("password can't be empty")
	ErrEmailAlreadyRegistered = errors.New("email already registered")
	ErrInvalidUserId          = errors.New("invalid id")
	ErrSignInFailed           = errors.New("sign in failed")
)

func ValidateSignUp(user *pb.User) error {
	// TODO: check white space

	if !primitive.IsValidObjectID(user.Id) {
		return ErrInvalidUserId
	}

	if user.Email == "" {
		return ErrEmptyEmail
	}

	if user.Password == "" {
		return ErrEmptyPassword
	}

	if user.Name == "" {
		return ErrEmptyName
	}

	return nil
}

func FormatEmail(email string) string {
	return strings.TrimSpace(strings.ToLower(email))
}
