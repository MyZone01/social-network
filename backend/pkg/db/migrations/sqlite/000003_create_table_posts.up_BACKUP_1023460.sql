<<<<<<< HEAD
CREATE TABLE IF NOT EXISTS posts (
    id UUID PRIMARY KEY,
    user_id INTEGER REFERENCES users(id),
=======
CREATE TABLE IF NOT EXISTS Posts (
    id INTEGER PRIMARY KEY,
    user_id INTEGER REFERENCES Users(id),
>>>>>>> master
    title TEXT,
    content TEXT,
    image_url TEXT,
    privacy TEXT CHECK(privacy IN ('public', 'private', 'almost private', 'unlisted')),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);