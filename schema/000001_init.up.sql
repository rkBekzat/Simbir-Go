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
    id serial not null unique ,
    transport_id int REFERENCES transport(id) on delete cascade not null ,
    user_id int REFERENCES users(id) on delete cascade not null ,
    rent_type varchar not null,
    renting_ended boolean,
    started_at timestamp,
    ended_at timestamp
);