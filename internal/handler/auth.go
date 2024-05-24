package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"server-techno-flow/internal/domain"
)

func (h *Handler) signIn(c *gin.Context) {
	fmt.Println("signIp called")

	var signInDto domain.UserSignInDto

	if err := c.BindJSON(&signInDto); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error sign in dto": err.Error()})
		return
	}

	fmt.Println(signInDto)

	user, err := h.services.Auth.SignIn(signInDto)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error while signing in": err.Error()})
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"signed in user": user,
	})
}

func (h *Handler) signOut(c *gin.Context) {
	fmt.Println("signOut called")

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": "signed out",
	})
}
