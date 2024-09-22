package controller

import (
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gorilla/sessions"
)

type HomePage struct {
	logger       *log.Logger
	sessionStore *sessions.CookieStore
}

func NewHomePage(logger *log.Logger, sessionStore *sessions.CookieStore) *HomePage {
	return &HomePage{
		logger:       logger,
		sessionStore: sessionStore,
	}
}

func (hp *HomePage) Home(w http.ResponseWriter, r *http.Request) {
	path := filepath.Join("templates", "home.html")
	fi, err := os.Stat(path)
	if os.IsNotExist(err) || fi.IsDir() {
		hp.logger.Println("path not found", path)
		path = filepath.Join("/home/opc/templates", "home.html")
	}

	token := r.Header.Get("token")
	if token == "" {
		token = "test-token"
	}
	session, _ := hp.sessionStore.Get(r, "selfReg")
	session.Values["selfToken"] = token
	sessions.Save(r, w)

	log.Println("HomePage")
	if token != "" {
		w.Header().Set("token", token)
	}
	http.ServeFile(w, r, path)
}
