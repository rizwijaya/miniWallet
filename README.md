### Mini Wallet Rest API
---
### Table of Contents

- [About](#about)
- [Getting Started](#getting-started)
- [Running Guide](#running-guide)
- [Project Structure](#project-structure)
- [Author](#author)
---

### About

Mini Wallet is a RESTful API designed for managing wallet operations such as deposits, withdrawals, and balance inquiries. This project is built using the Go programming language, leveraging the Fiber framework for high performance and scalability. It implements Clean Architecture, as popularized by Uncle Bob, ensuring maintainability and separation of concerns. The project uses a PostgreSQL database for data storage.

---

### Getting Started

##### Clone the Repository

To begin, clone the repository using the following command:
```bash
git clone https://github.com/rizwijaya/miniWallet.git
```

##### Environment Configuration
Set up your .env file by referring to .env.example. Below are the required configurations:

###### Application Settings
	• APP_DEBUG: Application debug mode (e.g., true, false)
	• APP_NAME: Application name
	• APP_SECRET: Secret key for JWT Token
	• APP_URL: Application base URL
	• APP_PORT: Port where the application runs

###### Setup Database
	• DB_HOST: Database hostname
	• DB_PORT: Database port
	• DB_NAME: Database name
	• DB_USER: Database username
	• DB_PASSWORD: Database password

###### Note: Create an empty database and link it to the application. On the first run, the system will automatically create the necessary tables and columns using migration.

#### Initialize the Project

Run the following commands to initialize the Go project:
```bash
go mod init
go mod tidy
```

Alternatively, use the Makefile for initialization:
```bash
make init
```

Once configured, proceed to run the project as outlined in the [Running Guide](#running-guide).

---
### Running Guide

You can run the project in three ways: directly, using a Makefile, or with nodemon.

+ ##### Method 1: Direct Execution
    Run the main file directly using:
    ```bash
    go run ./app/main.go
    ```
+ ##### Method 2: Using Makefile
    Use the following command to run via the Makefile:
    ```bash
    Make run
    ```
+ ##### Method 3: Using Nodemon
    For automatic reloading, run the application with nodemon:
    ```bash
    nodemon --exec go run ./app/main.go
    ```
    Or, use the Makefile:
    ```bash
    Make run-nodemon
    ```
    For detailed usage of the Makefile, refer to the Makefile in the project.
+ ##### Building and Running Binary
    To build a binary:
    ```bash
    go build ./app/main.go
    ```
    Or, use the Makefile:
    ```bash
    make build
    ```
    Run the binary directly:
    ```bash
    ./main
    ```
    Or, start using the Makefile:
    ```
    make start
    ```

----
### Project Structure
This project is based on the Clean Architecture model. Below is a simplified representation:
| Layer | Directory |
|--------------------------------|----------------|
| Frameworks & Drivers | Infrastructures|
| Interface | Interfaces |
| Usecases | Usecases |
| Entities | Domain |

Clean Architecture ensures clear boundaries and separation of concerns, promoting scalability and maintainability.

---
### Author 
##### <img src="https://github.com/rizwijaya/rizwijaya/blob/main/Assets/Developer.gif" height="22px"> Rizqi Wijaya
<p>
    <a href="https://www.linkedin.com/in/rizwijaya" target="_blank"><img alt="LinkedIn" src="https://img.shields.io/badge/linkedin-%230077B5.svg?&style=for-the-badge&logo=linkedin&logoColor=white" /></a> 
    <a href="mailto:rizwijaya58@gmail.com" target="_blank"><img alt="Gmail" src="https://img.shields.io/badge/gmail-D14836?&style=for-the-badge&logo=gmail&logoColor=white" /></a>   
</p>