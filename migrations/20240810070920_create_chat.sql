-- +goose Up
create table if not exists chats (
    id bigserial primary key,
    name text not null,
    created_at timestamp default now() not null
);

create table if not exists chat_users (
    id bigserial primary key,
    chat_id bigint not null references chats(id),
    user_name text not null
);

-- +goose Down
drop table if exists chat_users;
drop table if exists chats;
