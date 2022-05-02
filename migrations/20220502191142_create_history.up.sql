CREATE TABLE currencies (
    id serial PRIMARY KEY,
    name VARCHAR(255) NOT NULL
);

CREATE TABLE price_history (
    id serial PRIMARY KEY,
    tstamp bigint NOT NULL,
    price decimal NOT NULL,
    fk_currencies int REFERENCES currencies(id)
);
