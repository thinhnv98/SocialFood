create table if not exists account(
    id serial primary key,
    email text,
    password text,
    type text
);

insert into account(email, password, type)
values ('Admin', '1', 'vendor'),
       ('Thinh', '1', 'customer');