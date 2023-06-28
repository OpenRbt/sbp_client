create table users (
    id uuid default gen_random_uuid() not null constraint user_id_pk primary key,
    identity_uid text not null,
    deleted boolean default false not null
);

create table wash_servers (
    id uuid default gen_random_uuid() not null constraint wash_servers_id_pk primary key,
    owner uuid not null constraint wash_servers_owner_fk references users,
    title text not null,
    service_key text,
    terminal_key text,
    terminal_password text,
    description text not null,
    deleted boolean default false not null
);

CREATE TABLE transactions (
    id uuid default gen_random_uuid() not null constraint transactions_id_pk primary key,
    server_id TEXT NOT NULL,
    post_id TEXT NOT NULL,
    payment_id_bank TEXT,
    amount Integer NOT NULL,
    status TEXT,
    data_create timestamp default now(),
    data_update timestamp default now()
);