#!/bin/bash

# Navigate to api-gateway and run make run-api
cd api-gateway || exit
make run-api &

# Navigate back to the root directory
cd ..

# Navigate to customer-service and run make run-api
cd customer-service || exit
make run-api &

# Navigate back to the root directory
cd ..

# Navigate to payment-method-service and run make run-api
cd payment-method-service || exit
make run-api &

# Navigate back to the root directory
cd ..

# Navigate to transaction-service and run make run-api
cd transcation-service || exit
make run-api &
