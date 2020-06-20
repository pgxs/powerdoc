CREATE TABLE users
(
    uid               SERIAL PRIMARY KEY,
    username          char(50),
    password          char(128),
    name              char(30),
    email             char(100),
    last_logged_in_at timestamptz,
    created_at        timestamptz,
    deleted_at        timestamptz DEFAULT NULL
);
CREATE INDEX idx_deleted_at on users (deleted_at);