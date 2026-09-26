package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/image-generator/internal/models"
	"github.com/image-generator/utils"
)

type tokenResponse struct {
	RefreshToken string `json:"refresh"`
	AccessToken  string `json:"access"`
}

// anything of type this must satisfy both the function interfaces.
type UserStore interface {
	LoginUser(credentials models.Login) (int, error)
}

type TokenService interface {
	GenerateTokens(userId int64, email, requirement string) (string, string, error)
}

type LoginHandler struct {
	Store UserStore
	Token TokenService
}

func (l *LoginHandler) HandleLogin(w http.ResponseWriter, r *http.Request) {
	var ld models.Login
	json.NewDecoder(r.Body).Decode(&ld)
	fmt.Println("[LOGIN HANDLER]: Login Data", ld)

	// now validate all of the given data and add to the database
	err := utils.ValidateLoginCredentials(ld)

	if err != nil {
		response := models.Response{
			Success: false,
			Message: "Login credentials can't be validated.",
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(response)
		return
	}

	// if the user doesn't exists
	id, err := l.Store.LoginUser(ld)
	if id == -1 {
		response := models.Response{
			Success: false,
			Message: "User doesnot exists",
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(404)
		json.NewEncoder(w).Encode(response)
		return
	}

	// provide the user the access-token and refresh-token
	refreshToken, accessToken, err := l.Token.GenerateTokens(int64(id), ld.Email, "both")
	fmt.Println("[LOGIN HANDLER]: Access token and refresh token: ", refreshToken, accessToken)
	if err != nil {
		fmt.Println("[LOGIN HANDLER] ", err)
		response := models.Response{
			Success: false,
			Message: "Server error.",
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(response)
		return
	}

	response := models.Response{
		Success: true,
		Message: "Successfully logged in.",
		Data: tokenResponse{
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
		},
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(response)

}
