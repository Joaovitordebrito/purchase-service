CREATE TABLE purchase(
    uuid VARCHAR(36) NOT NULL PRIMARY KEY,
    description VARCHAR(50) NOT NULL,
    purchase_amount FLOAT NOT NULL,
    transaction_date VARCHAR(10)
);