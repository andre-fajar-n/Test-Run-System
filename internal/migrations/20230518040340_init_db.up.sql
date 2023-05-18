create table if not exists product (
    id serial primary key,
    name varchar(100) not null,
    stock integer not null,
    expiry_date date,
    version integer not null default 0,
    created_at timestamptz not null default now(),
    updated_at timestamptz,
    deleted_at timestamptz,
    created_by varchar(10),
    updated_by varchar(10),
    deleted_by varchar(10)
);

create index idx_product_lower_name on product (lower(name));

create table if not exists product_activity_history (
    id serial primary key,
    product_id integer not null,
    action varchar(1024),
    created_at timestamptz not null default now(),
    created_by varchar(10),
    constraint fk_product foreign key(product_id) references product(id) on delete restrict on update restrict
);