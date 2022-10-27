package handler

import (
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sethvargo/go-envconfig"
	"go.uber.org/fx"
	"net/http"
)

var Module = fx.Module("handler",
	fx.Provide(
		NewConfig,
		NewGinEngine,
	),
	fx.Invoke(
		RegisterHandler,
	),
)

type Config struct {
	Port int `env:"PORT,default=8081"`
}

func NewConfig() (*Config, error) {
	var cfg Config
	err := envconfig.Process(context.Background(), &cfg)
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}

func NewGinEngine(
	lc fx.Lifecycle,
	cfg *Config,
) *gin.Engine {

	gin.DisableConsoleColor()
	gin.SetMode(gin.ReleaseMode)

	// Create Gin router
	r := gin.New()

	r.Use(
		gin.LoggerWithConfig(gin.LoggerConfig{
			SkipPaths: []string{"/health"},
		}),
	)

	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	r.Use(gin.Recovery())

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			addr := fmt.Sprintf(":%d", cfg.Port)
			// In production, we'd want to separate the Listen and Serve phases for
			// better error-handling.
			go func() {
				err := r.Run(addr)
				if err != nil && !errors.Is(err, http.ErrServerClosed) {
					panic(err)
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			// TODO see https://github.com/gin-gonic/examples/tree/master/graceful-shutdown/graceful-shutdown
			return nil
		},
	})

	return r
}

// RegisterHandler  handler
func RegisterHandler(r *gin.Engine) {
	r.POST("/query", func(c *gin.Context) {
		c.Writer.Write([]byte("querying..."))
	})

	// for k8s health check
	r.GET("/health", func(c *gin.Context) {
		c.Writer.Write([]byte("ok"))
	})
}
