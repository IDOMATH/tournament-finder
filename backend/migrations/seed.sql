CREATE TABLE tournaments (
    id serial PRIMARY KEY,
    name VARCHAR(50),
    location_name VARCHAR(255),
    street_address VARCHAR(255),
    city VARCHAR(50),
    state VARCHAR(13),
    organizer_id INTEGER,
    is_full BOOLEAN,
    is_boy_varsity BOOLEAN,
    is_girls_varsity BOOLEAN,
    is_boys_jv BOOLEAN,
    is_girls_js BOOLEAN,
    is_boys_ms BOOLEAN,
    is_girls_ms BOOLEAN,
    is_boys_youth BOOLEAN,
    is_girls_youth BOOLEAN
)

CREATE TABLE users (
    id serial PRIMARY KEY,
    name VARCHAR(255)
    email VARCHAR(320),
    updated_at TIMESTAMP,
    created_at TIMESTAMP
)