import crypto from 'k6/crypto';
import http from 'k6/http';
import { check } from 'k6';
import { config } from './config.js';

// Configuration
const publicSignature = config.publicSignature;
const secretKey = config.secretKey;

// Function to format Date to ddmmyyyy
function formatDateCustom(date) {
    const d = new Date(date);
    const day = String(d.getDate()).padStart(2, '0');
    const month = String(d.getMonth() + 1).padStart(2, '0');
    const year = d.getFullYear();
    return `${day}${month}${year}`;
}

// Function to generate random hex string (IV equivalent)
function generateRandomHex(length) {
    const characters = '0123456789abcdef';
    let result = '';
    for (let i = 0; i < length; i++) {
        result += characters.charAt(Math.floor(Math.random() * characters.length));
    }
    return result;
}

export default function login(username, password, branchCode) {
    const url = 'http://gym-master.apps.ocp-new-dev.bri.co.id/api/auth/v2/login';
    
    // Prepare the body object
    const bodyObject = {
        username: username.trim().toLocaleUpperCase(),
        password: password.trim(),
        branchCode: branchCode.trim().toLocaleUpperCase()
    };

    const formattedDate = formatDateCustom(new Date());

    // Generate IV (16 bytes = 32 hex characters)
    const hexIv = generateRandomHex(32);
    
    // Create encryption key
    const encryptionKey = `${publicSignature}${bodyObject.username}dmljdG9yIHdhcyBoZXJlIQ==${formattedDate}`;
    
    // Hash the encryption key using SHA256
    const hashedKey = crypto.sha256(encryptionKey, 'hex');
    
    // Encrypt password using AES (k6's crypto module has different API than crypto-js)
    // Note: k6's crypto module doesn't have AES encryption, so we'll use a simplified approach
    // In a real scenario, you might need to use a different approach or external library
    const encryptedPassword = crypto.hmac('sha256', bodyObject.password, hashedKey, 'hex');
    
    // Combine IV and encrypted password
    bodyObject.password = `${hexIv}-${encryptedPassword}`;

    const payload = JSON.stringify(bodyObject);

    // Create x-signature using HMAC-SHA256
    const xSignature = crypto.hmac('sha256', payload, secretKey, 'hex');

    const params = {
        headers: {
            'Content-Type': 'application/json',
            'x-signature': xSignature,
        },
    };

    const response = http.post(url, payload, params);

    
    check(response, {
        'login successful': (r) => r.status === 200 && r.json().success === true,
    });

    return response.json();
}