ALTER TABLE public.washes DROP CONSTRAINT unique_terminal_key_terminal_password;

ALTER TABLE public.washes
ADD CONSTRAINT unique_terminal_key_terminal_password UNIQUE (terminal_key, terminal_password) WHERE (deleted = false);