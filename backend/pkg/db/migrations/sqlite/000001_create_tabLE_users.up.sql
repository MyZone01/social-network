CREATE TABLE IF NOT EXISTS Users (
    id INTEGER PRIMARY KEY,
    email TEXT UNIQUE,
    password TEXT,
    first_name TEXT,
    last_name TEXT,
    date_of_birth DATE,
    avatar_image TEXT,
    nickname TEXT,
    about_me TEXT,
    is_public INTEGER,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    delete_at TIMESTAMP
);