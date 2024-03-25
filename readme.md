## API

### Description:
This project is an API developed in Go, utilizing JWT for authentication and MongoDB for data storage. It serves as a comprehensive demonstration of API development principles and integration with essential tools and technologies.

### Introduction:
This API project serves as a learning exercise for developing APIs in Go, incorporating features such as JWT authentication and MongoDB integration. The primary goal is to provide a practical example of API development using modern technologies.

### Installation:
1. Clone the repository.
2. Ensure Go is installed on your system.
3. Install necessary dependencies using `go mod tidy`.

### Usage:
To use this API:
- Ensure MongoDB is running locally or configured properly.
- Set up environment variables (if any).
- Build the project using `go build`.
- Run the executable file generated.

### Authentication:
JWT authentication is implemented in this API. Users can obtain a JWT token by logging in with valid credentials. This token should be included in the headers of subsequent requests to access protected endpoints.


### Deployment:
To deploy the application:
1. Create a `bin` folder in the root of the project.
2. Build the project using `go build`.
3. Place the generated `.exe` file and `config.yml` in the `bin` folder.
4. Run the `.exe` file from the `bin` folder.
5. Test the deployed application using Postman or similar tools.

### Technologies Used:
- Go: Programming language for backend development.
- JWT: Authentication mechanism.
- MongoDB: NoSQL database for data storage.

### License:
This project is licensed under the [MIT License](LICENSE).

---
**How to release our application?**
To release the application:
1. Create a `bin` folder in the root of the project to receive the desired output.
2. In the root of the project, execute the following command in PowerShell: `go build`.
3. Obtain the `.exe` file generated and test it using Postman.
4. Place the `.exe` file along with `config.yml` in the `bin` folder.
5. Run the application from the `bin` folder.

