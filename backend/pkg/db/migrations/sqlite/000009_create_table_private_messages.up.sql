CREATE TABLE IF NOT EXISTS PrivateMessages (
    id UUID PRIMARY KEY,
    sender_id UUID REFERENCES Users(id),
    receiver_id UUID REFERENCES Users(id),
    content TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);