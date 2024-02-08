CREATE TABLE IF NOT EXISTS Followers (
    id UUID PRIMARY KEY,
    follower_id INTEGER REFERENCES Users(id),
    followee_id INTEGER REFERENCES Users(id),
    status TEXT CHECK(status IN ('requested', 'accepted', 'declined')),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);