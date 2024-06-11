import { createTransaction } from '../src/transaction.js';
// import { getCustomerTransactions } from '../src/customer_transaction.js';


export let options = {
    stages: [
        { duration: '30s', target: 100 },
        { duration: '1m', target: 200 },
        { duration: '2m', target: 500 },
        { duration: '2m', target: 0 },
    ],
};

export default function () {
    createTransaction();
    // getCustomerTransactions('4b9b992b-0f3d-4204-b0dd-995f430857ae');
    
}
