package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"server-techno-flow/internal/domain"
	"strconv"
)

func (h *Handler) CreateUser(c *gin.Context) {
	var userDto domain.UserCreateDto

	if err := c.BindJSON(&userDto); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	id, err := h.services.User.CreateUser(userDto)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) GetAllUsers(c *gin.Context) {
	users, err := h.services.User.GetAllUsers()

	if err != nil {
		fmt.Printf("Error Fetching Users: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"users": users,
	})
}

func (h *Handler) GetUserById(c *gin.Context) {
	paramId := c.Param("id")
	fmt.Printf("GetUserById Called, param: %s\n", paramId)

	id, err := strconv.Atoi(paramId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.services.User.GetUserById(id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"user": user,
	})
}

func (h *Handler) DeleteUser(c *gin.Context) {
	fmt.Printf("DeleteUser Called, param: %s\n", c.Param("id"))
	paramId := c.Param("id")

	id, err := strconv.Atoi(paramId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err = h.services.User.DeleteUser(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"user_id": id,
	})
}

func (h *Handler) UpdateUser(c *gin.Context) {
	fmt.Printf("UpdateUser Called, param: %s\n", c.Param("id"))
	paramId := c.Param("id")

	id, err := strconv.Atoi(paramId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var userUpdateDto domain.UserUpdateDto

	if err := c.BindJSON(&userUpdateDto); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	errUpdate := h.services.User.UpdateUser(id, userUpdateDto)

	if errUpdate != nil {
		fmt.Printf("Error Updating User: %v\n", errUpdate)
		c.JSON(http.StatusInternalServerError, gin.H{"error": errUpdate.Error()})
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}
