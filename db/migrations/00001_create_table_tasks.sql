-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
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
SELECT 'down SQL query';
DROP TABLE tasks;
-- +goose StatementEnd
