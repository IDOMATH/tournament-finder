CREATE TABLE tournaments (
    id serial PRIMARY KEY,
    name VARCHAR(50),
    location_name VARCHAR(255),
    location_address VARCHAR(255),
    organizer_name VARCHAR(50),
    organizer_email VARCHAR(320),
    age_division INTEGER
)

CREATE TABLE users (
    id serial PRIMARY KEY,
    name VARCHAR(255)
    email VARCHAR(320),
    updated_at TIMESTAMP,
    created_at TIMESTAMP
)