create table if not exists users (
    id serial not null PRIMARY KEY,
    name varchar(24) not null,
    password varchar(24) not null,
    email varchar(50) not null UNIQUE
);

create table if not exists posts (
    id serial not null PRIMARY KEY,
    header varchar(64) not null,
    text varchar(5120) not null,
    author int not null,
    created_at timestamp not null default now(),
    FOREIGN KEY (author) references users(id)
);