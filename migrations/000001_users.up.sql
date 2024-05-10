CREATE TABLE public.users (
    user_id uuid DEFAULT gen_random_uuid(),
    first_name VARCHAR NULL,
    last_name VARCHAR NULL,
    email VARCHAR NOT NULL unique,
    phone VARCHAR NOT NULL unique,
    password VARCHAR NOT NULL,
    role VARCHAR NOT NULL,
    birth DATE NULL,
    gender VARCHAR null,
    image text null,
	created_at TIMESTAMP without time zone not null DEFAULT NOW(),
	updated_at TIMESTAMP without time zone null,
	CONSTRAINT users_pk PRIMARY KEY (user_id)
);