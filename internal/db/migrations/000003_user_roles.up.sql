CREATE TABLE IF NOT EXISTS user_roles(
    id SERIAL PRIMARY KEY,
    user_id INTEGER REFERENCES users (id),
    role_id INTEGER REFERENCES roles (id)
);