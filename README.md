# Mekari Technical Assessment

This project is a backend technical test for the recruitment process at Mekari. It is designed to evaluate the technical skills and understanding of backend development concepts of prospective developers.

## Description

The project involves creating a RESTful API for employee management using Golang, PostgreSQL, and Docker.

## Tech Stack
Primary Techstack
- Go
- PostgreSQL
- Docker

Golang Library
- [Gin](https://github.com/gin-gonic/gin) : Rest api framework
- [Gorm](https://github.com/go-gorm/gorm) : ORM database management
- [Wire](https://github.com/google/wire) : Automatic dependency injection
- [Zap](https://github.com/uber-go/zap) : Logging
- [Testify](https://github.com/stretchr/testify): Testing assertion
- [SQLMock](https://github.com/DATA-DOG/go-sqlmock): Testing mock sql
- [GoMock/Mockgen](https://github.com/uber-go/mock): Generating mock interface for unit testing

## Environment Variables

To run this project, you will need to add the following environment variables (.env file):

```env
PORT=3000
DB_HOST=localhost
DB_PORT=5432
DB_USER=root
DB_PASSWORD=root
DB_NAME=postgres
DB_TIMEZONE=Asia/Jakarta
ENV=LOCAL
```

if you want run this project on dockerize environment, you need to change DB_HOST=host.docker.internal. Optional: ENV change to PROD to set gin to release mode.

## Installation
Clone the project to your local directory
```bash
git clone https://github.com/aliirsyaadn/mekari-technical-assessment
```

Change directory to the project folder
```bash
cd mekari-technical-assessment
```

Run server 
```bash
make run
```

Run server on dockerize environment
```bash
make docker-up
```

The server will start running at localhost and port that specify on environment variables.


## Development
If you add changes on dependency, you can add your dependency on `app/provider.go` and simply run this command to automatic generate dependency injection
```bash
make wire
```


To run all of the test you can run this command
```bash
make unit-test
```

## API Documentation
The API documentation for this project is provided in the docs folder. It includes an `OpenAPI specification` file and a `Postman collection`.

The OpenAPI specification provides a standard, language-agnostic interface to RESTful APIs which allows both humans and computers to discover and understand the capabilities of the service without access to source code, documentation, or through network traffic inspection.

If you prefer to use Postman for API testing, you can import the provided Postman collection (`postman_collection.json`) and environment file (`postman_environment.json`). These files include pre-configured requests for all the API endpoints, making it easy to test the API's functionality.

Choose the method that best suits your workflow. Both methods provide a overview of the API's capabilities, including available endpoints, request/response formats, and error codes.

## Organization Code Explaination
The `Router-Handler-Service-Repository pattern` is a common architectural pattern used in web development, especially in Golang projects. This pattern helps to structure the code in a way that promotes separation of concerns, code reusability, and maintainability.

1. `Router`: The router is responsible for handling all the routing logic of the application. It maps HTTP requests to specific handler functions based on the request method (GET, POST, PUT, DELETE) and the URL. This makes the application's endpoints clear and easy to manage. It also allows developers to easily add new routes or modify existing ones.

2. `Handler`: Handlers are functions that are executed when a specific route is hit. They contain the business logic to handle the request and send a response. Handlers interact with services to perform operations and send the results back to the client. By separating handlers from the routing logic, we can keep our code clean and easy to read.

3. `Service`: Services contain the core business logic of the application. They interact with repositories to perform CRUD operations on the database. By separating the business logic into services, we can ensure that each function has a single responsibility, making the code easier to test and maintain.

4. `Repository`: Repositories are responsible for interacting with the database. They contain the SQL queries or ORM operations to perform CRUD operations. By isolating database operations in repositories, we can change the database or the ORM without affecting the business logic.

This pattern is beneficial as it promotes code modularity and separation of concerns. Each part of the application has a specific role and interacts with the other parts through well-defined interfaces. This makes the code easier to understand, test, and maintain. It also allows for better scalability as new features or changes can be implemented with minimal impact on the existing code.

In addition to the Router-Handler-Service-Repository pattern, this project also utilizes Dependency Injection for easier management of dependencies. This approach simplifies the process of adding or modifying dependencies, making the code more flexible and maintainable.

The `config` folder is used for managing configurations such as database connections, logging settings, and automatic migrations with Gorm. Centralizing configuration settings in this manner promotes consistency and ease of access across the application.

The `model` folder is used to store primary data structures that interact with the database and handle data parsing. This separation aids in maintaining a clean and organized codebase.

Lastly, the `entity` folder contains additional struct files that help reduce code redundancy. By reusing these entities across different parts of the application, we can ensure consistency and reduce the potential for errors.

### Testing
To ensure good test coverage and maintain high code quality, I employ a modified version of `Test-Driven Development` (TDD). This approach involves writing tests first and then developing features that meet the tests' requirements.

At the repository level, I use `sqlmock` to validate SQL queries. This allows me to create unit tests for database interactions without needing an actual database. It ensures that the SQL queries are correctly formed and behave as expected.

For the service layer, which contains the core business logic, I use `mocking` repositories. This approach allows me to isolate the business logic and test it independently from the database. It ensures that the services correctly handle data and execute the expected operations.

At the handler level, I use a combination of `mocking` and `httptest` to simulate HTTP requests and responses. This allows me to test the handlers' functionality, ensuring they correctly process requests, call the appropriate services, and return the expected responses.

This testing strategy ensures each component of the application is thoroughly tested in isolation, leading to a robust and reliable application

## License
This project is MIT licensed.