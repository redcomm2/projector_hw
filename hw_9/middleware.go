package main

import (
	"fmt"
	"net/http"
	"strings"
	"time"
)

func authorize(role string) func(http.HandlerFunc) http.HandlerFunc {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			sessionToken, err := r.Cookie("session_token")
			if err != nil || sessionToken.Value == "" {
				http.Redirect(w, r, "/login", http.StatusSeeOther)
				return
			}

			tokenParts := strings.Split(sessionToken.Value, ":")
			if len(tokenParts) != 3 || tokenParts[0] != role {
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}

			next(w, r)
		}
	}
}

func setSession(w http.ResponseWriter, userID int, classIndex int) {
	token := fmt.Sprintf("%s:%d:%d", "teacher", userID, classIndex)
	if userID != classes[classIndex].Teacher.ID {
		token = fmt.Sprintf("%s:%d:%d", "student", userID, classIndex)
	}
	http.SetCookie(w, &http.Cookie{
		Name:    "session_token",
		Value:   token,
		Path:    "/",
		Expires: time.Now().Add(1 * time.Minute),
	})
}
