# Build JWT authenticated Restful APIs with Golang

### Dependencies

#### Download dependencies
```shell
go mod download
```

#### Create DB and Table
```sql
CREATE DATABASE dbname;
USE dbname;

CREATE TABLE users (
    id         serial4 NOT NULL,
    email      varchar NOT NULL,
    "password" varchar NOT NULL,
    CONSTRAINT users_email_unique UNIQUE (email),
    CONSTRAINT users_pk PRIMARY KEY (id)
);
```

#### Create .env file and modify it
```shell
cp .env.dist .env
```

### Run the program
```shell
go run main.go
```

### Curl for end-points

##### Signup
```shell
curl --location --request POST 'http://localhost:8000/signup' \
--header 'Content-Type: text/plain' \
--data-raw '{
    "email": "email@example.com",
    "password": "password"
}'
```

##### Login
```shell
curl --location --request POST 'http://localhost:8000/login' \
--header 'Content-Type: text/plain' \
--data-raw '{
    "email": "email@example.com",
    "password": "password"
}'
```

##### Get all users
```shell
curl --location --request GET 'http://localhost:8000/get_all_users' \
--header 'Authorization: Bearer eyJhbG...fT2XZs'
```
