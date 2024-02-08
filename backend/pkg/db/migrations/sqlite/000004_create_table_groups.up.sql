CREATE TABLE IF NOT EXISTS Groups (
    id UUID PRIMARY KEY,
    title TEXT,
    description TEXT,
    banner_url TEXT,
    creator_id UUID REFERENCES Users(id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);