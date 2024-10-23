package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	app "oracle.com/self/partner-test-env/App"
	config "oracle.com/self/partner-test-env/Config"
)

func main() {
	logger := log.New(os.Stdout, "", log.LstdFlags|log.Lshortfile)
	conf := config.NewConfig()
	app := app.NewApp(conf, logger)
	app.Run()

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGTERM, syscall.SIGKILL, os.Interrupt, os.Kill)
	sig := <-sigs
	logger.Printf("Got shutdown interupt signal %v\n", sig)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	app.Shutdown(ctx)

}
