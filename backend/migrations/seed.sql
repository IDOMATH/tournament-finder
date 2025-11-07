CREATE TABLE tournaments (
    id serial PRIMARY KEY,
    name VARCHAR(50),
    location_name VARCHAR(255),
    street_address VARCHAR(255),
    city VARCHAR(50),
    state VARCHAR(13),
    organizer_id INTEGER,
    boy_varsity INTEGER,
    girls_varsity INTEGER,
    boys_jv INTEGER,
    girls_js INTEGER,
    boys_ms INTEGER,
    girls_ms INTEGER,
    boys_youth INTEGER,
    girls_youth INTEGER
)

CREATE TABLE users (
    id serial PRIMARY KEY,
    name VARCHAR(255),
    email VARCHAR(320) UNIQUE,
    password_hash VARCHAR(255),
    is_organizer BOOLEAN,
    is_coach BOOLEAN,
    updated_at TIMESTAMP,
    created_at TIMESTAMP
)

CREATE TABLE schedule (
    coach_id INTEGER,
    tournament_id INTEGER
)