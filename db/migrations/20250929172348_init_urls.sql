-- +goose Up
-- +goose StatementBegin
create table url_mappings
(
    id         bigserial primary key,
    original_url TEXT not null UNIQUE,
    short_url char(10) not null unique,
    created_at timestamptz not null default now()

);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table url_mappings;
-- +goose StatementEnd