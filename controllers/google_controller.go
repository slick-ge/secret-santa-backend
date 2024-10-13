package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"golang.org/x/oauth2"
	"net/http"
	"secret-santa-backend/utils/utils"
)

var googleOauthConfig *oauth2.Config

// Google OAuth2 callback handler
func googleCallbackHandler(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")
	token, err := googleOauthConfig.Exchange(context.Background(), code)

	if err != nil {
		fmt.Println("Error exchanging code: " + err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	userInfo, err := utils.GetUserInfo(token.AccessToken)
	if err != nil {
		fmt.Println("Error getting user info: " + err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	signedToken, err := utils.SignJWT(userInfo)
	if err != nil {
		fmt.Println("Error signing token: " + err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"token": signedToken})
}
