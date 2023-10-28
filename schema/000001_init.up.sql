CREATE TABLE users
(
    id serial not null unique,
    is_admin boolean,
    balance integer,
    name varchar not null,
    username varchar not null  unique,
    password_hash varchar not null
);

CREATE TABLE transport
(
    id serial not null unique ,
    owner_id serial REFERENCES users(id),
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

CREATE TABLE rent_history
(
    id serial not null,
    transport_id serial REFERENCES transport(id),
    user_id serial REFERENCES users(id),
    rent_type varchar,
    started_at date,
    ended_at date
);