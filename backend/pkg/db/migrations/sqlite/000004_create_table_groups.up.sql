CREATE TABLE IF NOT EXISTS Groups (
    id INTEGER PRIMARY KEY,
    title TEXT,
    description TEXT,
    banner_url TEXT,
    creator_id INTEGER REFERENCES Users(id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);