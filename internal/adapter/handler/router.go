package handler

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Router struct {
	*gin.Engine
}

func NewRouter(
	productHandler ProductHandler,
) (*Router, error) {

	// CORS
	config := cors.DefaultConfig()
	allowedOrigins := os.Getenv("HTTP_ALLOWED_ORIGINS")
	originsList := strings.Split(allowedOrigins, ",")
	config.AllowOrigins = originsList

	router := gin.New()
	router.Use(gin.LoggerWithFormatter(customLogger), gin.Recovery(), cors.New(config))

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Welcome to the Go Hexagonal API",
		})

	})

	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "OK",
		})

	})

	// Product handler
	product := router.Group("/products")
	{
		product.GET("/", productHandler.ListProducts)
		product.GET("/:id", productHandler.GetProduct)
		product.POST("/", productHandler.CreateProduct)
		product.PUT("/:id", productHandler.UpdateProduct)
		product.DELETE("/:id", productHandler.DeleteProduct)
	}

	return &Router{
		router,
	}, nil

}

// Serve starts the HTTP server
func (r *Router) Serve(listenAddr string) error {
	return r.Run(listenAddr)
}

// customLogger is a custom Gin logger
func customLogger(param gin.LogFormatterParams) string {
	return fmt.Sprintf("[%s] - %s \"%s %s %s %d %s [%s]\"\n",
		param.TimeStamp.Format(time.RFC1123),
		param.ClientIP,
		param.Method,
		param.Path,
		param.Request.Proto,
		param.StatusCode,
		param.Latency.Round(time.Millisecond),
		param.Request.UserAgent(),
	)
}
