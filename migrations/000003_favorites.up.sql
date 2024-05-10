create table public.favorites (
    favorite_id uuid default gen_random_uuid(),
	user_id uuid not null,
    product_id uuid not null,
    created_at timestamp without time zone not null default now(),
    updated_at timestamp without time zone null,
    constraint favorites_pk primary key (product_id),
	CONSTRAINT fk_favorites_users
    FOREIGN KEY (user_id)
        REFERENCES users(user_id)
        ON DELETE CASCADE,
	CONSTRAINT fk_favorites_products
    FOREIGN KEY (product_id)
        REFERENCES products(product_id)
        ON DELETE CASCADE
);