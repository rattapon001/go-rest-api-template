CREATE TABLE IF NOT EXISTS allocate(
    id serial PRIMARY KEY,
    sku VARCHAR,
    qty INT,
    order_line VARCHAR
);