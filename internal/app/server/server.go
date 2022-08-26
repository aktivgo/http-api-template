package server

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"wire-di/internal/domain"
)

func createChannel() (chan os.Signal, func()) {
	stopCh := make(chan os.Signal, 1)
	signal.Notify(stopCh, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)

	return stopCh, func() {
		close(stopCh)
	}
}

func start(server *http.Server) {
	log.Println("application started")
	if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		panic(err)
	} else {
		log.Println("application stopped gracefully")
	}
}

func shutdown(ctx context.Context, server *http.Server) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		panic(err)
	} else {
		log.Println("application shutdowned")
	}
}

func CreateUserHttpServer(
	port string,
	userHandler domain.UserHandler,
) {
	ctx := context.Background()

	router := buildRouter(
		userHandler,
	)

	s := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}
	go start(s)

	stopCh, closeCh := createChannel()
	defer closeCh()
	log.Println("notified:", <-stopCh)

	shutdown(ctx, s)
}
