
create table if not exists saved (
    id text primary key,
    user text not null,
    retro text not null,
    foreign key(user) references user(id) on delete cascade,
    foreign key(retro) references retro(id) on delete cascade
);
