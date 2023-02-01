CREATE TABLE IF NOT EXISTS posts (
    id bigint NOT NULL,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone,
    author text DEFAULT 'Ricardo'::text,
    title text,
    content text,
    primary key (id)
);


ALTER TABLE posts OWNER TO postgres;
