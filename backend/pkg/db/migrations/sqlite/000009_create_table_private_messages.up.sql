CREATE TABLE IF NOT EXISTS PrivateMessages (
    id INTEGER PRIMARY KEY,
    sender_id INTEGER REFERENCES Users(id),
    receiver_id INTEGER REFERENCES Users(id),
    content TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);