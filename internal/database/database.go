package database

import (
	"context"
	"database/sql"
	"embed"
	"fmt"

	"github.com/AsgerNoer/talks/write-less-code/internal/database/postgresql"
	"github.com/AsgerNoer/talks/write-less-code/internal/models"
	"github.com/google/uuid"
	"github.com/pressly/goose/v3"

	// require the lib/pg package for the postgresql driver
	_ "github.com/lib/pq"
)

//go:embed postgresql/migrations/*.sql
var migrationFS embed.FS

type DB struct {
	db      *sql.DB
	queries *postgresql.Queries
}

func (db *DB) GetReminders(ctx context.Context, ids ...uuid.UUID) (reminders []models.Reminder, err error) {
	var rows []postgresql.Reminder

	switch {
	case ids != nil:
		rows, err = db.queries.GetRemindersWithId(ctx, ids)
		if err != nil {
			return reminders, fmt.Errorf("getting reminders based on ids from database: %w", err)
		}
	default:
		rows, err = db.queries.GetAllReminders(ctx)
		if err != nil {
			return reminders, fmt.Errorf("getting all reminders from database: %w", err)
		}
	}

	for _, row := range rows {
		reminders = append(reminders, models.UnmarshallFromDB(row))
	}

	return reminders, nil
}

func New(dialect string, db *sql.DB) (*DB, error) {
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("pinging database: %w", err)
	}

	goose.SetBaseFS(migrationFS)

	if err := goose.SetDialect(dialect); err != nil {
		return nil, fmt.Errorf("setting dialect for migration: %w", err)
	}

	if err := goose.Up(db, "postgresql/migrations"); err != nil {
		return nil, fmt.Errorf("appling latest migration: %w", err)
	}

	return &DB{
		db:      db,
		queries: postgresql.New(db),
	}, nil
}
