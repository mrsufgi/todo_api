package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	thd "github.com/mrsufgi/todo_api/internal/todos/delivery/http"
	tr "github.com/mrsufgi/todo_api/internal/todos/repository/pg"
	ts "github.com/mrsufgi/todo_api/internal/todos/service"

	"github.com/julienschmidt/httprouter"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	log "github.com/sirupsen/logrus"
)

func main() {
	// logger setup
	setupLogger(true)

	conf, err := sqlx.Connect("postgres", "host=database port=5432 user=postgres password=postgres dbname=demo sslmode=disable")
	if err != nil {
		log.Fatalln(err)
	}

	trepo := tr.NewPgTodosRepository(conf)
	mservice := ts.NewTodoService(trepo)

	port := ":3000"
	router := httprouter.New()
	router.GET("/metrics", Metrics(promhttp.Handler()))
	router.GET("/health", Health)

	s := &http.Server{
		Addr:         port,
		Handler:      router,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		// MaxHeaderBytes: 1 << 20,
	}

	// handler setup his own router
	thd.NewTodosHandler(router, mservice)

	go func() {
		log.Infof("start http server on port %s", port)
		if err := s.ListenAndServe(); err != nil {
			log.Println("HTTP server shutting down")
			if err != http.ErrServerClosed {
				log.Fatalf("closed unexpected error %v", err)
			}

			s.Close()
		}
	}()

	gracefulShutdown(s)
}

func Metrics(h http.Handler) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		h.ServeHTTP(w, r)
	}
}

func Health(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.WriteHeader(http.StatusOK)
}

func gracefulShutdown(s *http.Server) {
	interruptChan := make(chan os.Signal, 1)
	ctxTimeout := time.Second * 10

	signal.Notify(interruptChan, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-interruptChan

	ctx, cancel := context.WithTimeout(context.Background(), ctxTimeout)
	if err := s.Shutdown(ctx); err != nil {
		log.Fatalf("Error while shutting down %v", err)
	}

	cancel()
	os.Exit(0)
}

func setupLogger(debug bool) {
	log.SetFormatter(&log.JSONFormatter{})

	if debug {
		log.SetLevel(log.DebugLevel)
	}
}
