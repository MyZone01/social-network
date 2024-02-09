-- Users Table
CREATE TABLE Users (
    id UUID PRIMARY KEY UNIQUE,
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
    id UUID PRIMARY KEY,
    follower_id UUID REFERENCES Users(id),
    followee_id UUID REFERENCES Users(id),
    status TEXT CHECK(status IN ('requested', 'accepted', 'declined')),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Posts Table
CREATE TABLE Posts (
    id UUID PRIMARY KEY,
    user_id UUID REFERENCES Users(id),
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
    id UUID PRIMARY KEY,
    title TEXT,
    description TEXT,
    banner_url TEXT,
    creator_id UUID REFERENCES Users(id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Group Members Table
CREATE TABLE GroupMembers (
    id UUID PRIMARY KEY,
    group_id UUID REFERENCES Groups(id),
    member_id UUID REFERENCES Users(id),
    status TEXT CHECK(status IN ('invited', 'requesting', 'accepted', 'declined')),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Group Posts Table
CREATE TABLE GroupPosts (
    id UUID PRIMARY KEY,
    group_id UUID REFERENCES Groups(id),
    post_id UUID REFERENCES Posts(id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Events Table
CREATE TABLE Events (
    id UUID PRIMARY KEY,
    group_id UUID REFERENCES Groups(id),
    title TEXT,
    description TEXT,
    datetime DATETIME,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Events Participant Table
CREATE TABLE EventsParticipants (
    id UUID PRIMARY KEY,
    event_id UUID REFERENCES Events(id),
    member_id UUID REFERENCES Users(id),
    response TEXT CHECK(response IN ('going', 'not_going')),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Private Messages Table
CREATE TABLE PrivateMessages (
    id UUID PRIMARY KEY,
    sender_id UUID REFERENCES Users(id),
    receiver_id UUID REFERENCES Users(id),
    content TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Group Messages Table
CREATE TABLE GroupMessages (
    id UUID PRIMARY KEY,
    group_id UUID REFERENCES Groups(id),
    sender_id UUID REFERENCES Users(id),
    content TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Notifications Table
CREATE TABLE Notifications (
    id UUID PRIMARY KEY,
    user_id UUID REFERENCES Users(id),
    type TEXT CHECK(type IN ('follow_request', 'group_invitation', 'new_message', 'new_event')),
    message TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Sessions Table
CREATE TABLE Sessions (
    id UUID PRIMARY KEY,
    user_id UUID  REFERENCES Users(id),
    session_token TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    expires_at TIMESTAMP,
    deleted_at TIMESTAMP
);