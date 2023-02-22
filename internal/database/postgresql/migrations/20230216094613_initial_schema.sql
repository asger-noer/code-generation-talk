-- +goose Up
-- +goose StatementBegin
CREATE TYPE reminder_status AS ENUM (
    'created',
    'doing',
    'done'
);

CREATE TABLE IF NOT EXISTS reminders (
    "id" UUID DEFAULT gen_random_uuid() NOT NULL,
    "title" TEXT NOT NULL,
    "description" TEXT,
    "status" reminder_status DEFAULT 'created' NOT NULL,
    "created" TIMESTAMP DEFAULT NOW() NOT NULL,
    PRIMARY KEY(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS (
    reminders
);

DROP TYPE reminder_status;
-- +goose StatementEnd
