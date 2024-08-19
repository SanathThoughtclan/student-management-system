# Student Management System
Go for backend and Mongo for DB,
The Project includes JWT-based authentication, appropriate HTTP status codes, context-based user information retrieval, and logging.

## Features

- **Student Management**: Create, update, delete, and retrieve student records.
- **JWT Authentication**: Secure access to the API endpoints using JSON Web Tokens.
- **Context-Based User Identification**: Automatically handle user information in API requests via context.
- **Logging**: Log successful and failed operations for better traceability.

## Project Structure

- **cmd/app**: Contains the entry point for the application.
- **config**: Handles application configuration, including server, database, and JWT settings.
- **handlers**: Manages HTTP request handling for different routes.
- **middlewares**: Contains middleware for JWT authentication.
- **models**: Defines the data models used in the application.
- **services**: Implements business logic for managing students.
- **utils**: Includes utility functions such as JWT generation and context management.

## Configuration

The application configuration is stored in `config.yaml`. This file includes settings for the server, database, and JWT secret key.


## Running the Application

### install Dependencies:-
go mod tidy

### Run the Application:-
go run cmd/app/main.go
or run the existing build - .\myapp.exe

### View the changes/manipulations:-
 in DB using mongoDB compass or express by connecting to port mongodb://localhost:27017

## API Endpoints with a sample curl command:-

1. POST/Register: register a new admin:-

& curl -X POST http://localhost:8080/register -H "Content-Type: application/json" -d '{\"UserID\":\"1\",\"username\":\"john\",\"password\":\"123\"}'

2. POST/Login: Authenticate and login the admin:-

& curl -X POST http://localhost:8080/login -H "Content-Type: application/json" -d '{\"username\":\"sanath\",\"password\":\"123\"}'

- This generates and prints a Auth token which is madeuse in subsequent operations.

3. POST /api/students: Create a new student record.

& curl -X POST http://localhost:8080/api/students `
-H "Authorization: auth token generated while logging in " `
-H "Content-Type: application/json" `
-d '{\"first_name\": \"Raj\", \"last_name\": \"kholi\", \"course\": \"Math\", \"grade\": \"A\"}'


4. GET /api/students: Retrieve all student records.

& curl -X GET http://localhost:8080/api/students `
-H "Authorization:auth token generated while logging in"



5. GET /api/students/{id}: Retrieve a student record by ID.
& curl -X DELETE http://localhost:8080/api/students/  <! --student's objectID eg -->  66c344879f96fc80f5765a5b `
-H "Authorization:auth token generated while logging in"



6. PUT /api/students/{id}: Update a student record by ID.

& curl -X PUT http://localhost:8080/api/students/  <! -- student's objectID eg -->  66c344879f96fc80f5765a5b `
-H "Authorization:auth token generated while logging in" `
-H "Content-Type: application/json" `
-d '{\"first_name\": \"Raj\", \"last_name\": \"kholi\", \"course\": \"Math\", \"grade\": \"q\"}'


7. DELETE /api/students/{id}: Delete a student record by ID.

& curl -X DELETE http://localhost:8080/api/students/ <! -- student's objectID eg -->  66c344879f96fc80f5765a5b `
-H "Authorization:auth token generated while logging in"