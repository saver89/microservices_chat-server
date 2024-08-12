-- +goose Up
create table if not exists messages (
    id bigserial primary key,
    chat_id bigint not null,
    from_user text not null,
    text text not null,
    sent_at timestamp  not null,
    created_at timestamp default now() not null
);

-- +goose Down
drop table if exists messages;