create database invite;

create table if not exists guest (
     id serial unique,
     name varchar,
     email varchar,
     presence varchar,
     drink jsonb,
     food jsonb,
     music text,
     transfer varchar,
     accommodation varchar,
     companion integer references guest(id),
     date_created timestamp default now()
);