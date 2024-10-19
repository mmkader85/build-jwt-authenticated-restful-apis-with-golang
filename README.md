# Build JWT authenticated Restful APIs with Golang

### Run the program

```
$ ./user_jwt
```

### Curl for end-points

##### Signup
```
curl --location --request POST 'http://localhost:8000/signup' \
--header 'Content-Type: text/plain' \
--data-raw '{
    "email": "email@example.com",
    "password": "password"
}'
```

##### Login
```
curl --location --request POST 'http://localhost:8000/login' \
--header 'Content-Type: text/plain' \
--data-raw '{
    "email": "email@example.com",
    "password": "password"
}'
```

##### Get all users
```
curl --location --request GET 'http://localhost:8000/get_all_users' \
--header 'Authorization: Bearer eyJhbG...fT2XZs'
```
