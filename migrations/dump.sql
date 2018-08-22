create table orders (
    order_id serial primary key,
    customer_id int NOT NULL,
    product_id int NOT NULL,
    created_at timestamp default current_timestamp
);
