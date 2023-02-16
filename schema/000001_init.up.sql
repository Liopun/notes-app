CREATE TABLE users (
    id              serial not null unique,
    name            varchar(255) not null,
    username        varchar(255) not null unique,
    hashed_password varchar(255) not null
)

CREATE TABLE users_lists (
    id      serial not null unique,
    user_id int references users(id) on delete cascade  not null,
    list_id int references notes_lists(id) on delete cascade not null
)

CREATE TABLE notes_lists (
    id          serial not null unique,
    title       varchar(255)  not null,
    description varchar(255)
)

CREATE TABLE notes_items {
    id          serial not null unique,
    title       varchar(255)  not null,
    description varchar(255) not null,
    archived    boolean not null default false
}

CREATE TABLE lists_items {
    id      serial not null unique,
    item_id int references notes_items(id) on delete cascade not null,
    list_id int references notes_lists(id) on delete cascade not null
}