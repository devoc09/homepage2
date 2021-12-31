package server

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/devoc09/homepage2/server/http"
)

func Run() {
	os.Exit(run(context.Background()))
}

func run(ctx context.Context) int {
	termCh := make(chan os.Signal, 1)
	trapSignals := []os.Signal{syscall.SIGTERM, syscall.SIGINT}
	signal.Notify(termCh, trapSignals...)

	s := http.NewServer()
	errCh := make(chan error, 1)

	go func() {
		errCh <- s.StartServer(443)
	}()

	select {
	case <-termCh:
		s.StopServer(ctx)
		return 0
	case <-errCh:
		return 1
	}
}
