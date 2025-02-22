CREATE TABLE users (
    user_id BIGINT,
    username VARCHAR(50) UNIQUE NOT NULL,
    fullname VARCHAR(50) NOT NULL,
    hash_pw TEXT NOT NULL,
    created_at BIGINT,
    PRIMARY KEY (user_id)
);