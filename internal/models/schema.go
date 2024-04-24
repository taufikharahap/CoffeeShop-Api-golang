package models

var schema = `
CREATE TABLE public.users (
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

create table public.products (
    product_id uuid default gen_random_uuid(),
    name varchar not null,
    category varchar not null,
    price int not null,
	discount numeric(3,2) null,
    image text null,
	description text null,
    created_at timestamp default now(),
    update_at timestamp null,
    constraint products_pk primary key (product_id)
);

create table public.favorites (
    favorite_id uuid default gen_random_uuid(),
	user_id uuid not null,
    product_id uuid not null,
    created_at timestamp default now(),
    update_at timestamp null,
    constraint favorites_pk primary key (product_id),
	CONSTRAINT fk_favorites_users
    FOREIGN KEY (user_id)
        REFERENCES users(user_id)
        ON DELETE CASCADE,
	CONSTRAINT fk_favorites_products
    FOREIGN KEY (product_id)
        REFERENCES products(product_id)
        ON DELETE CASCADE
);`
