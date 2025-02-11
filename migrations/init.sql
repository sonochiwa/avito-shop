CREATE TABLE users
(
    id         SERIAL PRIMARY KEY,
    username   VARCHAR(64) NOT NULL,
    password   VARCHAR(64) NOT NULL,
    balance    INT         NOT NULL DEFAULT 1000,
    created_at TIMESTAMP   NOT NULL DEFAULT now(),
    CONSTRAINT username_unique UNIQUE (username)
);

CREATE TABLE products
(
    id    SERIAL PRIMARY KEY,
    title VARCHAR(32) NOT NULL,
    price INT         NOT NULL,
    CONSTRAINT title_unique UNIQUE (title)
);

CREATE TABLE inventory
(
    user_id    INTEGER REFERENCES users (id),
    product_id INTEGER REFERENCES products (id),
    quantity   INTEGER NOT NULL,
    CONSTRAINT user_id_product_id_unique UNIQUE (user_id, product_id)
);

CREATE TABLE ledger
(
    id         SERIAL PRIMARY KEY,
    from_user  INTEGER REFERENCES users (id),
    to_user    INTEGER REFERENCES users (id) CHECK ( to_user != from_user),
    amount     INT       NOT NULL CHECK (amount > 0),
    created_at TIMESTAMP NOT NULL DEFAULT now()
);

INSERT INTO products (title, price)
VALUES ('t-shirt', 80),
       ('cup', 20),
       ('book', 50),
       ('pen', 10),
       ('powerbank', 200),
       ('hoody', 300),
       ('umbrella', 200),
       ('socks', 10),
       ('wallet', 50),
       ('pink-hoody', 500);