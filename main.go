package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	app "oracle.com/self/partner-test-env/App"
	config "oracle.com/self/partner-test-env/Config"
	"oracle.com/self/partner-test-env/database"
)

func main() {
	logger := log.New(os.Stdout, "", log.LstdFlags|log.Lshortfile)

	// Initialize database connection
	database.InitializeDB()
	logger.Println("Database initialized successfully.")

	conf := config.NewConfig()
	app := app.NewApp(conf, logger)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	log.Println("Setting up static file handler for /images/")
	http.Handle("/images/", http.StripPrefix("/images/", http.FileServer(http.Dir("images"))))
	app.Run()

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGTERM, syscall.SIGKILL, os.Interrupt, os.Kill)
	sig := <-sigs
	logger.Printf("Got shutdown interupt signal %v\n", sig)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	app.Shutdown(ctx)

}
