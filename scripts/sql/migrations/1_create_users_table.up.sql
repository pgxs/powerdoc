CREATE TABLE users
(
    uid               INTEGER PRIMARY KEY AUTO_INCREMENT,
    username          char(50),
    password          char(100),
    name              char(30),
    email             char(100),
    last_logged_in_at datetime,
    created_at        datetime,
    deleted_at        datetime DEFAULT NULL
);
CREATE INDEX idx_deleted_at on users (deleted_at);