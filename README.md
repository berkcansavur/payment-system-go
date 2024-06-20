
# Payment System

## Overview
Payment System is a service that handles client registration, payment processing, and basket management. It utilizes MongoDB as its database and provides a RESTful API for interactions.

## Project Structure

payment-system/
├── cmd/
│ └── server/
│  └── server.go
| |_main.go 
├── domain/
│ └── entity/
│  ├── basket.go
│  ├── client.go
│  └── payment.go
├── interfaces/
│ ├── ibasket_repository.go
│ └── iclient_repository.go
├── pkg/
  ├── config/
  │ └── config.go
  ├── controller/
  │ ├── basket_controller.go
  │ ├── client_controller.go
  │ └── payment_controller.go
  ├── database/
  │ └── database.go
  ├── repository/
  │ ├── basket_repository.go
  │ └── client_repository.go
  ├── router/
  │ └── router.go
  └── usecase/
  ├── basket_usecase.go
  ├── client_usecase.go
  └── payment_usecase.go

## Prerequisites

- Go 1.16+
- Docker
- MongoDB

## Setup

1. **Clone the repository**
    ```sh
    git clone https://github.com/yourusername/payment-system.git
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
    ```

4. **Run the application**
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

### Payment
- `POST /payment/create` - Create a payment

### Basket
- `POST /basket/create` - Create a new basket

## Logging
The application uses `logrus` for logging. Logs are generated at various levels (info, error) to provide insights into the application flow and for debugging purposes.

## License
This project is licensed under the MIT License.
