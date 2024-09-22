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

	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Origin", "*")                            // Allow all origins, or specify your domain
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")          // Allow specific methods
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization") // Allow specific headers

	log.Println("HomePage")
	http.ServeFile(w, r, path)
}
