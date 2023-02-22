package httpserver

import (
	"fmt"
	"net/http"

	"github.com/AsgerNoer/talks/write-less-code/internal/database"
	"github.com/AsgerNoer/talks/write-less-code/internal/httpserver/rest"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type server struct {
	repo *database.DB
}

// GetAllReminders implements ServerInterface
func (s *server) GetAllReminders(ctx echo.Context) error {
	req := ctx.Request()

	reminders, err := s.repo.GetReminders(req.Context())
	if err != nil {
		return fmt.Errorf("getting reminders: %w", err)
	}

	return ctx.JSON(http.StatusOK, reminders)
}

// GetsingleReminder implements ServerInterface
func (s *server) GetSingleReminder(ctx echo.Context, id uuid.UUID) error {
	req := ctx.Request()

	reminders, err := s.repo.GetReminders(req.Context(), id)
	if err != nil {
		return fmt.Errorf("getting reminders: %w", err)
	}

	return ctx.JSON(http.StatusOK, reminders[0])
}

func NewServer(repo *database.DB) *echo.Echo {
	echo := echo.New()

	server := &server{
		repo: repo,
	}

	echo.Use(middleware.Logger())

	// Register handlers from the generated service
	rest.RegisterHandlers(echo, server)

	return echo
}
