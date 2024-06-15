package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"server-techno-flow/internal/entities"
	"strconv"
)

func (h *Handler) CreateUser(c *gin.Context) {
	const op = "user/Handler.CreateUser"
	var userDto entities.UserCreateDto

	if err := c.BindJSON(&userDto); err != nil {
		ResponseError(c, "failed to bind user dto", err.Error(), http.StatusInternalServerError)
		return
	}

	id, err := h.services.User.CreateUser(c, userDto)

	if err != nil {
		ResponseError(c, "failed to create user", err.Error(), http.StatusInternalServerError)
		return
	}

	ResponseSuccess(c, "user created successfully", map[string]interface{}{"user_id": id})
}

func (h *Handler) GetAllUsers(c *gin.Context) {
	const op = "user/Handler.GetAllUsers"
	users, err := h.services.User.GetAllUsers(c)

	if err != nil {
		ResponseError(c, "failed to fetch users", err.Error(), http.StatusInternalServerError)
		return
	}

	ResponseSuccess(c, "users fetched successfully", users)
}

func (h *Handler) GetUserById(c *gin.Context) {
	const op = "user/Handler.GetUserById"
	paramId := c.Param("id")

	id, err := strconv.Atoi(paramId)
	if err != nil {
		ResponseError(c, "invalid query id param", err.Error(), http.StatusBadRequest)
		return
	}

	user, err := h.services.User.GetUserById(c, id)

	if err != nil {
		ResponseError(c, "failed to fetch user", err.Error(), http.StatusInternalServerError)
		return
	}

	ResponseSuccess(c, "user fetched successfully", user)
}

func (h *Handler) DeleteUser(c *gin.Context) {
	const op = "user/Handler.DeleteUser"
	paramId := c.Param("id")

	id, err := strconv.Atoi(paramId)
	if err != nil {
		ResponseError(c, "invalid query id param", err.Error(), http.StatusBadRequest)
		return
	}

	_, err = h.services.User.DeleteUser(c, id)
	if err != nil {
		ResponseError(c, "failed to delete user", err.Error(), http.StatusInternalServerError)
		return
	}

	ResponseSuccess(c, "user deleted successfully", map[string]interface{}{"user_id": id})
}

func (h *Handler) UpdateUser(c *gin.Context) {
	const op = "user/Handler.UpdateUser"
	paramId := c.Param("id")

	id, err := strconv.Atoi(paramId)
	if err != nil {
		ResponseError(c, "invalid query id param", err.Error(), http.StatusBadRequest)
		return
	}

	var userUpdateDto entities.UserUpdateDto

	if err := c.BindJSON(&userUpdateDto); err != nil {
		ResponseError(c, "failed to bind user dto", err.Error(), http.StatusInternalServerError)
		return
	}

	err = h.services.User.UpdateUser(c, id, userUpdateDto)

	if err != nil {
		ResponseError(c, "failed to update user", err.Error(), http.StatusInternalServerError)
		return
	}

	ResponseSuccess(c, "user updated successfully", map[string]interface{}{"user_id": id})
}
