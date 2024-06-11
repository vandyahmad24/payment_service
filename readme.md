# Payment Service

## Running Service 
./run_services.sh



### Table DDL

```
CREATE TABLE customers (
    customer_id VARCHAR(50) PRIMARY KEY,
    name VARCHAR(100),
    email VARCHAR(100),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_customers_email ON customers(email);


CREATE TABLE payment_methods (
    payment_method_id INT AUTO_INCREMENT PRIMARY KEY,
    method_name VARCHAR(50)
);

CREATE INDEX idx_payment_methods_method_name ON payment_methods(method_name);


CREATE TABLE transactions (
    transaction_id VARCHAR(50) PRIMARY KEY,
    amount DECIMAL(10, 2),
    currency VARCHAR(10),
    payment_method_id INT,
    description VARCHAR(255),
    customer_id VARCHAR(50),
    status VARCHAR(20),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (payment_method_id) REFERENCES payment_methods(payment_method_id),
    FOREIGN KEY (customer_id) REFERENCES customers(customer_id)
);

create index customer_id
    on transactions (customer_id);

create index payment_method_id
    on transactions (payment_method_id);


```