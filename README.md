# Payment System

## Overview
Payment System is a service that handles client registration, payment processing, and basket management by iyzico integration. It utilizes MongoDB as its database and provides a RESTful API for interactions.

## Project Structure
```sh
payment-system/
├── cmd/
│   └── server/
│       └── server.go
├── domain/
│   └── entity/
│       ├── basket.go
│       ├── client.go
│       └── payment.go
|       └── authorization.go 
|       └── basket.request.go 
|       └── payment.response.go 
|       └── iyzico_payment.response.go 
|       └── iyzico_payment.request.go 
|       └── client.response.go 
|       └── client.request.go 
├── pkg/
│   ├── config/
│   │   └── config.go
│   ├── controller/
│   │   ├── basket.go
│   │   ├── client.go
│   │   └── payment.go
│   ├── database/
│   │   └── database.go
│   ├── repository/
│   │   ├── basket.go
│   │   └── client.go
|   |   └── iyzico_authorization.go
|   |   └── iyzico.go
|   ├── formatter/
|   |   ├── formatter.go
|   ├── interfaces/
│   |   ├── basket.go
│   |   └── client.go
|   |   └── payment.go
│   ├── router/
│   │   └── router.go
│   └── usecase/
│       ├── basket.go
│       ├── client.go
│       └── payment.go
└── main.go
```
## Prerequisites

- Go 1.16+
- Docker
- MongoDB

## Setup

1. **Clone the repository**
    ```sh
    git clone https://github.com/berkcansavur/payment-system.git
    cd payment-system
    ```

2. **Run MongoDB using Docker**
    ```sh
    docker run --name mongodb -d -p 27017:27017 mongo
    ```
    or
    ```sh
    docker compose up -d
    ```

3. **Load environment variables**
    Create a `.env` file in the root directory with the following content:
    ```env
    DB_URI=mongodb://localhost:27017
    IYZICO_API_KEY=your_api_key
    IYZICO_SECRET=your_secret
    IYZICO_BASE_URL=https://sandbox-api.iyzipay.com
    ```


4. **Launch Json for Debugging**
    ```json
    {
        "version": "0.2.0",
        "configurations": [
            {
                "name": "Debug Main",
                "type": "go",
                "request": "launch",
                "mode": "debug",
                "envFile": "${workspaceFolder}/.env",
                "program": "${workspaceFolder}/main.go"
            }
        ]
    }
    ```


5. **Run the application**
    ```sh
    go run main.go
    ```
  
## API Endpoints

### Client
- `POST /client/create` - Create a new client
- `DELETE /client/{id}` - Delete a client
- `POST /client/{id}` - Update a client
- `GET /client/{id}` - Get client details
- `GET /client/{id}/cards` - Get client cards
- `POST /client/{id}/cards` - Add a card to client
- `DELETE /client/{id}/cards` - Remove a card from client
- `POST /client/{id}/checkout` - Checkout with current active basket

### Payment
- `POST /payment/create` - Create a payment

### Basket
- `POST /basket/create/{clientId}` - Create a new basket for client
- `POST /basket/abort/{id}` - Delete basket with id

## Logging
The application uses standard Go logger for logging. Logs are generated at various levels (info, error) to provide insights into the application flow and for debugging purposes.
