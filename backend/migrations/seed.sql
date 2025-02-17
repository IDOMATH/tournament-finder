CREATE TABLE tournaments (
    id serial PRIMARY KEY,
    name VARCHAR(50),
    location_name VARCHAR(255),
    location_address VARCHAR(255),
    organizer_id INTEGER,
    is_full BOOLEAN,
    age_division INTEGER
)

CREATE TABLE users (
    id serial PRIMARY KEY,
    name VARCHAR(255)
    email VARCHAR(320),
    updated_at TIMESTAMP,
    created_at TIMESTAMP
)