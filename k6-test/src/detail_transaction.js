import http from 'k6/http';
import { check, sleep } from 'k6';

const BASE_URL = 'http://localhost:8889';
const AUTH_TOKEN = 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiNGI5Yjk5MmItMGYzZC00MjA0LWIwZGQtOTk1ZjQzMDg1N2FlIiwiZXhwIjoxNzE4MTcxMTk5LCJpYXQiOjE3MTgwODQ3OTl9.q33lubQC-PU8mfu_sjt687CvKsf2f329DqoK6-zfdFI';

export function getTransaction(transactionId) {
    let res = http.get(`${BASE_URL}/transactions/${transactionId}`, {
        headers: { 'Authorization': `Bearer ${AUTH_TOKEN}` }
    });

    check(res, { 'GET transaction status was 200': (r) => r.status == 200 });
    sleep(1);
}