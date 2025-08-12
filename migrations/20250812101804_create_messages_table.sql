-- +goose Up
CREATE TABLE messages (
    id SERIAL PRIMARY KEY,
    sender VARCHAR(255) NOT NULL,
    recipient VARCHAR(255) NOT NULL,
    content TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT NOW()
);

-- +goose Down
DROP TABLE messages;
