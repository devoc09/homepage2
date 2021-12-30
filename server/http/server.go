package http

import (
	"context"
	"fmt"

	"github.com/labstack/echo/v4"
)

type Server struct {
	server *echo.Echo
}

func NewServer() *Server {
	return &Server{
		echo.New(),
	}
}

func (s *Server) StartServer(port int) error {
	s.server.File("/", "public/index.html")
	s.server.Static("/static", "static")

	if err := s.server.Start(fmt.Sprintf(":%d", port)); err != nil {
		return fmt.Errorf("failed to serve: %w", err)
	}
	return nil
}

func (s *Server) StopServer(ctx context.Context) error {
	if err := s.server.Shutdown(ctx); err != nil {
		return fmt.Errorf("failed to shutdown: %w", err)
	}
	fmt.Println("Server Shutdown")
	return nil
}
