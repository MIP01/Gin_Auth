# Description
***
A secure API Auth built using Gin, utilizing JWT for authentication, along with data validation and password hashing using bcrypt. The system is designed to provide a  centralized authentication, enabling seamless user management and access control across various services.

# How to setup
***
1. create .env and insert
    * ```DATABASE_URL=your-username:your-password@tcp(localhost:3306)/your-db_name?charset=utf8mb4&parseTime=True&loc=Local```
    * ```JWT_SECRET_KEY = your-secret-key```
2. execute ```go mod init Gin_Auth```
3. execute ```go mod tidy```
