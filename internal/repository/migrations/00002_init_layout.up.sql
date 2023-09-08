-- users
CREATE TYPE user_role AS ENUM ('user', 'admin');

create table users (
    id uuid default gen_random_uuid() not null constraint user_id_pk primary key,
    identity_uid text not null,
    role user_role DEFAULT 'user' NOT NULL,
    deleted boolean default false not null,
    created_at timestamp default now(),
    updated_at timestamp default now()
);

-- washes
create table washes (
    id uuid default gen_random_uuid() not null constraint washes_id_pk primary key,
    password text,
    owner_id uuid not null constraint washes_owner_fk references users,
    title text not null,
    terminal_key text,
    terminal_password text,
    description text not null,
    deleted boolean default false not null
);

-- transactions
CREATE TYPE transaction_status AS ENUM (
    'new',
    'authorized',
    'confirmed_not_synced',
    'confirmed',
    'canceling',
    'canceled',
    'unknown'
);

CREATE TABLE transactions (
    id uuid default gen_random_uuid() not null constraint transactions_id_pk primary key,
    wash_id TEXT NOT NULL,
    post_id TEXT NOT NULL,
    payment_id_bank TEXT,
    amount Integer NOT NULL,
    status transaction_status DEFAULT 'new' NOT NULL,
    created_at timestamp default now(),
    update_at timestamp default now()
);