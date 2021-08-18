package services

import (
	"authentication/authentication/models"
	"authentication/authentication/repository"
	"authentication/authentication/validators"
	"authentication/pb"
	"authentication/security"
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"strings"
	"time"
)

type authService struct {
	usersRepository repository.UsersRepository
}

func NewAuthService(usersRepository repository.UsersRepository) pb.AuthServiceServer {
	return &authService{usersRepository: usersRepository}
}

func (a authService) SignUp(ctx context.Context, user *pb.User) (*pb.User, error) {
	err := validators.ValidateSignUp(user)
	if err != nil {
		return nil, err
	}
	user.Password, err = security.EncryptPassword(user.Password)
	if err != nil {
		return nil, err
	}
	user.Name = strings.TrimSpace(user.Name)
	user.Email = validators.FormatEmail(user.Email)
	found, err := a.usersRepository.GetByEmail(user.Email)
	if err == mongo.ErrNoDocuments {
		_user := new(models.User)
		_user.FromProtoBuffer(user)
		err := a.usersRepository.Save(_user)
		if err != nil {
			return nil, err
		}
		return _user.ToProtoBuffer(), nil
	}

	if found == nil {
		return nil, err
	}

	return nil, validators.ErrEmailAlreadyRegistered
}

func (a authService) GetUser(ctx context.Context, request *pb.GetUserRequest) (*pb.User, error) {
	if !primitive.IsValidObjectID(request.Id) {
		return nil, validators.ErrInvalidUserId
	}
	found, err := a.usersRepository.GetById(request.Id)
	if err != nil {
		return nil, err
	}

	return found.ToProtoBuffer(), nil
}

func (a authService) ListUsers(request *pb.ListUsersRequest, stream pb.AuthService_ListUsersServer) error {
	users, err := a.usersRepository.GetAll()
	if err != nil {
		return err
	}

	for _, user := range users {
		err := stream.Send(user.ToProtoBuffer())
		if err != nil {
			return err
		}
	}
	return nil
}

func (a authService) UpdateUser(ctx context.Context, user *pb.User) (*pb.User, error) {
	if !primitive.IsValidObjectID(user.Id) {
		return nil, validators.ErrInvalidUserId
	}

	_user, err := a.usersRepository.GetById(user.Id)
	if err != nil {
		return nil, err
	}
	_user.Name = strings.TrimSpace(user.Name)
	if _user.Name == "" {
		return nil, validators.ErrEmptyName
	}
	if user.Name == _user.Name {
		return _user.ToProtoBuffer(), nil
	}

	_user.Name = user.Name
	_user.Updated = time.Now()
	err = a.usersRepository.Update(_user)
	return _user.ToProtoBuffer(), err
}

func (a authService) DeleteUser(ctx context.Context, request *pb.GetUserRequest) (*pb.DeleteUserResponse, error) {
	if !primitive.IsValidObjectID(request.Id) {
		return nil, validators.ErrInvalidUserId
	}
	err := a.usersRepository.Delete(request.Id)

	if err != nil {
		return nil, err
	}

	return &pb.DeleteUserResponse{Id: request.Id}, nil
}

func (a authService) SignIn(ctx context.Context, request *pb.SignInRequest) (*pb.SignInResponse, error) {
	request.Email = validators.FormatEmail(request.Email)
	user, err := a.usersRepository.GetByEmail(request.Email)

	if err != nil {
		return nil, validators.ErrSignInFailed
	}
	err = security.VerifyPassword(user.Password, request.Password)
	if err != nil {
		return nil, validators.ErrSignInFailed
	}

	token, err := security.NewToken(user.Id.Hex())
	if err != nil {
		return nil, validators.ErrSignInFailed
	}
	return &pb.SignInResponse{
		User:  user.ToProtoBuffer(),
		Token: token,
	}, nil
}
