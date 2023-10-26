CREATE TABLE users
(
    id serial not null unique,
    name varchar not null,
    username varchar not null  unique,
    password_hash varchar not null
);

CREATE TABLE transport
(
    id 
    can_be_rented boolean,
    transport_type varchar,
    model varchar,
    color varchar,
    identifier varchar,
    description varchar not null,
    latitude real,
    longitude real,
    minute_price real,
    day_price real
);