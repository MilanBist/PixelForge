package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/image-generator/internal/models"
	"github.com/image-generator/utils"
)


type Refresh struct{
	Token TokenService
}
type refreshToken struct{
	RefreshToken string		`json:"refreshToken"`
}

func(tkn Refresh) HandleNewAccessToken(w http.ResponseWriter, r *http.Request){
	// get a new access token based on certain string
	var token refreshToken
	json.NewDecoder(r.Body).Decode(&token)

	fmt.Println("Refresh Token is: ", token.RefreshToken)

	if token.RefreshToken == ""{
		response := models.Response{
			Success: false,
			Message: "can't get.",
			Data: "Attach refreshToken as refreshToken: tokenstring",
		}

		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(&response)
		return
	}

	// else first validate the token string
	userId, email, err := utils.ValidateToken(token.RefreshToken, "refresh")
	if err != nil{
		response := models.Response{
			Success: false,
			Message: "expired",
			Data: "Attach correct token as refreshToken: <token_string>",
		}
		w.WriteHeader(401)
		json.NewEncoder(w).Encode(&response)
		return
	}

	// respond with the new access token
	_, accessToken, err := tkn.Token.GenerateTokens(int64(userId), email, "access")
	if err != nil{
		response := models.Response{
			Success: false,
			Message: "Error in generating new access token.",
		}

		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(&response)
	}

	response := models.Response{
		Success: true,
		Message: "Access token generated",
		Data: map[string]string{
			"accessToken": accessToken,
		},
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(&response)

}