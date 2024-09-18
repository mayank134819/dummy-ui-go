package controller

import (
	"log"
	"net/http"
	"os"
	"path/filepath"
)

type HomePage struct {
	logger *log.Logger
}

func NewHomePage(logger *log.Logger) *HomePage {
	return &HomePage{
		logger: logger,
	}
}

func (hp *HomePage) Home(w http.ResponseWriter, r *http.Request) {
	path := filepath.Join("templates", "home.html")
	fi, err := os.Stat(path)
	if os.IsNotExist(err) || fi.IsDir() {
		hp.logger.Println("path not found", path)
		// return
	}
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }
	log.Println("HomePage")
	http.ServeFile(w, r, path)
}
