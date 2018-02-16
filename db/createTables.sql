CREATE TABLE subjects (
    id serial primary key,
    courseID serial,
    name varchar(128) not null
);
CREATE TABLE courses (
    id serial primary key,
    name varchar(256) not null
);