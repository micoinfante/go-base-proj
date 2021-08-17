package routes

import (
	"authentication/api/handlers"
	"net/http"
)

func NewAuthRoute(authHandler handlers.AuthHandlers) []*Route {
	return []*Route{
		{
			Path:    "/signup",
			Method:  http.MethodPost,
			Handler: authHandler.SignUp,
		},
		{
			Path:    "/signin",
			Method:  http.MethodPost,
			Handler: authHandler.SignIn,
		},
		{
			Path:         "/users",
			Method:       http.MethodGet,
			Handler:      authHandler.GetUsers,
			AuthRequired: true,
		},
		{
			Path:         "/users/{id}",
			Method:       http.MethodGet,
			Handler:      authHandler.GetUser,
			AuthRequired: true,
		},
		{
			Path:         "/users/{id}",
			Method:       http.MethodPut,
			Handler:      authHandler.PutUser,
			AuthRequired: true,
		},
		{
			Path:         "/users/{id}",
			Method:       http.MethodDelete,
			Handler:      authHandler.DeleteUser,
			AuthRequired: true,
		},
	}
}
