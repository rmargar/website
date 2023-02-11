CREATE TABLE IF NOT EXISTS posts (
    id SERIAL PRIMARY KEY,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone,
    author text DEFAULT 'Ricardo'::text,
    title text,
    content text,
    tags varchar[],
    summary text,
    url_path varchar unique not null
);
