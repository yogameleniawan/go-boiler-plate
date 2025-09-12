package main

import (
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"

	containerPkg "github.com/absendulu-project/backend/container"
	"github.com/absendulu-project/backend/pkg/config"
	"github.com/absendulu-project/backend/pkg/server"
)

var environtment string

func init() {
	env := flag.String("env", "development", "Environment (development/production)")
	flag.Parse()

	environtment = *env

	switch *env {
	case "development":
		config.LoadConfig("./config/config.development.yaml")
	case "staging":
		config.LoadConfig("./config/config.staging.yaml")
	case "production":
		config.LoadConfig("./config/config.production.yaml")
	}
}

func main() {
	// setup server
	container, err := containerPkg.New()
	if err != nil {
		log.Fatal(err)
	}

	err = container.Invoke(Start)
	if err != nil {
		log.Fatal(err)
	}

	// add quit signal
	quit := make(chan os.Signal, 1)
	container.Provide(func() chan os.Signal {
		return quit
	})

	signal.Notify(quit, os.Interrupt, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	err = container.Invoke(Shutdown)
	if err != nil {
		log.Fatal(err)
	}
}

func Start(
	svr server.Server,
) {
	log.Println("Starting server...")
	if err := svr.Start(); err != nil {
		log.Fatal(err)
	}

	if environtment == "production" {
		// for the security purpose, we need to remove file configuration after server start
		// This can be used in case the container where the backend is located is hacked,
		// information related to DB configuration, etc. cannot be accessed.
		log.Println("Deleting configuration after service running...")
		if err := os.Remove("./config/config.production.yaml"); err != nil {
			log.Fatal(err)
		}
	}
}

func Shutdown(
	quit chan os.Signal,
	svr server.Server,
) {

	q := <-quit
	log.Println("got signal:", q)

	log.Println("Shutting down server...")
	if err := svr.Stop(); err != nil {
		log.Fatal(err)
	}

	log.Println("service gracefully shutdown")
}
