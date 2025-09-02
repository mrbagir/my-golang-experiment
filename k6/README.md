# K6 Authentication Test Project

This project contains k6 performance tests with authentication functionality that includes password encryption and request signing, converted from a Postman pre-request script.

## Files

- `auth.js` - Authentication module with encryption and signature functionality
- `k6-test.js` - Main k6 test script
- `config.js` - Configuration file for API endpoints and credentials

## Features

- Password encryption using AES-like algorithm
- Request signing using HMAC-SHA256
- Date-based encryption key generation
- Configurable environment variables

## Setup

1. Install k6 if you haven't already:
   ```bash
   # Windows
   choco install k6
   
   # macOS
   brew install k6
   
   # Linux
   sudo apt-key adv --keyserver hkp://keyserver.ubuntu.com:80 --recv-keys C5AD17C747E3415A3642D57D77C6C491D6AC1D69
   echo "deb https://dl.k6.io/deb stable main" | sudo tee /etc/apt/sources.list.d/k6.list
   sudo apt-get update
   sudo apt-get install k6
   ```

2. Set environment variables (recommended):
   ```bash
   export K6_NEXT_PUBLIC_SIGNATURE="your-actual-public-signature"
   export K6_SECRET_KEY="your-actual-secret-key"
   export K6_TEST_USERNAME="your-test-username"
   export K6_TEST_PASSWORD="your-test-password"
   export K6_TEST_BRANCH_CODE="your-test-branch-code"
   ```

   Or on Windows:
   ```cmd
   set K6_NEXT_PUBLIC_SIGNATURE=your-actual-public-signature
   set K6_SECRET_KEY=your-actual-secret-key
   set K6_TEST_USERNAME=your-test-username
   set K6_TEST_PASSWORD=your-test-password
   set K6_TEST_BRANCH_CODE=your-test-branch-code
   ```

## Usage

Run the k6 test:
```bash
k6 run k6-test.js
```

Run with custom options:
```bash
k6 run --vus 10 --duration 1m k6-test.js
```

## Authentication Flow

The authentication process includes:

1. **Input Sanitization**: Username and branch code are trimmed and converted to uppercase
2. **Date Formatting**: Current date is formatted as ddmmyyyy
3. **Encryption Key Generation**: Combines public signature, username, fixed string, and date
4. **Password Encryption**: Password is encrypted using the generated key
5. **Request Signing**: The entire request payload is signed using HMAC-SHA256
6. **Header Addition**: The signature is added as `x-signature` header

## Notes

- The encryption implementation has been adapted for k6's crypto module limitations
- In production, ensure all sensitive keys are properly secured
- The IV generation uses a simplified random hex generator for k6 compatibility
- Consider using k6's built-in options for managing test data and credentials

## Troubleshooting

- Ensure all required environment variables are set
- Verify that the API endpoint is accessible
- Check that the signature and encryption keys are correct
- Review the k6 logs for any authentication errors
