create table if not exists user(
    id text primary key,
    name text not null,
    username text unique not null,
    password text not null
);

create table if not exists template(
    id text primary key,
    user text not null,
    categories text not null,
    foreign key(user) references user(id) on delete cascade
);

create table if not exists retro(
    id text primary key,
    user text not null,
    created text not null,
    template text not null,
    visible integer default 0,
    foreign key(user) references user(id) on delete cascade,
    foreign key(template) references template(id) on delete cascade
);

create table if not exists record(
    id text primary key,
    retro text not null,
    author text not null,
    category text not null,
    content text not null,
    likes integer default 0,
    foreign key(retro) references retro(id) on delete cascade
);

create table if not exists comment (
    id text primary key,
    record text not null,
    author text not null,
    likes integer default 0,
    content text not null,
    foreign key(record) references record(id) on delete cascade
)
