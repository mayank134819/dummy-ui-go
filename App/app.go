package app

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	config "oracle.com/self/partner-test-env/Config"
	"oracle.com/self/partner-test-env/controller"
)

type App struct {
	logger       *log.Logger
	server       *http.Server
	router       *mux.Router
	sessionStore *sessions.CookieStore
	config       *config.Config
}

func NewApp(config *config.Config, logger *log.Logger) *App {
	app := &App{
		logger:       logger,
		config:       config,
		sessionStore: sessions.NewCookieStore([]byte("test-secret-key")),
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
	homepage := controller.NewHomePage(app.logger, app.sessionStore)
	signup := controller.NewSignUp(app.logger, app.sessionStore)
	signin := controller.NewSignIn(app.logger, app.sessionStore)

	helloWorld := controller.NewHelloWorld(app.logger)
	sign_in_zoom := controller.NewSignInZoom(app.logger)
	sign_up_zoom := controller.NewSignUpZoom(app.logger)
	showSubscriptionDetails := controller.NewShowSubscriptionDetails(app.logger)

	app.router.HandleFunc("/", homepage.Home)
	app.router.HandleFunc("/signup", signup.SignUp)
	app.router.HandleFunc("/signin", signin.SignIn)

	app.router.HandleFunc("/hello", helloWorld.Hello)

	app.router.HandleFunc("/signupzoom",sign_up_zoom.SignUpZoom)
	app.router.HandleFunc("/signinzoom",sign_in_zoom.SignInZoom)

	app.router.HandleFunc("/showSubscriptionDetails", showSubscriptionDetails.Show)
	app.router.HandleFunc("/showSubscriptionDetails/{subscriptionToken}", showSubscriptionDetails.Show)

	app.router.HandleFunc("/activateSubscription", showSubscriptionDetails.Activate)

}

func (app *App) Run() {
	app.logger.Println("Running server on: ", app.config.ServerAddr)
	app.logger.Fatalln(app.server.ListenAndServe())
}

func (app *App) Shutdown(ctx context.Context) {
	app.logger.Panic("Shuting down server")
	app.server.Shutdown(ctx)
}
