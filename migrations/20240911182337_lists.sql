-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
CREATE TABLE whitelists
(
    id     SERIAL PRIMARY KEY,
    subnet CIDR NOT NULL UNIQUE
);

CREATE TABLE blacklists
(
    id     SERIAL PRIMARY KEY,
    subnet CIDR NOT NULL UNIQUE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
DROP TABLE whitelists;
DROP TABLE blacklists;
-- +goose StatementEnd
