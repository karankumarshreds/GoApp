CREATE TABLE IF NOT EXISTS customers (
    customer_id INT PRIMARY KEY,
    name VARCHAR(20) NOT NULL,
    date_of_birth VARCHAR(20) NOT NULL,
    city VARCHAR(20) NOT NULL,
    zipcode VARCHAR(20) NOT NULL,
    status BOOLEAN DEFAULT false
);

INSERT INTO customers (customer_id, name, date_of_birth, city, zipcode, status)
VALUES
(1000, 'Steve1', '00/00/01', 'Delhi', '110021', true),
(1001, 'Steve2', '00/00/02', 'Delhi', '110022', true),
(1002, 'Steve3', '00/00/03', 'Delhi', '110023', true),
(1003, 'Steve4', '00/00/04', 'Delhi', '110024', true),
(1004, 'Steve5', '00/00/05', 'Delhi', '110025', true);

CREATE TABLE IF NOT EXISTS accounts (
    account_id SERIAL PRIMARY KEY,
    customer_id INT UNIQUE NOT NULL,
    opening_date TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    amount FLOAT NOT NULL DEFAULT 0,
    account_type VARCHAR NOT NULL,
    status VARCHAR DEFAULT ''
);