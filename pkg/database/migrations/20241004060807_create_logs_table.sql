-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS logs (
    id SERIAL PRIMARY KEY,
    content TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS logs;
-- +goose StatementEnd
