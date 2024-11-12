package controller

import (
    "log"
    "net/http"
    "oracle.com/self/partner-test-env/model"
)

type SignInZoom struct {
    logger *log.Logger
}

func NewSignInZoom(logger *log.Logger) *SignInZoom {
    return &SignInZoom{
        logger: logger,
    }
}

func (si *SignInZoom) SignInZoom(w http.ResponseWriter, r *http.Request) {
    if r.Method == http.MethodPost {
        r.ParseForm()
        username := r.FormValue("username")
        password := r.FormValue("password")

        // Check if the user exists and the password is correct
        user, exists := model.Users[username]
        if !exists || user.Password != password {
            http.Error(w, "Invalid username or password", http.StatusUnauthorized)
            return
        }

        http.Redirect(w, r, "/dashboard", http.StatusSeeOther) // Redirect to dashboard after successful login
    } else {
        http.ServeFile(w, r, "templates/sign_in_zoom.html")
    }
}
