CREATE TABLE IF NOT EXISTS group_posts (
    id UUID PRIMARY KEY,
    group_id UUID REFERENCES groups(id),
    post_id UUID REFERENCES posts(id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);