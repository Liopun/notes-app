CREATE TABLE users (
    id              SERIAL NOT NULL UNIQUE,
    name            VARCHAR(255) NOT NULL,
    username        VARCHAR(255) NOT NULL UNIQUE,
    hashed_password VARCHAR(255) NOT NULL
)

CREATE TABLE users_lists (
    id      SERIAL NOT NULL UNIQUE,
    user_id int REFERENCES users(id) ON DELETE CASCADE  NOT NULL,
    list_id int REFERENCES notes_lists(id) ON DELETE CASCADE NOT NULL
)

CREATE TABLE notes_lists (
    id          SERIAL NOT NULL UNIQUE,
    title       VARCHAR(255)  NOT NULL,
    description TEXT
)

CREATE TABLE notes_items {
    id          SERIAL NOT NULL UNIQUE,
    title       VARCHAR(255)  NOT NULL,
    description TEXT NOT NULL,
    archived    boolean NOT NULL default false
}

CREATE TABLE lists_items {
    id      SERIAL NOT NULL UNIQUE,
    item_id int REFERENCES notes_items(id) ON DELETE CASCADE NOT NULL,
    list_id int REFERENCES notes_lists(id) ON DELETE CASCADE NOT NULL
}