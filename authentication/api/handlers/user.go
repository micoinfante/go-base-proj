package handlers

import (
	"authentication/api/restutil"
	"authentication/pb"
	"encoding/json"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

type AuthHandlers interface {
	SignUp(w http.ResponseWriter, r *http.Request)
	PutUser(w http.ResponseWriter, r *http.Request)
	GetUser(w http.ResponseWriter, r *http.Request)
	GetUsers(w http.ResponseWriter, r *http.Request)
	DeleteUser(w http.ResponseWriter, r *http.Request)
}

type authHandler struct {
	authSvcClient pb.AuthServiceClient
}

func NewAuthServiceClient(authSvcClient pb.AuthServiceClient) AuthHandlers {
	return &authHandler{authSvcClient: authSvcClient}
}

func (h *authHandler) SignUp(w http.ResponseWriter, r *http.Request) {
	if r.Body == nil {
		restutil.WriteError(w, http.StatusBadRequest, restutil.ErrEmptyBody)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(r.Body)
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		restutil.WriteError(w, http.StatusBadRequest, err)
		return
	}
	user := new(pb.User)
	err = json.Unmarshal(body, user)
	if err != nil {
		restutil.WriteError(w, http.StatusBadRequest, err)
		return
	}
	user.Created = time.Now().Unix()
	user.Updated = user.Created
	user.Id = primitive.NewObjectID().Hex()
	response, err := h.authSvcClient.SignUp(r.Context(), user)
	if err != nil {
		restutil.WriteError(w, http.StatusUnprocessableEntity, err)
		return
	}
	restutil.WriteAsJson(w, http.StatusCreated, response)
}

func (h *authHandler) PutUser(w http.ResponseWriter, r *http.Request) {
	if r.Body == nil {
		restutil.WriteError(w, http.StatusBadRequest, restutil.ErrEmptyBody)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(r.Body)
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		restutil.WriteError(w, http.StatusBadRequest, err)
		return
	}
	user := new(pb.User)
	err = json.Unmarshal(body, user)
	if err != nil {
		restutil.WriteError(w, http.StatusBadRequest, err)
		return
	}
	vars := mux.Vars(r)
	user.Id = vars["id"]

	response, err := h.authSvcClient.UpdateUser(r.Context(), user)
	if err != nil {
		restutil.WriteError(w, http.StatusUnprocessableEntity, err)
		return
	}
	restutil.WriteAsJson(w, http.StatusOK, response)
}

func (h *authHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	response, err := h.authSvcClient.GetUser(r.Context(), &pb.GetUserRequest{Id: vars["id"]})
	if err != nil {
		restutil.WriteError(w, http.StatusBadGateway, err)
		return
	}
	restutil.WriteAsJson(w, http.StatusOK, response)
}

func (h *authHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
	stream, err := h.authSvcClient.ListUsers(r.Context(), &pb.ListUsersRequest{})
	if err != nil {
		restutil.WriteError(w, http.StatusUnprocessableEntity, err)
		return
	}
	var users []*pb.User

	for {
		user, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			restutil.WriteError(w, http.StatusBadGateway, err)
			return
		}

		users = append(users, user)
	}
	restutil.WriteAsJson(w, http.StatusOK, users)
}

func (h *authHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	response, err := h.authSvcClient.DeleteUser(r.Context(), &pb.GetUserRequest{Id: vars["id"]})
	if err != nil {
		restutil.WriteError(w, http.StatusBadGateway, err)
		return
	}
	w.Header().Set("Entity", response.Id)
	restutil.WriteAsJson(w, http.StatusNoContent, nil)
}
