# Mini Multi-Currency Wallet Service

A full-stack technical assessment implementing a simplified multi-currency wallet system inspired by modern fintech platforms.

Built with:

* **Backend:** Go, Gin, GORM, SQLite
* **Frontend:** React, TypeScript, Vite, React Query

---

# Overview

This project simulates a simplified multi-currency wallet service where users can:

* Create wallet accounts
* Deposit funds
* Transfer money between accounts
* Perform cross-currency transfers using predefined FX rates
* View account balances
* View transaction history
* Prevent duplicate transfers using Idempotency Keys

The implementation focuses on correctness, API design, and reliability rather than extensive UI styling.

---

# Features

## Backend

* Create wallet accounts
* Retrieve all accounts
* Retrieve a single account
* Deposit money
* Transfer between accounts
* Cross-currency conversion
* Transaction ledger
* Paginated transaction history
* Idempotency protection
* Input validation
* Graceful error handling

---

## Frontend

* Accounts dashboard
* Account details page
* Deposit funds
* Transfer funds
* Dynamic FX conversion preview
* Transaction history
* Loading states
* Empty states
* Error handling
* Success feedback
* Responsive layout

---

# Tech Stack

## Backend

* Go
* Gin
* GORM
* SQLite

## Frontend

* React
* TypeScript
* Vite
* TanStack React Query
* Axios
* React Router

---

# Project Structure

```text
mini-wallet-service/

├── backend/
│   ├── database/
│   ├── handlers/
│   ├── models/
│   ├── repositories/
│   ├── routes/
│   ├── services/
│   ├── utils/
│   ├── main.go
│   └── go.mod
│
└── frontend/
    ├── src/
    │   ├── api/
    │   ├── components/
    │   ├── hooks/
    │   ├── pages/
    │   ├── services/
    │   ├── types/
    │   ├── utils/
    │   └── App.tsx
    └── package.json
```

---

# API Endpoints

## Create Account

```
POST /accounts
```

Request

```json
{
    "name":"John Doe",
    "currency":"USD"
}
```

---

## Get All Accounts

```
GET /accounts
```

---

## Get Account

```
GET /accounts/:id
```

---

## Deposit

```
POST /accounts/:id/deposit
```

Request

```json
{
    "amount":1000
}
```

---

## Transfer

```
POST /transfers
```

Headers

```
Idempotency-Key: 6b8a4d0d-4d39-44b9-a26b-9f0dc6d6c9ef
```

Request

```json
{
    "from_account_id":"ACCOUNT_ID",
    "to_account_id":"ACCOUNT_ID",
    "amount":100
}
```

---

## Get Transactions

```
GET /accounts/:id/transactions?page=1
```

---

# Static Exchange Rates

```text
USD -> NGN = 1550
USD -> EUR = 0.92
USD -> GBP = 0.79
```

Conversions between non-USD currencies are performed using the application's conversion utility based on these fixed reference rates.

---

# Cross-Currency Transfers

Example

```
100 USD

↓

155,000 NGN
```

The frontend displays a live conversion preview before the transfer is submitted.

---

# Idempotency

Transfers require an **Idempotency-Key** header.

Duplicate requests with the same key are safely ignored.

Example

```
POST /transfers

Idempotency-Key:
abc123xyz
```

Submitting the same request again returns:

```json
{
    "message":"Transfer already processed"
}
```

without changing account balances.

---

# Pagination

Transaction history supports pagination.

Example

```
GET /accounts/{id}/transactions?page=2
```

Current implementation returns:

* page
* limit
* transactions

---

# Error Handling

The API returns meaningful errors for scenarios such as:

* Invalid request body
* Missing idempotency key
* Account not found
* Unsupported currency conversion
* Insufficient funds
* Invalid transfer destination
* Internal server errors

---

# Frontend Features

## Accounts Page

Displays

* Account name
* Currency
* Current balance

---

## Account Details

Displays

* Account metadata
* Current balance
* Paginated transaction history

---

## Transfer Page

Supports

* Source account selection
* Destination account selection
* Amount input
* Live currency conversion preview
* Validation
* Loading state
* Success state
* Error state

---

# Running the Project

## Backend

```bash
cd backend

go mod tidy

go run main.go
```

Backend runs on

```
http://localhost:8080
```

---

## Frontend

```bash
cd frontend

npm install

npm run dev
```

Frontend runs on

```
http://localhost:5173
```

---

# Testing Data

Example accounts

| Name               | Currency | Balance |
| ------------------ | -------- | ------- |
| John Doe           | USD      | 1000    |
| Sarah Johnson      | NGN      | 500000  |
| Michael Brown      | EUR      | 750     |
| Emma Wilson        | GBP      | 400     |
| Corporate Treasury | USD      | 10000   |
| Payroll Account    | NGN      | 2500000 |

---

# Assumptions

The assessment brief leaves several implementation details open. The following assumptions were made:

* SQLite is used for simplicity and portability.
* Fixed exchange rates are hardcoded.
* Idempotency keys are stored in memory for demonstration purposes.
* Account balances are stored as floating-point values for simplicity. In a production financial system, fixed-point integers or decimal types should be used to avoid floating-point precision issues.
* Authentication and authorization are intentionally omitted because they are outside the scope of the assessment.

---

# Future Improvements

Given more time, the following enhancements would be implemented:

* Persistent idempotency storage
* Database transactions for atomic transfers
* Optimistic locking
* Authentication and authorization
* Docker support
* Swagger/OpenAPI documentation
* Comprehensive unit and integration tests
* CI/CD pipeline
* PostgreSQL support
* Redis-backed idempotency store
* Decimal arithmetic for monetary values
* Audit logging
* Role-based permissions

---

# Video Demonstration

A short demonstration video is available here:

**Add your Loom, YouTube (unlisted), or Google Drive link here**

---

# Author

**Polycarp Atalor**

FullStack Developer

GitHub:
https://github.com/justPoly
