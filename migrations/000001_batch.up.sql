CREATE TABLE IF NOT EXISTS batches(
    id serial PRIMARY KEY,
    sku VARCHAR,
    ref VARCHAR,
    qty INT,
    eta VARCHAR
);