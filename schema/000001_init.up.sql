CREATE TABLE users
(
    id serial not null unique,
    name varchar not null,
    username varchar not null  unique,
    password_hash varchar not null
);
