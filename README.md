# Zota Developer Challenge

## Overview

The challenge involves integrating Zota to handle two key payment flows:

- **Deposit (Non-Credit Card)**
- **Status Check**

At least one unit test is mandatory.

## Getting Started

### Prerequisites

Before you begin, ensure you have the following environment variables set up in your system:

```plaintext
ENDPOINT_ID="<your_endpoint_id>"
MERCHANT_SECRET_KEY="<your_merchant_secret_key>"
MERCHANT_ID="<your_merchant_id>"
URL="<zota_api_url>"
```

These are crucial for the authentication and successful requests to the Zota API.

### Documentation

For detailed instructions and endpoints, refer to the [public documentation](https://doc.zota.com/deposit/1.0/) here.

## Running the Code

### Unit Tests

To ensure your integration works as expected, run the unit tests with the following command:

```bash
go test ./...
```

This command recursively tests all packages in the project.

### Example Code

To run the example code, execute the following command after setting up your environment variables:

```bash
go run example.go
```

This will initiate the deposit and status check flows based on the provided example.
