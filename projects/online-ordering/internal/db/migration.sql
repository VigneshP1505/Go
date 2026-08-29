CREATE TABLE orders(
    id uuid primary key,
    customer_id uuid not null,
    restaurant_id uuid not null,
    total_amount numeric(10,2) not null,
    order_status varchar(30) not null,
    created_at timestamp default now(),
    updated_at timestamp default now()
)

create table order_items (
    id uuid primary key,
    order_id uuid references orders(id) on delete cascade,
    item_name text not null,
    quantity int not null,
    price numeric(10,2) not null
)

create index idx_orders_customer on orders(customer_id);
create index idx_orders_restaurant on orders(restaurant_id);
create index idx_orders_order_status on orders(order_status)j