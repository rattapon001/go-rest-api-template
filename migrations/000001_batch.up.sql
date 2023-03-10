CREATE TABLE IF NOT EXISTS batch(
    id serial PRIMARY KEY,
    sku VARCHAR,
    ref VARCHAR,
    qty INT,
    eta VARCHAR
);