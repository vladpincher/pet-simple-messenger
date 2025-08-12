package server

import (
	"database/sql"
	"fmt"
	"time"

	"my-pet-simple-messenger/internal/config"
	"my-pet-simple-messenger/internal/logger"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

const (
	greetingWorld = "World"
)

type httpServer struct {
	route    *gin.Engine
	port     string
	database *sql.DB
	logger   *logger.Logger
}

func PortInitialization() string {
	viper.AutomaticEnv()
	viper.SetDefault("PORT", "8080")
	return fmt.Sprintf(":%s", viper.GetString("PORT"))
}

func NewServer(config *config.Config, logg *logger.Logger, db *sql.DB) *httpServer {

	r := gin.New()
	r.Use(gin.Recovery())

	s := &httpServer{
		route:    r,
		port:     config.Port,
		database: db,
		logger:   logg,
	}

	r.GET("/hello", s.withLogging(), s.HomeHandler)
	r.NoRoute(s.withLogging(), s.NotFoundHandler)

	return s
}

func (s *httpServer) Start() error {
	fmt.Printf("Starting server on %v\n", s.port)
	return s.route.Run(s.port)
}

func (s *httpServer) Close() {
	_ = s.logger.Sync()
}

func (s *httpServer) HomeHandler(c *gin.Context) {
	c.JSON(200, gin.H{"Hello": greetingWorld})
}

func (s *httpServer) NotFoundHandler(c *gin.Context) {
	c.JSON(404, gin.H{"error": "Not found"})
}

func (s *httpServer) withLogging() gin.HandlerFunc {
	return func(c *gin.Context) {
		ts := time.Now().UTC().Format(time.RFC3339)
		s.logger.Debugf("got request at %s method=%s path=%s", ts, c.Request.Method, c.Request.URL.Path)
		c.Next()
	}
}
