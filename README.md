# **Shoppers Backend API (Golang)**

---

## **Overview**

This project is a monolithic eCommerce backend API built with Go and the Gin framework. It provides a robust, scalable, and secure foundation for managing all eCommerce operations, including user authentication, product management, order processing, payment handling, and more. The application is designed with a modular structure, making it easy to maintain and extend.

---

## **Features**

- **User Management**: User registration, login, profile management, and authentication using JWT.
- **Product Catalog**: CRUD operations for products and categories, including search and filters.
- **Cart Management**: Add, update, and remove items in the shopping cart.
- **Order Management**: Order placement, tracking, and status updates.
- **Payment Processing**: Integration with multiple payment gateways.
- **Admin Panel**: Administrative endpoints for managing users, products, orders, and reports.
- **Analytics and Reporting**: Sales reports, user activity, and inventory tracking.
- **Notifications**: Email and SMS notifications for order updates.
- **Security**: Follows best practices for data protection, authentication, and authorization.

---

## **Technology Stack**

- **Language**: Go
- **Framework**: Gin
- **Database**: PostgreSQL (primary), Redis (caching), SQLite (for local development/testing)
- **ORM**: GORM
- **Authentication**: JWT
- **Message Broker**: RabbitMQ (optional, for event handling)
- **File Storage**: AWS S3 or local storage for product images
- **Testing**: Go testing tools with mocks
- **Docker**: Containerization for deployment

---

## **Project Structure**

```plaintext
project/
├── cmd/
│   └── server/
│       └── main.go             # Entry point of the application
├── internal/                   # Application-specific business logic
│   ├── modules/                # Modularized features
│   │   ├── users/              # User module
│   │   │   ├── controllers/    # Handles HTTP requests for users
│   │   │   │   ├── user_controller.go
│   │   │   ├── services/       # Contains business logic
│   │   │   │   ├── user_service.go
│   │   │   ├── repositories/   # Handles database interactions
│   │   │   │   ├── user_repository.go
│   │   │   ├── routes/         # Defines routes for the user module
│   │   │   │   ├── user_routes.go
│   │   ├── products/           # Product module
│   │   │   ├── controllers/    # Handles HTTP requests for products
│   │   │   │   ├── product_controller.go
│   │   │   ├── services/       # Contains business logic
│   │   │   │   ├── product_service.go
│   │   │   ├── repositories/   # Handles database interactions
│   │   │   │   ├── product_repository.go
│   │   │   ├── routes/         # Defines routes for the product module
│   │   │   │   ├── product_routes.go
│   │   ├── orders/             # Order module
│   │       ├── controllers/    # Handles HTTP requests for orders
│   │       ├── services/       # Contains business logic
│   │       ├── repositories/   # Handles database interactions
│   │       ├── routes/         # Defines routes for the order module
│   │
│   ├── middlewares/            # Application-wide middleware
│   │   ├── auth.go             # JWT authentication
│   │   ├── logging.go          # Request/response logging
│   │   ├── recovery.go         # Error recovery middleware
├── pkg/                        # Shared libraries/utilities
│   ├── config/                 # Application configuration
│   │   ├── config.go           # Reads and parses environment variables
│   ├── db/                     # Database initialization
│   │   ├── postgres.go         # PostgreSQL connection
│   │   ├── redis.go            # Redis connection
│   ├── logger/                 # Centralized logging utility
│   │   ├── logger.go
│   ├── utils/                  # Helper functions used across modules
│       ├── hash.go             # Password hashing
│       ├── response.go         # JSON response helpers
│       ├── validator.go        # Request validation
├── tests/                      # Unit and integration tests
│   ├── modules/
│   │   ├── users/              # Tests for user module
│   │   ├── products/           # Tests for product module
│   │   ├── orders/             # Tests for order module
│   ├── middlewares_test.go
├── docs/                       # API documentation
│   ├── swagger.yaml            # OpenAPI specification
├── go.mod                      # Go module file
├── go.sum
└── README.md                   # Project documentation
```

---

## **Installation**

### Prerequisites
- Go 1.20+
- PostgreSQL
- Redis (optional, for caching)
- Docker (optional, for containerization)

### Setup
1. Clone the repository:
   ```bash
   git clone https://github.com/your-username/ecommerce-backend.git
   cd ecommerce-backend
   ```

2. Install dependencies:
   ```bash
   go mod tidy
   ```

3. Set up environment variables:
   Create a `.env` file based on `.env.example` and configure the database, JWT secret, and other settings.

4. Run database migrations:
   ```bash
   go run cmd/migrate/main.go
   ```

5. Start the server:
   ```bash
   go run cmd/server/main.go
   ```

6. Access the API at `http://localhost:8080`.

---

## **API Endpoints**

### **Authentication**
- `POST /auth/register`: Register a new user
- `POST /auth/login`: User login
- `GET /auth/profile`: Get user profile (requires token)
- `PUT /auth/profile`: Update user profile

### **Product Management**
- `GET /products`: Get a list of products
- `GET /products/:id`: Get product details
- `POST /products`: Create a new product (Admin)
- `PUT /products/:id`: Update a product (Admin)
- `DELETE /products/:id`: Delete a product (Admin)

### **Category Management**
- `GET /categories`: Get a list of categories
- `POST /categories`: Create a new category (Admin)
- `PUT /categories/:id`: Update a category (Admin)
- `DELETE /categories/:id`: Delete a category (Admin)

### **Cart Management**
- `GET /cart`: Get the user's cart
- `POST /cart/add`: Add an item to the cart
- `PUT /cart/update`: Update cart item quantity
- `DELETE /cart/remove`: Remove an item from the cart

### **Order Management**
- `POST /orders`: Place an order
- `GET /orders/:id`: Get order details
- `GET /orders`: Get a list of user orders
- `PUT /orders/:id/cancel`: Cancel an order

### **Payment Processing**
- `POST /payments/process`: Process a payment
- `GET /payments/status/:id`: Get payment status

### **Admin Operations**
- `GET /admin/users`: List all users
- `GET /admin/orders`: List all orders
- `GET /admin/reports`: Generate reports

---

## **Best Practices**

1. **Modular Design**:
    - Each module has its own business logic, controllers, and routes, ensuring clean separation of concerns.

2. **Secure Authentication and Authorization**:
    - JWT for secure token-based authentication.
    - Role-based access control for endpoints.

3. **Database Transactions**:
    - Wrap critical operations (e.g., order placement and payment) in transactions to ensure data consistency.

4. **Error Handling**:
    - Centralized error handling middleware for uniform responses.

5. **Caching**:
    - Use Redis for frequently accessed data like product listings.

6. **Scalability**:
    - Modular architecture allows easy scaling of individual components.

7. **Logging**:
    - Structured logging for debugging and monitoring.

8. **Testing**:
    - Write unit and integration tests for all critical paths.

---

## **Deployment**

1. Build the application:
   ```bash
   go build -o ecommerce-backend ./cmd/server
   ```

2. Create a Docker image:
   ```bash
   docker build -t ecommerce-backend .
   ```

3. Run with Docker Compose:
   ```bash
   docker-compose up
   ```

---

## **Contributing**

We welcome contributions! Please submit a pull request or create an issue for discussion.

---

## **License**

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

---

## **Acknowledgments**

Special thanks to the Go and Gin communities and contributors to the open-source libraries used in this project.
