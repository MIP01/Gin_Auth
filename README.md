# How to setup
***

1. create .evv and add
    * ```DATABASE_URL=your-username:your-password@tcp(localhost:3306)/your-db_name?charset=utf8mb4&parseTime=True&loc=Local```
    * ```JWT_SECRET_KEY = your-secret-key```
2. execute ```go mod init go_auth```
3. execute ```go mod tidy```
