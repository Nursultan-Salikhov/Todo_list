CREATE TABLE tasks
(
    id serial not null unique,
    name varchar(1000) not null,
    done boolean not null default false
);