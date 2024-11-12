package controller

import (
    "log"
    "net/http"
    "oracle.com/self/partner-test-env/model"
)

type SignUpZoom struct {
    logger *log.Logger
}

func NewSignUpZoom(logger *log.Logger) *SignUpZoom {
    return &SignUpZoom{
        logger: logger,
    }
}

func (su *SignUpZoom) SignUpZoom(w http.ResponseWriter, r *http.Request) {
    if r.Method == http.MethodPost {
        r.ParseForm()
        name := r.FormValue("name")
        username := r.FormValue("username")
        password := r.FormValue("password")

        // Save to in-memory DB
        model.Users[username] = model.User{Name: name, Username: username, Password: password}

        http.Redirect(w, r, "/signinzoom", http.StatusSeeOther)
    } else {
        http.ServeFile(w, r, "templates/sign_up_zoom.html")
    }
}
