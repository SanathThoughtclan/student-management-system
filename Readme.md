# Student Management System
Go for backend and Mongo for DB,
The Project includes JWT-based authentication, appropriate HTTP status codes, context-based user information retrieval, and logging.
- One type of User - Admin
- Two types of models - User(admin) and Student

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

- The application configuration is stored in `config.yaml`. This file includes settings for the server, database, and JWT secret key.


## Running the Application

### install Dependencies:-
- go mod tidy

### Run the Application:-
- go run cmd/app/main.go
 or run the existing build - .\myapp.exe

### View the changes/manipulations:-
- Using mongoDB compass or express by connecting to port mongodb://localhost:27017

##  API Endpoints with a sample curl command:-

1. POST/Register: register a new admin.

	& curl -X POST http://localhost:8080/register -H "Content-Type: application/json" -d '{\"UserID\":\"1\",\"username\":\"john\",\"password\":\"123\"}'

2. POST/Login: Authenticate and login the admin using username and password.

	& curl -X POST http://localhost:8080/login -H "Content-Type: application/json" -d '{\"username\":\"sanath\",\"password\":\"123\"}'

- This generates and prints a Auth token which is madeuse in subsequent operations.

3. POST /api/students: Create a new student record.

	& curl -X POST http://localhost:8080/api/students `
	-H "Authorization: auth token generated while logging in " `
	-H "Content-Type: application/json" `
	-d '{\"first_name\": \"Raj\", \"last_name\": \"kholi\", \"course\": \"Math\", \"grade\": \"A\"}'


5. GET /api/students: Retrieve all student records.

	& curl -X GET http://localhost:8080/api/students `
	-H "Authorization:auth token generated while logging in"



6. GET /api/students/{id}: Retrieve a student record by ID.
   
	& curl -X DELETE http://localhost:8080/api/students/  student's objectID as generated by Mongo eg   66c344879f96fc80f5765a5b `
	-H "Authorization:auth token generated while logging in"



7. PUT /api/students/{id}: Update a student record by ID.

	& curl -X PUT http://localhost:8080/api/students/  student's objectID eg:  66c344879f96fc80f5765a5b `
	-H "Authorization:auth token generated while logging in " `
	-H "Content-Type: application/json" `
	-d '{\"first_name\": \"Raj\", \"last_name\": \"kholi\", \"course\": \"Math\", \"grade\": \"q\"}'


8. DELETE /api/students/{id}: Delete a student record by ID.

	& curl -X DELETE http://localhost:8080/api/students/ student's objectID eg  66c344879f96fc80f5765a5b `
	-H "Authorization:auth token generated while logging in "

## Passing user-name and ID from http transport layer to the db layer-

- 1.Middleware Layer(Auth):

The middleware injects the admin name(username) and ID into the context of the request.
```
username := claims["username"].(string)
userID := claims["user_id"].(string)
ctx := utils.NewContextWithUserName(r.Context(), username)
ctx = utils.NewContextWithUserID(ctx, userID)

```
- 2.HTTP Layer(handler):

The handler retrieves the admin name(username) from the context and uses it for further processing, such as creating and updating info of students.
```
	username, _ := utils.GetUsernameFromContext(ctx)
	student.UpdatedBy = username

```
- 3.Service Layer:

The service layer uses the username from the context passed down by the handler.
```
	user, err := s.repo.GetByUsername(ctx, username)
```
- 4 Repository(DB) Layer:
  The service passes the username to the database layer, which can use it to identify users and manipulate them.
  ```
  filter := bson.D{{Key: "username", Value: username}} 
	err := r.collection.FindOne(ctx, filter).Decode(&user)
  ```

## Persisting the configuration information in a config file:-
The config.go file loads the configuration from the confing.yaml file

- main() calls → LoadConfig().
- LoadConfig() calls → viper.ReadInConfig() to read the YAML file.
- viper.Unmarshal(config) populates the Config struct.
- Config struct is returned to main().

## Logging:-

- Log Initialization: The logging system is initialized in utils/logger.go using InitLogger() function. Logs are written to app.log with timestamps and log levels (INFO and ERROR).

- Logging Functions:

- LogInfo(message string, id string): Logs informational messages with an ID .
- LogError(message string, err error): Logs error messages with error details.

####  error handling and user input handing implemented for most cases
