# EVERMOS BACKEND API
##### Backend service for Evermos Project Based Internship delivered on Rakamin Academy. This API was built using Go, Gin, MySQL, and GORM, implementing authentication, authorization, simplify product management and transaction processing.  

## Features
* User registration & login with JWT
* Role-based authorization (Admin & User)
* Automatic store creation for new users
* Product management (ownership-based)
* Category management (admin-only)
* Address management
* Transaction processing with product snapshot (transaction log)
* External API integration (EMSIFA – Indonesian region data)

## Tech Stack
This App were built with these resources:
* [GO/Golang](https://go.dev/doc/) - Bahasa Pemrograman yang dikembangkan oleh Google.
* [Gin Web Framework](https://gin-gonic.com/en/docs/) - One of high-performance HTTP web framework that wrriten on Go/Golang.
* [MySQL](https://dev.mysql.com/doc/) - Integrated Database service that are most used for developers.
* [GORM](https://gorm.io/docs/) - An ORM Library for Go/Golang Projects that are easy to use
* [JWT Authentication](https://www.jwt.io/introduction#what-is-json-web-token) - JSON Web Token (JWT) is an open standard (RFC 7519) that defines a compact and self-contained way for securely transmitting information between parties as a JSON object.

## Authentication & Authorization
* JWT is used for securing protected endpoints
* User identity is extracted from JWT payload
* Admin-only access is enforced using middleware
* Ownership validation is handled server-side (no sensitive data from request body)

## External API (EMSIFA)
This project integrates EMSIFA Public API for Indonesian region reference data

Proxy endpoints:
* GET /provinces
* GET /cities?province_id={id}

Data from EMSIFA is NOT stored in the database, only forwarded to the client.

## API Endpoints
### Auth
* POST /register
* POST /login

### Store & Product
* POST /products
* GET /my-products

### Category (Admin Only)
* POST /categories

### Address
* POST /addresses
* GET /my-addresses

### Transaction
* POST /transactions

### Region Reference
* GET /provinces
* GET /cities?province_id={id}

## Requirement
[x] Visual Studio Code
[x] Laragon/XAMPP (to run MySQL)
[x] Postman

## How to Run
### 1. Clone Repository
```
git clone https://github.com/wreck046/evermos-backend
cd evermos-backend
```

### 2. Setup Database
Create MySQL database:
```
CREATE DATABASE evermos_db
```

### 3. Configure Database
update database credetials in:
```
config/database.go
```

### 4. Install Dependencies
```
go mod tidy
```

### 5. Run Application
```
go run main.go
```
server will run on (default):
`http://localhost:8080`

## Testing
* API can be tested using Postman
* Authentication endpoints return JWT token
* Protected endpoints require `Authorization: Bearer <token>`

## Notes
* Password hashing (bcrypt) could be added as a future improvement
* Database transactions could be wrapped for atomic operations
* Endpoint naming follows REST principles but may differ slightly from reference

## Author
Developed as part of Evermos x Rakamin Academy Project Based Internship Program as Backend Developer

## Assumptions & Limitations
### Assumptions
* Each user has exactly one store, which is created automatically after registrations completed
* A user could ONLY manage products, addresses, and transactions that belong to their own account
* Category management is restricted to admin users only.
* Product price and name are `snapshotted` at the time of transaction to preserve transaction history integrity.
* external region data (provinces and cities) is retrieved from EMSIFA Public API and NOT persisted in the database.

### Limitation
* Password are currently stored as `plain text` for simplicity and could be enchanced using bcrypt hashing.
* Database operations are not fully wrapped in SQL transactions (atomicity could be improved).
* Stock managemnet and product quantity validation are not implemented.
* Transaction history listing endpoints (e.g. GET /transactions) are not implemented.
* External API data from EMSIFA is not cached locally.

## Credit
My LinkedIn profile [here](https://www.linkedin.com/in/vladimir-ivan-koroh-a2695a213/)
My other Repositories and projects could be found [here](https://github.com/wreck046?tab=repositories)
