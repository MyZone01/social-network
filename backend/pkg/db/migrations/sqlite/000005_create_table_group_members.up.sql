CREATE TABLE IF NOT EXISTS GroupMembers (
    id UUID PRIMARY KEY,
    group_id UUID REFERENCES Groups(id),
    member_id UUID REFERENCES Users(id),
    status TEXT CHECK(status IN ('invited', 'requesting', 'accepted', 'declined')),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);