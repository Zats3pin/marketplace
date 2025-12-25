CREATE TABLE users (
                       id SERIAL PRIMARY KEY,
                       username VARCHAR(50) NOT NULL UNIQUE,
                       email VARCHAR(100) NOT NULL UNIQUE,
                       created_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE products (
                          id SERIAL PRIMARY KEY,
                          name VARCHAR(100) NOT NULL,
                          price BIGINT NOT NULL,
                          stock INT NOT NULL DEFAULT 0,
                          created_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE orders (
                        id SERIAL PRIMARY KEY,
                        user_id INT REFERENCES users(id),
                        total BIGINT NOT NULL,
                        status VARCHAR(20) NOT NULL DEFAULT 'created',
                        created_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE order_items (
                             id SERIAL PRIMARY KEY,
                             order_id INT REFERENCES orders(id),
                             product_id INT REFERENCES products(id),
                             quantity INT NOT NULL,
                             price BIGINT NOT NULL
);