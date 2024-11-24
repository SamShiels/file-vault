package server

import (
	"fmt"
	"vault/internal/database/repository"
	"vault/internal/endpoint"
	"vault/internal/service"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Server struct {
	echo *echo.Echo
	api  *endpoint.Api
}

func (s *Server) Start() {
	s.echo = echo.New()

	s.echo.Use(middleware.Logger())
	s.echo.Use(middleware.Recover())

	fileRepo := &repository.FileRepository{}
	fileService := service.NewFileService(fileRepo)
	api := endpoint.NewApi(fileService)

	s.echo.Static("", "public")
	s.echo.POST("/upload", api.Upload)

	err := s.echo.Start(fmt.Sprintf(":%s", "8082"))
	if err != nil {
		s.echo.Logger.Fatal(err)
	}
}
