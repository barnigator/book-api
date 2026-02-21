package app

import (
	"context"
	"fmt"
	"net/http"
	"os/signal"
	"syscall"

	"github.com/barnigator/book-api/config"
	"github.com/barnigator/book-api/internal/delivery"
)

func NewServer(cfg *config.Config, h *delivery.Handler) *http.Server {
	mux := http.NewServeMux()
	delivery.RegisterRoutes(h, mux)

	return &http.Server{
		Addr:    cfg.Port,
		Handler: mux,
	}
}

func StartServer(srv *http.Server) {
	go func() {
		fmt.Printf("Listening on %s\n", srv.Addr)
		err := srv.ListenAndServe()
		if err != nil {
			fmt.Printf("Stopped listening: %v\n", err)
		}
	}()
}

func StopServer(srv *http.Server) {
	shutdown, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	<-shutdown.Done()

	fmt.Println("Shutting down server...")
	if err := srv.Shutdown(context.Background()); err != nil {
		fmt.Printf("Shutdown with error: %v\n", err)
	}

	fmt.Println("Shutdown complete")
}
