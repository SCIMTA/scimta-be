package main

import (
	"net/http"
	"os"
	"scimta-be/db"
	"scimta-be/router"
	"scimta-be/services"
	"strings"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	_ "scimta-be/docs" // docs is generated by Swag CLI, you have to import it.

	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger" // echo-swagger middleware
)

type Validator struct {
	validator *validator.Validate
}

func (v *Validator) Validate(i interface{}) error {
	if err := v.validator.Struct(i); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

// @title Swagger Example API
// @version 1.0
// @description Conduit API
// @title Conduit API

// @host eyeh:8080
// @BasePath /api

// @schemes http https
// @produce	application/json
// @consumes application/json

// @in header
// @name Authorization

func main() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal().Err(err).Msg("Error loading .env file")
	}

	// Init logger
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	logger := zerolog.New(os.Stdout)

	// Init Echo
	e := echo.New()
	e.Validator = &Validator{validator: validator.New()}
	// e.Logger.SetLevel(zerolog.DebugLevel)

	// Middleware
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogURI:    true,
		LogStatus: true,
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			logger.Info().
				Str("URI", v.URI).
				Int("status", v.Status).
				Msg("request")

			return nil
		},
	}))
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))
	e.Use(middleware.GzipWithConfig(middleware.GzipConfig{
		Skipper: func(c echo.Context) bool {
			return strings.Contains(c.Request().URL.Path, "docs")
		},
	}))

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	// Swagger
	e.GET("/docs/*", echoSwagger.WrapHandler)

	// API Group
	v1 := e.Group("/api")

	// Init DB
	d := db.New()
	db.AutoMigrate(d)

	// Services
	userServices := services.NewUserServices(d)

	// Routers
	router.NewUserRouter(v1, userServices)

	// Start server
	e.Logger.Fatal(e.Start(os.Getenv("BE_HOST") + ":" + os.Getenv("BE_PORT")))
}
