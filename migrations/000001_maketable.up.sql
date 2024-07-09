CREATE TABLE users (
    id UUID PRIMARY KEY NOT NULL,
    name VARCHAR(65) NOT NULL,
    username VARCHAR(30) NOT NULL,
    bio VARCHAR(300),
    birth_day DATE,
    email VARCHAR(40) NOT NULL,
    password TEXT NOT NULL,
    avatar TEXT, 
    coint INT DEFAULT 0,
    score INT DEFAULT 0,
    refresh_token TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT NULL
);

CREATE TABLE admins (
    id SERIAL PRIMARY KEY,
    email VARCHAR(40),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE socials(
    location_name VARCHAR(35),
    location_url TEXT,
    education_name VARCHAR(35),
    education_url TEXT,
    telegram_name VARCHAR(35),
    telegram_url TEXT,
    twitter_name VARCHAR(35),
    twitter_url TEXT,
    instagram_name VARCHAR(35),
    instagram_url TEXT,
    youtube_name VARCHAR(35),
    youtube_url TEXT,
    linkedin_name VARCHAR(35),
    linkedin_url TEXT,
    website_name VARCHAR(35),
    website_url TEXT,
    user_id UUID REFERENCES users(id)
);

CREATE TYPE badgeme AS ENUM ('month', 'extra');
CREATE TABLE badges(
    id SERIAL PRIMARY KEY,
    name VARCHAR(100),
    badge_date DATE,
    badge_type badgeme,
    picture TEXT
);


CREATE TABLE user_badge(
    user_id UUID REFERENCES users(id),
    badge_id INT REFERENCES badges(id)
);

-- Course


CREATE TABLE languages(
    id SERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    picture TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT NULL
);

CREATE TABLE levels(
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    real_level INT NOT NULL,
    picture TEXT NOT NULL,
    language_id INT REFERENCES languages(id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT NULL
);

CREATE TABLE topics (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    level_id INT REFERENCES levels(id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT NULL
);

CREATE TABLE lessons (
    id SERIAL PRIMARY KEY,
    lesson_type VARCHAR(1) NOT NULL,
    topic_id INT REFERENCES topics(id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT NULL
);

CREATE TABLE contents (
    id SERIAL PRIMARY KEY,
    lesson_id INT REFERENCES lessons(id),
    gentype SMALLINT,
    title TEXT,
    question TEXT,
    text_data TEXT,
    arr_text TEXT[],
    correct_answer INTEGER,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT NULL
);

CREATE TABLE content_files (
    id SERIAL PRIMARY KEY,
    content_id INT REFERENCES contents(id),
    sound_data TEXT,
    image_data TEXT,
    video_data TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT NULL
);

CREATE TABLE user_language (
    id SERIAL PRIMARY KEY,
    status BOOLEAN DEFAULT TRUE,
    user_id UUID REFERENCES users(id),
    language_id INT REFERENCES languages(id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE user_level (
    id SERIAL PRIMARY KEY,
    score INT DEFAULT 0 CHECK(score <= 100),
    status BOOLEAN DEFAULT TRUE,
    user_id UUID REFERENCES users(id),
    level_id INT REFERENCES levels(id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE user_topic (
    id SERIAL PRIMARY KEY,
    score INT DEFAULT 0 CHECK(score <= 100),
    user_id UUID REFERENCES users(id),
    topic_id INT REFERENCES topics(id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE user_lesson (
    id SERIAL PRIMARY KEY,
    score INT DEFAULT 0 CHECK(score <= 100),
    user_id UUID REFERENCES users(id),
    lesson_id INT REFERENCES lessons(id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE activitys (
    id SERIAL PRIMARY KEY,
    score INTEGER DEFAULT 0,
    lesson_id INT REFERENCES lessons(id),
    user_id UUID REFERENCES users(id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Place for certificate

CREATE TABLE certificates (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    pdfile TEXT NOT NULL,
    level_id INT REFERENCES levels(id),
    user_id UUID REFERENCES users(id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE books (
    id SERIAL PRIMARY KEY,
    name VARCHAR(65),
    picture TEXT,
    book_file TEXT,
    level_id INT REFERENCES levels(id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT NULL
);


CREATE TABLE user_messages (
    id SERIAL PRIMARY KEY,
    user_id UUID REFERENCES users(id),
    content TEXT NOT NULL, -- end of content we should add where is form // Language -> Level -> Topic -> Lesson
    status BOOLEAN, --default false
    comment TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT NULL
);


CREATE TABLE admin_messages (
    id SERIAL PRIMARY KEY,
    user_id UUID REFERENCES users(id),
    content TEXT NOT NULL, -- end of content we should add where is form // Language -> Level -> Topic -> Lesson
    status BOOLEAN, --default false
    comment TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT NULL
);