package server

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
)

const (
	wsReadBufferSize  = 1024
	wsWriteBufferSize = 1024
	shutdownTimeout   = 30 * time.Second
	writeTimeout      = 60 * time.Second
)

type Server struct {
	server *http.Server
	Router *chi.Mux
}

type Opts struct {
	Addr string
}

func NewServer(opts Opts) *Server {
	router := chi.NewRouter()

	server := &http.Server{
		Addr:         opts.Addr,
		Handler:      router,
		WriteTimeout: writeTimeout,
	}

	return &Server{
		server: server,
		Router: router,
	}
}

func (s *Server) ListenAndServe() error {
	idleConnsClosed := make(chan struct{})

	go func() {
		sigint := make(chan os.Signal, 1)

		// interrupt signal sent from terminal
		signal.Notify(sigint, os.Interrupt)
		// sigterm signal sent from kubernetes
		signal.Notify(sigint, syscall.SIGTERM)

		<-sigint

		// We received an interrupt signal, shut down.
		ctx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
		defer cancel()

		if err := s.server.Shutdown(ctx); err != nil {
			// Error from closing listeners, or context timeout:
			log.Printf("HTTP server Shutdown: %v", err)
		}

		close(idleConnsClosed)
	}()

	if err := s.server.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
		return fmt.Errorf("http server ListenAndServe returned error: %w", err)
	}

	<-idleConnsClosed

	return nil
}
