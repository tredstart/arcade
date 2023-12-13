alter table record
    add column likes integer default 0;

alter table retro
    add column visible integer default 0;

create table if not exists comments (
    id text primary key,
    record text not null,
    author text not null,
    likes integer default 0,
    content text not null,
    foreign key(record) references record(id)
)
