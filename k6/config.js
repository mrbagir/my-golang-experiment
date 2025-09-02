// Configuration file for k6 test
// You can set these values as environment variables or modify them here

export const config = {
    // API Configuration
    baseUrl: 'http://gym-master.apps.ocp-new-dev.bri.co.id',
    
    // Signature and encryption keys
    // Set these as environment variables: K6_NEXT_PUBLIC_SIGNATURE and K6_SECRET_KEY
    publicSignature: __ENV.NEXT_PUBLIC_SIGNATURE || 'your-public-signature-here',
    secretKey: __ENV.SECRET_KEY || 'your-secret-key-here',
    
    // Test credentials (you might want to set these as environment variables too)
    testCredentials: {
        username: __ENV.TEST_USERNAME || 'testuser',
        password: __ENV.TEST_PASSWORD || 'testpassword',
        branchCode: __ENV.TEST_BRANCH_CODE || 'testbranch'
    }
};
