CREATE TABLE IF NOT EXISTS GroupMessages (
    id INTEGER PRIMARY KEY,
    group_id INTEGER REFERENCES Groups(id),
    sender_id INTEGER REFERENCES Users(id),
    content TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);