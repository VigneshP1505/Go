CREATE TABLE orders(
    id uuid primary key,
    customer_id uuid not null,
    restaurant_id uuid not null,
    total_amount numeric(10,2) not null,
    order_status varchar(30) not null,
    created_at timestamp default now(),
    updated_at timestamp default now()
);

CREATE TABLE order_items (
    id uuid primary key,
    order_id uuid references orders(id) on delete cascade,
    item_name text not null,
    quantity int not null,
    price numeric(10,2) not null
);

create index idx_orders_customer on orders(customer_id);
create index idx_orders_restaurant on orders(restaurant_id);
create index idx_orders_order_status on orders(order_status);

create table customers(
    id uuid primary key,
    customer_name string not null,
    customer_email string not null,
    unit string not null,
    street string not null,
    city string not null,
    postal string not null
)

create index idx_customers_email on customers(customer_email)

create table outbox_events(
    id uuid primary key,
    event_type varchar(100) not null,
    aggregate_id uuid not null,
    payload jsonb not null,
    published boolean default false,
    created_at timestamp default now(),
    published_at timestamp
)

create index idx_outbox_events on outbox_events(published, created_at)