//Filename: cmd/api/main.go

package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

// The Application Version Number
const version = "1.0.0"

// The Configuration Settings
type config struct {
	port int
	env  string // Development, Staging, Production, ETC.
}

// Dependency Injection
type application struct {
	config config
	logger *log.Logger
}

func main() {
	var cfg config

	// Read in flags that are needed to populate our config
	flag.IntVar(&cfg.port, "port", 4000, "API Server Port") // When using a struct we must use IntVar
	flag.StringVar(&cfg.env, "env", "development", "Enviorment ( Development | Staging | Production )")
	flag.Parse()

	//Create a logger
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	//Create an instance of our application struct
	// We are using the application struct for dependecy injection
	app := &application{
		config: cfg,
		logger: logger,
	}

	// Create our new servemux
	mux := http.NewServeMux()
	mux.HandleFunc("/v1/healthcheck", app.healthcheckHandler)

	// Create HTTP Server
	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.port),
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	// Start our Server
	logger.Printf("Starting %s Server on %s", cfg.env, srv.Addr)
	err := srv.ListenAndServe()
	logger.Fatal(err)

}
