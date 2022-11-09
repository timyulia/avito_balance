CREATE TABLE users
(
    id  int not null unique,
    amount int not null DEFAULT 0
);

CREATE TABLE reserved
(
    order_id int not null unique,
    user_id int not null,
    service_id int not null,
    amount int not null

);

CREATE TABLE report
(
    id serial not null unique,
    service_id int not null,
    amount int not null,
    date DATE
);

CREATE TABLE history
(
    id serial not null unique,
    user_id int not null,
    reason varchar(255),
    amount int not null,
    date DATE
);

CREATE TABLE service
(
    service_id int not null unique,
    name varchar(255)
);

