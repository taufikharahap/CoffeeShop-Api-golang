CREATE TABLE users (
    user_id uuid DEFAULT gen_random_uuid(),
    first_name VARCHAR NULL,
    last_name VARCHAR NULL,
    email VARCHAR NOT NULL unique,
    phone VARCHAR NOT NULL unique,
    password VARCHAR NOT NULL,
    birth DATE NULL,
    gender VARCHAR null,
    image text null,
	created_at TIMESTAMP DEFAULT NOW(),
	updated_at TIMESTAMP null,
	CONSTRAINT users_pk PRIMARY KEY (user_id)
);

INSERT INTO public.users (email, phone, password)
    values 
        ('user1@email.com', '08123456789', 'user1'),
        ('user2@email.com', '08123456780', 'user2');