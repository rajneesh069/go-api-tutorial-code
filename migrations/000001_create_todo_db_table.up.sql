CREATE TABLE IF NOT EXISTS todos(
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    completed BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE OR REPLACE FUNCTION set_updated_at()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = now();
    RETURN NEW;
END
$$ LANGUAGE plpgsql;

CREATE TRIGGER trigger_set_updated_at
BEFORE UPDATE ON todos
FOR EACH ROW
EXECUTE FUNCTION set_updated_at();
