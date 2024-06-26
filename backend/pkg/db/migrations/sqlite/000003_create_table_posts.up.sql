CREATE TABLE IF NOT EXISTS posts (
    id UUID PRIMARY KEY,
    user_id INTEGER REFERENCES users(id),
    title TEXT,
    content TEXT,
    image_url TEXT,
    privacy TEXT CHECK(privacy = 'public' OR privacy = 'private' OR privacy = 'almost private' OR privacy = 'unlisted'),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);
