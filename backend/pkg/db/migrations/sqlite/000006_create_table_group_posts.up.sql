CREATE TABLE IF NOT EXISTS group_posts (
    id UUID PRIMARY KEY,
    group_id UUID REFERENCES groups(id),
    title TEXT,
    content TEXT,
    image_url TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);