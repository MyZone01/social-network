CREATE TABLE IF NOT EXISTS GroupMessages (
    id UUID PRIMARY KEY,
    group_id UUID REFERENCES Groups(id),
    sender_id UUID REFERENCES Users(id),
    content TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);