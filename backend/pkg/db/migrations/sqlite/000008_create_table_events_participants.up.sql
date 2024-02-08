CREATE TABLE IF NOT EXISTS EventsParticipants (
    id UUID PRIMARY KEY,
    event_id UUID REFERENCES Events(id),
    member_id UUID REFERENCES Users(id),
    response TEXT CHECK(response IN ('going', 'not_going')),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);