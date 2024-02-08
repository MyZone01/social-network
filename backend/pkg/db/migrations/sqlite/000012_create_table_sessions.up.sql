CREATE TABLE IF NOT EXISTS Sessions (
    id UUID PRIMARY KEY,
    user_id UUID REFERENCES Users(id),
    session_token TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    expires_at TIMESTAMP,
    deleted_at TIMESTAMP
);