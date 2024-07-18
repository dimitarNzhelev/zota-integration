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

## Feedback Revision

This section revises the feedback received, focusing on enhancing clarity and structure.

### Areas for Improvement

1. **Global Variables for Test State in Deposit & Status Check Files:**

   - **Issue:** Global variables are used to hold test-related state, leading to potential encapsulation and debugging challenges. Implicit dependencies introduced by these globals are considered a bad practice.
   - **Solution:** Implement dependency injection through interfaces to encapsulate state management and improve code maintainability and testability.

2. **Deposit Validation Issues:**

   - **Issue:** The current validation mechanism for deposits is ineffective due to missing struct tags required for validator execution. Additionally, a redundant validation block in both deposit and status check functions checks for non-empty predefined strings, which is unnecessary and never triggers an error.
   - **Solution:** Ensure all structs have appropriate tags for validation. Remove the ineffective validation block and implement a more robust validation strategy that accurately checks input data.

3. **Inconsistent and Redundant Comments:**

   - **Issue:** Comments throughout the codebase are inconsistent, with some questioning the necessity of certain code blocks (e.g., `// do i really need this?`). This indicates unresolved decisions and a lack of polish in the final code delivery.
   - **Solution:** Standardize comment format and review all comments for relevance and clarity. Remove rhetorical questions and ensure all code decisions are finalized and clearly justified in comments.

By addressing these areas, the codebase will benefit from improved structure, readability, and overall quality.
