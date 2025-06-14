package handlers

import (
	"context"
	"my-api/internal/userService"
	"my-api/internal/web/users"
)

type UserHandler struct {
	service userService.UserService
}

func NewUserHandler(s userService.UserService) *UserHandler {
	return &UserHandler{service: s}
}

func (h *UserHandler) GetUsers(_ context.Context, _ users.GetUsersRequestObject) (users.GetUsersResponseObject, error) {
	allUsers, err := h.service.GetAllUsers()
	if err != nil {
		return nil, err
	}

	response := users.GetUsers200JSONResponse{}

	for _, usr := range allUsers {
		user := users.User{
			Id:       &usr.ID,
			Email:    &usr.Email,
			Password: &usr.Password,
		}
		response = append(response, user)
	}

	return response, nil
}

func (h *UserHandler) PostUser(_ context.Context, request users.PostUserRequestObject) (users.PostUserResponseObject, error) {
	body := request.Body

	userToCreate := userService.User{
		Email:    body.Email,
		Password: body.Password,
	}

	createdUser, err := h.service.CreateUser(userToCreate.Email, userToCreate.Password)
	if err != nil {
		return nil, err
	}

	response := users.PostUser201JSONResponse{
		Id:       &createdUser.ID,
		Email:    &createdUser.Email,
		Password: &createdUser.Password,
	}

	return response, nil
}

func (h *UserHandler) PatchUserByID(_ context.Context, request users.PatchUserByIDRequestObject) (users.PatchUserByIDResponseObject, error) {
	id := request.Id
	body := request.Body

	updatedUser, err := h.service.UpdateUser(id, body.Email, body.Password)
	if err != nil {
		return nil, err
	}

	response := users.User{
		Id:       &updatedUser.ID,
		Email:    &updatedUser.Email,
		Password: &updatedUser.Password,
	}

	return users.PatchUserByID200JSONResponse(response), nil
}

func (h *UserHandler) DeleteUserByID(_ context.Context, request users.DeleteUserByIDRequestObject) (users.DeleteUserByIDResponseObject, error) {
	id := request.Id

	err := h.service.DeleteUser(id)
	if err != nil {
		return nil, err
	}

	return users.DeleteUserByID204Response{}, nil
}
