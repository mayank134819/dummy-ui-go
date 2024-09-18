package app

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	config "oracle.com/self/partner-dummy-env/Config"
	"oracle.com/self/partner-dummy-env/controller"
)

type App struct {
	logger *log.Logger
	server *http.Server
	router *mux.Router
	config *config.Config
}

func NewApp(config *config.Config, logger *log.Logger) *App {
	app := &App{
		logger: logger,
		config: config,
	}
	app.intialization()
	app.setRouter()
	return app
}

func (app *App) intialization() {
	app.router = mux.NewRouter()
	app.logger.Println("Server configer on port: ", app.config.ServerAddr)
	app.server = &http.Server{
		Addr:         app.config.ServerAddr,
		Handler:      app.router,
		ErrorLog:     app.logger,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}
}

func (app *App) setRouter() {
	homepage := controller.NewHomePage(app.logger)
	signup := controller.NewSingUp(app.logger)

	app.router.HandleFunc("/", homepage.Home)
	app.router.HandleFunc("/signup", signup.SignUp)
}

func (app *App) Run() {
	app.logger.Println("Running server on: ", app.config.ServerAddr)
	app.logger.Fatalln(app.server.ListenAndServe())
}

func (app *App) Shutdown(ctx context.Context) {
	app.logger.Panic("Shuting down server")
	app.server.Shutdown(ctx)
}
