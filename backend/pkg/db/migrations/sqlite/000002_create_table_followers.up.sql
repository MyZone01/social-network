CREATE TABLE IF NOT EXISTS followers (
    id UUID PRIMARY KEY,
    follower_id INTEGER REFERENCES users(id),
    followee_id INTEGER REFERENCES users(id),
    status TEXT CHECK(status IN ('requested', 'accepted', 'declined')),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);