package http

import (
	"context"
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"golang.org/x/crypto/acme/autocert"
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

	s.server.AutoTLSManager.Cache = autocert.DirCache("/var/www/.cache")
	s.server.Pre(middleware.HTTPSRedirect())

	// if err := s.server.Start(fmt.Sprintf(":%d", port)); err != nil {
	// 	return fmt.Errorf("failed to serve: %w", err)
	// }
	if err := s.server.StartAutoTLS(fmt.Sprintf(":%d", port)); err != nil {
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
