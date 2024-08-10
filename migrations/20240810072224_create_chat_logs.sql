-- +goose Up
create table if not exists chat_logs (
    id bigserial primary key,
    chat_id bigint not null,
    log varchar not null,
    created_at timestamp not null default now()
);

-- +goose Down
drop table if exists chat_logs;