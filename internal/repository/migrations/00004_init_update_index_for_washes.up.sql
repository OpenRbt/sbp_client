DROP INDEX IF EXISTS idx_unique_terminal_key_terminal_password;

CREATE UNIQUE INDEX idx_unique_terminal_key_terminal_password
ON public.washes (terminal_key, terminal_password)
WHERE (deleted = false)