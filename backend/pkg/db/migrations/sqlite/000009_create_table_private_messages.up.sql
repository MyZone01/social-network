CREATE TABLE IF NOT EXISTS UUID (
    id UUID PRIMARY KEY,
    sender_id UUID REFERENCES Users(id),
    receiver_id UUID REFERENCES Users(id),
    content TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);