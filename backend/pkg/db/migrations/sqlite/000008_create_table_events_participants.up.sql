CREATE TABLE IF NOT EXISTS EventsParticipants (
    id INTEGER PRIMARY KEY,
    event_id INTEGER REFERENCES Events(id),
    member_id INTEGER REFERENCES Users(id),
    response TEXT CHECK(response IN ('going', 'not_going')),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);