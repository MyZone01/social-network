CREATE TABLE IF NOT EXISTS Events (
    id UUID PRIMARY KEY,
    group_id UUID REFERENCES Groups(id),
    title TEXT,
    description TEXT,
    datetime DATETIME,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);