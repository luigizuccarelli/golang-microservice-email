package main

import (
	"github.com/gorilla/mux"
	"github.com/microlib/simple"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

var (
	logger simple.Logger
)

func startHttpServer() *http.Server {

	srv := &http.Server{Addr: ":" + os.Getenv("SERVER_PORT")}

	r := mux.NewRouter()
	r.HandleFunc("/api/v1/mail/send", MiddlewareSendMail).Methods("POST")
	r.HandleFunc("/api/v2/sys/info/isalive", IsAlive).Methods("GET")
	http.Handle("/", r)

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			logger.Error("Httpserver: ListenAndServe() error: " + err.Error())
		}
	}()

	return srv
}

func main() {
	err := ValidateEnvars()
	if err != nil {
		os.Exit(-1)
	}
	// read the log level
	logger.Level = os.Getenv("LOG_LEVEL")
	srv := startHttpServer()
	logger.Info("Starting server on port " + os.Getenv("SERVER_PORT"))
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	exit_chan := make(chan int)

	go func() {
		for {
			s := <-c
			switch s {
			case syscall.SIGHUP:
				exit_chan <- 0
			case syscall.SIGINT:
				exit_chan <- 0
			case syscall.SIGTERM:
				exit_chan <- 0
			case syscall.SIGQUIT:
				exit_chan <- 0
			default:
				exit_chan <- 1
			}
		}
	}()

	code := <-exit_chan

	if err := srv.Shutdown(nil); err != nil {
		panic(err)
	}
	logger.Info("Server shutdown successfully")
	os.Exit(code)
}
