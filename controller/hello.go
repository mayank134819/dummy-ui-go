package controller

import (
    "html/template"
    "net/http"
    "log"
)

type HelloWorld struct {
    logger *log.Logger
}

func NewHelloWorld(logger *log.Logger) *HelloWorld {
    return &HelloWorld{
        logger: logger,
    }
}

func (hw *HelloWorld) Hello(w http.ResponseWriter, r *http.Request) {
    tmpl, err := template.ParseFiles("templates/hello.html")
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "text/html")
    tmpl.Execute(w, nil)
}
