create database invite;

create table if not exists guest (
     id serial unique,
     name varchar,
     email varchar,
     presence bool,
     drink jsonb,
     food jsonb,
     music text,
     transfer bool,
     accommodation bool
);