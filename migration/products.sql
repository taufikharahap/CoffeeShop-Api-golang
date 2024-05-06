create table public.products (
    product_id uuid default gen_random_uuid(),
    name varchar not null,
    category varchar not null,
    price int not null,
	discount numeric(3,2) not null default 0,
    image_url text not null default '',
	description text null,
    created_at timestamp without time zone not null default now(),
    updated_at timestamp without time zone null,
    constraint products_pk primary key (product_id)
);

insert into public.products (name, category, price)
	values 
		('sanger coffee', 'coffee', 30000),
		('nasi goreng', 'foods', 25000),
		('teh susu', 'noncoffee', 10000);

insert into public.products (name, category, price, discount)
	values 
		('kopi susu', 'coffee', 30000, 0.10);