import http from 'k6/http';
import { check, sleep } from 'k6';
import { getTransaction } from '../src/detail_transaction.js';

const BASE_URL = 'http://localhost:8889';
const AUTH_TOKEN = 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiNGI5Yjk5MmItMGYzZC00MjA0LWIwZGQtOTk1ZjQzMDg1N2FlIiwiZXhwIjoxNzE4MTcxMTk5LCJpYXQiOjE3MTgwODQ3OTl9.q33lubQC-PU8mfu_sjt687CvKsf2f329DqoK6-zfdFI';

export function createTransactionAndGetTransaction() {
    let res = http.post(`${BASE_URL}/transactions`, JSON.stringify({
        amount: 1000,
        currency: 'USD',
        payment_method: 'credit_card',
        description: 'Payment for order #1234',
    }), {
        headers: {
            'Content-Type': 'application/json',
            'Authorization': `Bearer ${AUTH_TOKEN}`
        }
    });

    check(res, { 'POST status was 200': (r) => r.status == 200 });
    
    if (res.status === 200) {
        let responseBody = JSON.parse(res.body);
        let transactionId = responseBody.transaction_id;
        getTransaction(transactionId);
    }


    sleep(1);
}
