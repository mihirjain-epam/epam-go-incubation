package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"epam.com/web-services/library-management/login-service/authentication"
	"epam.com/web-services/library-management/login-service/authorization"
	"epam.com/web-services/library-management/login-service/models"
	"epam.com/web-services/library-management/login-service/repository"
)

func handleLogin(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var userCredentialFromRequest models.UserCredential
		err := json.NewDecoder(r.Body).Decode(&userCredentialFromRequest)
		if err != nil {
			log.Print(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if !authentication.IsUserAuthenticated(userCredentialFromRequest) {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		userCredentialFromRepo, err := repository.GetCredential(userCredentialFromRequest.UserName)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		validTokenPair, err := authorization.GetToken(int64(userCredentialFromRepo.Id))
		if err != nil {
			log.Print(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Add("JWT-Token", validTokenPair["access_token"].Token)
		cookie := http.Cookie{Name: "refresh_token", Value: validTokenPair["refresh_token"].Token, Expires: time.Unix(validTokenPair["refresh_token"].Exp, 0), HttpOnly: true}
		http.SetCookie(w, &cookie)
		w.WriteHeader(http.StatusOK)
		return
	}
	w.WriteHeader(http.StatusMethodNotAllowed)
}

func handleLogout(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var userCredential models.UserCredential
		err := json.NewDecoder(r.Body).Decode(&userCredential)
		if err != nil {
			log.Print(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		authorization.LogoutToken(int64(userCredential.Id))
		w.WriteHeader(http.StatusOK)
		return
	}
	w.WriteHeader(http.StatusMethodNotAllowed)
}

func handleRefreshToken(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		validAccessToken, err := authorization.RefreshAccessToken(w, r)
		if err != nil {
			log.Print(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		w.Header().Add("JWT-Token", validAccessToken)
		w.WriteHeader(http.StatusOK)
		return
	}
	w.WriteHeader(http.StatusMethodNotAllowed)
}

func SetupRoutes() {
	loginHandler := http.HandlerFunc(handleLogin)
	http.Handle("/login", loginHandler)
	logoutHandler := http.HandlerFunc(handleLogout)
	http.Handle("/logout", logoutHandler)
	refreshHandler := http.HandlerFunc(handleRefreshToken)
	http.Handle("/refresh", refreshHandler)
}
