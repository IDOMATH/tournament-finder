CREATE TABLE tournaments (
    id serial PRIMARY KEY,
    name VARCHAR(50),
    location_name VARCHAR(50),
    location VARCHAR(255),
    organizer_name VARCHAR(50),
    organizer_email VARCHAR(320),
    age_division INTEGER
)