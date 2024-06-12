package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"server-techno-flow/internal/entities"
	"time"
)

type SignInResponse struct {
	User        entities.User `json:"user"`
	AccessToken string        `json:"access_token"`
}

func (h *Handler) signIn(c *gin.Context) {
	const op = "auth/Handler.signIn"

	var signInDto entities.UserSignInDto

	if err := c.BindJSON(&signInDto); err != nil {
		ResponseError(c, "failed to bind sign-in dto", err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Println(signInDto)

	user, refreshToken, accessToken, err := h.services.Auth.SignIn(signInDto)
	if err != nil {
		ResponseError(c, "failed to sign-in", err.Error(), http.StatusInternalServerError)
		return
	}

	http.SetCookie(c.Writer, &http.Cookie{
		Name:     "refresh_token",
		Value:    refreshToken,
		Expires:  time.Now().Add(48 * time.Hour),
		HttpOnly: true,
		Secure:   false,
	})

	//c.SetCookie("refresh_token", refreshToken, 3600, "/", "", false, true)

	signInResponse := SignInResponse{
		User:        user,
		AccessToken: accessToken,
	}

	ResponseSuccess(c, "user signed in successfully", signInResponse)
}

func (h *Handler) signOut(c *gin.Context) {
	const op = "auth/Handler.signOut"

	refreshToken, err := c.Cookie("refresh_token")
	if err != nil {
		ResponseError(c, "refresh token not provided", err.Error(), http.StatusBadRequest)
		return
	}

	if err = h.services.Auth.SignOut(refreshToken); err != nil {
		ResponseError(c, "error during Logout", err.Error(), http.StatusInternalServerError)
		return
	}

	http.SetCookie(c.Writer, &http.Cookie{
		Name:     "refresh_token",
		Value:    "",
		Expires:  time.Now().AddDate(0, 0, -1),
		HttpOnly: true,
		Secure:   false,
	})

	ResponseSuccess(c, "successfully signed out", nil)
}

func (h *Handler) refresh(c *gin.Context) {
	const op = "auth/Handler.refresh"
	token, err := c.Cookie("refresh_token")
	if err != nil {
		ResponseError(c, "refresh token not provided", err.Error(), http.StatusUnauthorized)
		return
	}

	user, refreshToken, accessToken, err := h.services.Auth.Refresh(token)
	if err != nil {
		ResponseError(c, "failed to sign-in", err.Error(), http.StatusUnauthorized)
		return
	}

	http.SetCookie(c.Writer, &http.Cookie{
		Name:     "refresh_token",
		Value:    refreshToken,
		Expires:  time.Now().Add(48 * time.Hour),
		HttpOnly: true,
		Secure:   false,
	})

	refreshResponse := SignInResponse{
		User:        user,
		AccessToken: accessToken,
	}

	ResponseSuccess(c, "access token refreshed successfully", refreshResponse)
}
