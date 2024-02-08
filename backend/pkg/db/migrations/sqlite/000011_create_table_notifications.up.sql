CREATE TABLE IF NOT EXISTS Notifications (
    id INTEGER PRIMARY KEY,
    user_id INTEGER REFERENCES Users(id),
    type TEXT CHECK(type IN ('follow_request', 'group_invitation', 'new_message', 'new_event')),
    message TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);