CREATE TABLE IF NOT EXISTS Sessions (
    id INTEGER PRIMARY KEY,
    user_id INTEGER REFERENCES Users(id),
    session_token TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    expires_at TIMESTAMP,
    deleted_at TIMESTAMP
);