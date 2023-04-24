-- +goose Up
-- +goose StatementBegin
CREATE TABLE tasks(
                      id SERIAL PRIMARY KEY,
                      status TEXT,
                      name TEXT,
                      description TEXT,
                      performer TEXT,
                      deadline TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE tasks;
-- +goose StatementEnd
