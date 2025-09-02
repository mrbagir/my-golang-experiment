import http from 'k6/http';
import { check } from 'k6';
import login from './auth.js'; 

export let options = {
    vus: 1,
    duration: '30s',
};

export default function () {
    // Test login functionality
    const loginResponse = login('testuser', 'testpassword', 'testbranch');
    
    if (loginResponse && loginResponse.success) {
        console.log('Login successful');
        // You can use the token from loginResponse for authenticated requests
    } else {
        console.log('Login failed');
    }
    
    // Example of another HTTP request
    const url = 'https://test.k6.io';
    const response = http.get(url);
    
    check(response, {
        'status is 200': (r) => r.status === 200,
        'body size is greater than 1000 bytes': (r) => r.body.length > 1000,
    });
    
    console.log(`Response time was ${response.timings.duration} ms`);
}