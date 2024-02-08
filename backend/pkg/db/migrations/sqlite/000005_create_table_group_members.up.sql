CREATE TABLE IF NOT EXISTS GroupMembers (
    id INTEGER PRIMARY KEY,
    group_id INTEGER REFERENCES Groups(id),
    member_id INTEGER REFERENCES Users(id),
    status TEXT CHECK(status IN ('invited', 'requesting', 'accepted', 'declined')),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);