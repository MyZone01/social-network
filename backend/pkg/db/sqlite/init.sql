-- Users Table
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

-- Followers Table
CREATE TABLE IF NOT EXISTS Followers (
    id INTEGER PRIMARY KEY,
    follower_id INTEGER REFERENCES Users(id),
    followee_id INTEGER REFERENCES Users(id),
    status TEXT CHECK(status IN ('requested', 'accepted', 'declined')),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Posts Table
CREATE TABLE IF NOT EXISTS Posts (
    id INTEGER PRIMARY KEY,
    user_id INTEGER REFERENCES Users(id),
    title TEXT,
    content TEXT,
    image_url TEXT,
    privacy TEXT CHECK(privacy IN ('public', 'private', 'almost private', 'unlisted')),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);

-- Groups Table
CREATE TABLE IF NOT EXISTS Groups (
    id INTEGER PRIMARY KEY,
    title TEXT,
    description TEXT,
    banner_url TEXT,
    creator_id INTEGER REFERENCES Users(id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Group Members Table
CREATE TABLE IF NOT EXISTS GroupMembers (
    id INTEGER PRIMARY KEY,
    group_id INTEGER REFERENCES Groups(id),
    member_id INTEGER REFERENCES Users(id),
    status TEXT CHECK(status IN ('invited', 'requesting', 'accepted', 'declined')),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Group Posts Table
CREATE TABLE IF NOT EXISTS GroupPosts (
    id INTEGER PRIMARY KEY,
    group_id INTEGER REFERENCES Groups(id),
    post_id INTEGER REFERENCES Posts(id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Events Table
CREATE TABLE IF NOT EXISTS Events (
    id INTEGER PRIMARY KEY,
    group_id INTEGER REFERENCES Groups(id),
    title TEXT,
    description TEXT,
    datetime DATETIME,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Events Participant Table
CREATE TABLE IF NOT EXISTS EventsParticipants (
    id INTEGER PRIMARY KEY,
    event_id INTEGER REFERENCES Events(id),
    member_id INTEGER REFERENCES Users(id),
    response TEXT CHECK(response IN ('going', 'not_going')),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Private Messages Table
CREATE TABLE IF NOT EXISTS PrivateMessages (
    id INTEGER PRIMARY KEY,
    sender_id INTEGER REFERENCES Users(id),
    receiver_id INTEGER REFERENCES Users(id),
    content TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Group Messages Table
CREATE TABLE IF NOT EXISTS GroupMessages (
    id INTEGER PRIMARY KEY,
    group_id INTEGER REFERENCES Groups(id),
    sender_id INTEGER REFERENCES Users(id),
    content TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Notifications Table
CREATE TABLE IF NOT EXISTS Notifications (
    id INTEGER PRIMARY KEY,
    user_id INTEGER REFERENCES Users(id),
    type TEXT CHECK(type IN ('follow_request', 'group_invitation', 'new_message', 'new_event')),
    message TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Sessions Table
CREATE TABLE IF NOT EXISTS Sessions (
    id INTEGER PRIMARY KEY,
    user_id INTEGER REFERENCES Users(id),
    session_token TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    expires_at TIMESTAMP,
    deleted_at TIMESTAMP
);