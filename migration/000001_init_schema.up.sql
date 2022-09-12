create table if not exists users (
    id serial not null PRIMARY KEY,
    name varchar(24) not null,
    password varchar(24) not null,
    email varchar(50) not null UNIQUE,
    token varchar(30) not null
);

create table if not exists posts (
    id serial not null PRIMARY KEY,
    header varchar(64) not null,
    text varchar(5120) not null,
    author int not null,
    created_at timestamp not null default now(),
    time_to_live timestamp,
    FOREIGN KEY (author) references users(id)
);

create table if not exists accesses (
     user_id int,
     post_id int not null,
     access varchar(5) not null,
     FOREIGN KEY (user_id) references users(id),
     FOREIGN KEY (post_id) references posts(id),
     CONSTRAINT ucCodes UNIQUE (user_id, post_id, access)
);