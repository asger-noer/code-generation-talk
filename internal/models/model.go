package models

import (
	"encoding/json"
	"time"

	"github.com/AsgerNoer/talks/write-less-code/internal/database/postgresql"
	"github.com/AsgerNoer/talks/write-less-code/internal/httpserver/rest"
	"github.com/google/uuid"
)

type Status string

type Reminder struct {
	id          uuid.UUID
	status      Status
	title       string
	description string
	created     time.Time
}

func (r Reminder) MarshalJSON() ([]byte, error) {
	payload := rest.Reminder{
		Created:     r.created,
		Description: r.description,
		Id:          r.id,
		Status:      (rest.Status)(r.status),
		Title:       r.title,
	}

	return json.Marshal(payload)
}

func UnmarshallFromDB(row postgresql.Reminder) Reminder {
	return Reminder{
		id:          row.ID,
		status:      Status(row.Status),
		title:       row.Title,
		description: row.Description.String,
		created:     row.Created,
	}
}
