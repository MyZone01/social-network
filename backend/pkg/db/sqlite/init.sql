-- Users Table
CREATE TABLE Users (
    id TEXT PRIMARY KEY UNIQUE,
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
CREATE TABLE Followers (
    id TEXT PRIMARY KEY,
    follower_id TEXT REFERENCES Users(id),
    followee_id TEXT REFERENCES Users(id),
    status TEXT CHECK(status IN ('requested', 'accepted', 'declined')),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Posts Table
CREATE TABLE Posts (
    id TEXT PRIMARY KEY,
    user_id TEXT REFERENCES Users(id),
    title TEXT,
    content TEXT,
    image_url TEXT,
    privacy TEXT CHECK(privacy IN ('public', 'private', 'almost private', 'unlisted')),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);

-- Groups Table
CREATE TABLE Groups (
    id TEXT PRIMARY KEY,
    title TEXT,
    description TEXT,
    banner_url TEXT,
    creator_id TEXT REFERENCES Users(id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Group Members Table
CREATE TABLE GroupMembers (
    id TEXT PRIMARY KEY,
    group_id TEXT REFERENCES Groups(id),
    member_id TEXT REFERENCES Users(id),
    status TEXT CHECK(status IN ('invited', 'requesting', 'accepted', 'declined')),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Group Posts Table
CREATE TABLE GroupPosts (
    id TEXT PRIMARY KEY,
    group_id TEXT REFERENCES Groups(id),
    post_id TEXT REFERENCES Posts(id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Events Table
CREATE TABLE Events (
    id TEXT PRIMARY KEY,
    group_id TEXT REFERENCES Groups(id),
    title TEXT,
    description TEXT,
    datetime DATETIME,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Events Participant Table
CREATE TABLE EventsParticipants (
    id TEXT PRIMARY KEY,
    event_id TEXT REFERENCES Events(id),
    member_id TEXT REFERENCES Users(id),
    response TEXT CHECK(response IN ('going', 'not_going')),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Private Messages Table
CREATE TABLE PrivateMessages (
    id TEXT PRIMARY KEY,
    sender_id TEXT REFERENCES Users(id),
    receiver_id TEXT REFERENCES Users(id),
    content TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Group Messages Table
CREATE TABLE GroupMessages (
    id TEXT PRIMARY KEY,
    group_id TEXT REFERENCES Groups(id),
    sender_id TEXT REFERENCES Users(id),
    content TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Notifications Table
CREATE TABLE Notifications (
    id TEXT PRIMARY KEY,
    user_id TEXT REFERENCES Users(id),
    type TEXT CHECK(type IN ('follow_request', 'group_invitation', 'new_message', 'new_event')),
    message TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Sessions Table
CREATE TABLE Sessions (
    id TEXT PRIMARY KEY,
    user_id TEXT  REFERENCES Users(id),
    session_token TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    expires_at TIMESTAMP,
    deleted_at TIMESTAMP
);