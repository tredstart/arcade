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
    foreign key(user) references user(id)
);

create table if not exists retro(
    id text primary key,
    user text not null,
    created text not null,
    template text not null,
    url text unique not null,
    foreign key(user) references user(id),
    foreign key(template) references template(id)
);

create table if not exists record(
    id text primary key,
    retro text not null,
    author text not null,
    category text not null,
    content text not null,
    foreign key(retro) references retro(id)
);
