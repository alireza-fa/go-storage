CREATE TABLE users IF NOT EXISTS users(
    id SERIAL PRIMARY KEY,
    username VARCHAR(32) NOT NULL UNIQUE,
    email VARCHAR(64) NOT NULL, UNIQUE,
    hash_password VARCHAR(240) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
