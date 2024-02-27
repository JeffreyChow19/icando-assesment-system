package server

import (
	"context"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.uber.org/fx"
	"icando/internal/middleware"
	"icando/internal/route"
	"icando/lib"
	"icando/utils/logger"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

var Module = fx.Module("server", fx.Provide(NewServer))

type Server struct {
	Engine *gin.Engine
	Port   int
	Routes *route.Routes
}

func NewServer(routes *route.Routes, config *lib.Config) *Server {
	gin.SetMode(gin.ReleaseMode)
	engine := gin.New()
	engine.Use(gin.Recovery())
	engine.Use(middleware.LoggingMiddleware())
	engine.Use(
		cors.New(
			cors.Config{
				AllowOrigins:     strings.Split(config.Cors, ","),
				AllowMethods:     []string{"GET", "POST", "DELETE", "OPTIONS", "PUT", "PATCH"},
				AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
				ExposeHeaders:    []string{"Content-Length"},
				AllowCredentials: true,
				MaxAge:           12 * time.Hour,
			},
		),
	)
	routes.Setup(engine)
	return &Server{
		Engine: engine,
		Port:   config.ServicePort,
		Routes: routes,
	}
}

func (s *Server) Run() {
	err := godotenv.Load()
	if err != nil {
		logger.Log.Error("Error loading .env file")
	}

	address := fmt.Sprintf(":%d", s.Port)

	server := &http.Server{
		Addr:    address,
		Handler: s.Engine,
	}

	go func() {
		if err := s.startServer(server); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	logger.Log.Info("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		logger.Log.Fatal("Server Shutdown: ", err)
	}

	select {
	case <-ctx.Done():
		logger.Log.Info("Timeout 2s")
	}
	logger.Log.Info("Server exiting...")
}

func (s *Server) RunForTest() {
	err := godotenv.Load()
	if err != nil {
		logger.Log.Error("Error loading .env file")
	}

	address := fmt.Sprintf(":%d", s.Port)

	server := &http.Server{
		Addr:    address,
		Handler: s.Engine,
	}

	go func() {
		if err := s.startServer(server); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()
}

func (s *Server) startServer(server *http.Server) error {
	logger.Log.Info(fmt.Sprintf("Server started on port %d", s.Port))
	return server.ListenAndServe()
}
