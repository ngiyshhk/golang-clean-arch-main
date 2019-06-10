
-- +migrate Up
CREATE TABLE users (
    id INT NOT NULL primary key,
    name VARCHAR(64) NOT NULL,
    age INT NOT NULL DEFAULT 0
);

-- +migrate Down
DROP TABLE users;
