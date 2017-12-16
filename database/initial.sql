create table items (
    id serial primary key,
    title varchar(255) not null,
    description text not null,
    completed boolean default false not null,
    created_at timestamp default current_timestamp
);